package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// Producer handles publishing events to Kafka
type Producer struct {
	writer *kafka.Writer
}

// NewProducer creates a new Kafka producer
func NewProducer(brokers []string) *Producer {
	dialer := &kafka.Dialer{
		Timeout: 10,
	}
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:     brokers,
		Dialer:      dialer,
		MaxAttempts: 3,
	})

	return &Producer{
		writer: writer,
	}
}

// Publish publishes an event to a Kafka topic
func (p *Producer) Publish(ctx context.Context, topic string, event interface{}) error {
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	message := kafka.Message{
		Key:   []byte(fmt.Sprintf("%v", event)),
		Value: data,
	}

	err = p.writer.WriteMessages(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

// Close closes the producer
func (p *Producer) Close() error {
	if p.writer != nil {
		return p.writer.Close()
	}
	return nil
}

// PublishBatch publishes multiple events
func (p *Producer) PublishBatch(ctx context.Context, topic string, events []interface{}) error {
	messages := make([]kafka.Message, len(events))

	for i, event := range events {
		data, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("failed to marshal event at index %d: %w", i, err)
		}

		messages[i] = kafka.Message{
			Key:   []byte(fmt.Sprintf("%d", i)),
			Value: data,
		}
	}

	err := p.writer.WriteMessages(ctx, messages...)
	if err != nil {
		return fmt.Errorf("failed to publish batch: %w", err)
	}

	return nil
}
