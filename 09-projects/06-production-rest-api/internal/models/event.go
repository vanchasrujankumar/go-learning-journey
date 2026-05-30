package models

import "time"

// EventType represents the type of event
type EventType string

const (
	UserCreatedEvent    EventType = "user.created"
	UserUpdatedEvent    EventType = "user.updated"
	UserDeletedEvent    EventType = "user.deleted"
	ProductCreatedEvent EventType = "product.created"
	ProductUpdatedEvent EventType = "product.updated"
	ProductDeletedEvent EventType = "product.deleted"
)

// Event represents a Kafka event
type Event struct {
	ID        string      `json:"id"`
	Type      EventType   `json:"type"`
	Timestamp time.Time   `json:"timestamp"`
	Source    string      `json:"source"` // service name
	Data      interface{} `json:"data"`
	Version   string      `json:"version"` // event schema version
}

// UserEvent represents a user-related event
type UserEvent struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

// ProductEvent represents a product-related event
type ProductEvent struct {
	ProductID string    `json:"product_id"`
	Sku       string    `json:"sku"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int32     `json:"stock"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
