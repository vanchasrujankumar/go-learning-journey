package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SrujanKumar/go-learning-journey/09-projects/06-production-rest-api/internal/database"
	"github.com/SrujanKumar/go-learning-journey/09-projects/06-production-rest-api/internal/kafka"
	"github.com/SrujanKumar/go-learning-journey/09-projects/06-production-rest-api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userRepo *database.UserRepository
	producer *kafka.Producer
}

// NewUserHandler creates a new user handler
func NewUserHandler(userRepo *database.UserRepository, producer *kafka.Producer) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
		producer: producer,
	}
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		Phone:     req.Phone,
		Address:   req.Address,
		Status:    "active",
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := h.userRepo.Create(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Publish event to Kafka
	event := &models.Event{
		ID:        primitive.NewObjectID().Hex(),
		Type:      models.UserCreatedEvent,
		Timestamp: time.Now(),
		Source:    "user-service",
		Data: &models.UserEvent{
			UserID:    createdUser.ID.Hex(),
			Email:     createdUser.Email,
			Name:      createdUser.Name,
			Status:    createdUser.Status,
			Timestamp: createdUser.CreatedAt,
		},
		Version: "1.0",
	}

	h.producer.Publish(r.Context(), "user-events", event)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser.ToResponse())
}

// GetUser retrieves a user by ID
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.ToResponse())
}

// ListUsers retrieves all users
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepo.List(r.Context(), 0, 100)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responses := make([]*models.UserResponse, len(users))
	for i, user := range users {
		responses[i] = user.ToResponse()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

// UpdateUser updates a user
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.Status != "" {
		user.Status = req.Status
	}
	if req.Address != nil {
		user.Address = *req.Address
	}
	user.UpdatedAt = time.Now()

	if err := h.userRepo.Update(r.Context(), id, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Publish event to Kafka
	event := &models.Event{
		ID:        primitive.NewObjectID().Hex(),
		Type:      models.UserUpdatedEvent,
		Timestamp: time.Now(),
		Source:    "user-service",
		Data: &models.UserEvent{
			UserID:    user.ID.Hex(),
			Email:     user.Email,
			Name:      user.Name,
			Status:    user.Status,
			Timestamp: user.UpdatedAt,
		},
		Version: "1.0",
	}

	h.producer.Publish(r.Context(), "user-events", event)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.ToResponse())
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.userRepo.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Publish event to Kafka
	event := &models.Event{
		ID:        primitive.NewObjectID().Hex(),
		Type:      models.UserDeletedEvent,
		Timestamp: time.Now(),
		Source:    "user-service",
		Data: &models.UserEvent{
			UserID:    user.ID.Hex(),
			Email:     user.Email,
			Name:      user.Name,
			Timestamp: time.Now(),
		},
		Version: "1.0",
	}

	h.producer.Publish(r.Context(), "user-events", event)

	w.WriteHeader(http.StatusNoContent)
}
