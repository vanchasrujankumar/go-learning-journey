package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/SrujanKumar/go-learning-journey/09-projects/06-production-rest-api/internal/models"
	"github.com/segmentio/kafka-go"
)

// Consumer handles consuming events from Kafka
type Consumer struct {
	reader *kafka.Reader
}

// MessageHandler is a function that handles Kafka messages
type MessageHandler func(ctx context.Context, message *kafka.Message) error

// NewConsumer creates a new Kafka consumer
func NewConsumer(brokers []string, topic, groupID string) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:         brokers,
		Topic:           topic,
		GroupID:         groupID,
		StartOffset:     kafka.LastOffset,
		MaxBytes:        10e6, // 10MB
		CommitInterval:  1,
		PartitionWatchInterval: 0,
	})

	return &Consumer{
		reader: reader,
	}
}

// Start starts consuming messages
func (c *Consumer) Start(ctx context.Context, handler MessageHandler) error {
	defer c.Close()

	log.Println("Starting Kafka consumer...")

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			message, err := c.reader.FetchMessage(ctx)
			if err != nil {
				log.Printf("Error fetching message: %v", err)
				continue
			}

			if err := handler(ctx, &message); err != nil {
				log.Printf("Error handling message: %v", err)
				continue
			}

			if err := c.reader.CommitMessages(ctx, message); err != nil {
				log.Printf("Error committing message: %v", err)
			}
		}
	}
}

// Close closes the consumer
func (c *Consumer) Close() error {
	if c.reader != nil {
		return c.reader.Close()
	}
	return nil
}

// HandleUserEvent handles user events
func HandleUserEvent(ctx context.Context, message *kafka.Message) error {
	var event models.Event
	if err := json.Unmarshal(message.Value, &event); err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	// Handle different event types
	switch event.Type {
	case models.UserCreatedEvent:
		log.Printf("User created: %+v", event.Data)
		// Perform actions for user creation
	case models.UserUpdatedEvent:
		log.Printf("User updated: %+v", event.Data)
		// Perform actions for user update
	case models.UserDeletedEvent:
		log.Printf("User deleted: %+v", event.Data)
		// Perform actions for user deletion
	}

	return nil
}

// HandleProductEvent handles product events
func HandleProductEvent(ctx context.Context, message *kafka.Message) error {
	var event models.Event
	if err := json.Unmarshal(message.Value, &event); err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	// Handle different event types
	switch event.Type {
	case models.ProductCreatedEvent:
		log.Printf("Product created: %+v", event.Data)
		// Perform actions for product creation
	case models.ProductUpdatedEvent:
		log.Printf("Product updated: %+v", event.Data)
		// Perform actions for product update
	case models.ProductDeletedEvent:
		log.Printf("Product deleted: %+v", event.Data)
		// Perform actions for product deletion
	}

	return nil
}
