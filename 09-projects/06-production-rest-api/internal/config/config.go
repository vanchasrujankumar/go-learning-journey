package config

import (
	"os"
	"time"
)

// Config holds all application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Kafka    KafkaConfig
	JWT      JWTConfig
	Logger   LoggerConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port         string
	Env          string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// DatabaseConfig holds MongoDB configuration
type DatabaseConfig struct {
	URI      string
	Database string
	Timeout  time.Duration
}

// KafkaConfig holds Kafka configuration
type KafkaConfig struct {
	Brokers      []string
	UserTopic    string
	ProductTopic string
	GroupID      string
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret string
	Expiry time.Duration
}

// LoggerConfig holds logger configuration
type LoggerConfig struct {
	Level string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"),
			Env:          getEnv("ENV", "development"),
			ReadTimeout:  getDurationEnv("READ_TIMEOUT", 15*time.Second),
			WriteTimeout: getDurationEnv("WRITE_TIMEOUT", 15*time.Second),
			IdleTimeout:  getDurationEnv("IDLE_TIMEOUT", 60*time.Second),
		},
		Database: DatabaseConfig{
			URI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
			Database: getEnv("MONGO_DB", "go_api_db"),
			Timeout:  getDurationEnv("MONGO_TIMEOUT", 10*time.Second),
		},
		Kafka: KafkaConfig{
			Brokers:      []string{getEnv("KAFKA_BROKERS", "localhost:9092")},
			UserTopic:    getEnv("KAFKA_USER_TOPIC", "user-events"),
			ProductTopic: getEnv("KAFKA_PRODUCT_TOPIC", "product-events"),
			GroupID:      getEnv("KAFKA_GROUP_ID", "go-api-consumer-group"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
			Expiry: getDurationEnv("JWT_EXPIRY", 24*time.Hour),
		},
		Logger: LoggerConfig{
			Level: getEnv("LOG_LEVEL", "info"),
		},
	}
}

// getEnv retrieves an environment variable with a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getDurationEnv retrieves a duration environment variable with a default value
func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		duration, err := time.ParseDuration(value)
		if err == nil {
			return duration
		}
	}
	return defaultValue
}
