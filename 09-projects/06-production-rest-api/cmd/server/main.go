package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SrujanKumar/go-learning-journey/09-projects/06-production-rest-api/internal/config"
	"github.com/SrujanKumar/go-learning-journey/09-projects/06-production-rest-api/internal/database"
	"github.com/SrujanKumar/go-learning-journey/09-projects/06-production-rest-api/internal/handlers"
	"github.com/SrujanKumar/go-learning-journey/09-projects/06-production-rest-api/internal/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	log.Printf("Starting API server in %s environment on port %s", cfg.Server.Env, cfg.Server.Port)

	// Create root context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Connect to MongoDB
	mongoCtx, mongoCancel := context.WithTimeout(ctx, cfg.Database.Timeout)
	defer mongoCancel()

	mongo, err := database.NewMongoDB(mongoCtx, cfg.Database.URI, cfg.Database.Database)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongo.Close(ctx)

	log.Println("✓ Connected to MongoDB")

	// Create indexes
	if err := mongo.CreateIndexes(ctx); err != nil {
		log.Fatalf("Failed to create indexes: %v", err)
	}
	log.Println("✓ Created database indexes")

	// Initialize repositories
	userRepo := database.NewUserRepository(mongo)
	log.Println("✓ Initialized repositories")

	// Initialize Kafka producer
	producer := kafka.NewProducer(cfg.Kafka.Brokers)
	defer producer.Close()
	log.Println("✓ Initialized Kafka producer")

	// Initialize Kafka consumer (in background)
	consumer := kafka.NewConsumer(cfg.Kafka.Brokers, cfg.Kafka.UserTopic, cfg.Kafka.GroupID)
	go func() {
		if err := consumer.Start(ctx, kafka.HandleUserEvent); err != nil {
			log.Printf("Consumer error: %v", err)
		}
	}()
	log.Println("✓ Started Kafka consumer")

	// Setup router
	router := chi.NewRouter()

	// Add middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Timeout(cfg.Server.ReadTimeout))

	// Health check
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userRepo, producer)

	// API v1 routes
	router.Route("/api/v1", func(r chi.Router) {
		// User routes
		r.Route("/users", func(r chi.Router) {
			r.Post("/", userHandler.CreateUser)
			r.Get("/", userHandler.ListUsers)
			r.Get("/{id}", userHandler.GetUser)
			r.Put("/{id}", userHandler.UpdateUser)
			r.Delete("/{id}", userHandler.DeleteUser)
		})
	})

	// Create HTTP server
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	// Start server in goroutine
	go func() {
		log.Printf("API Server listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	log.Printf("Received signal: %v", sig)

	// Shutdown context with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	consumer.Close()
	log.Println("✓ Server shutdown complete")
}
