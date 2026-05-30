package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB represents MongoDB client wrapper
type MongoDB struct {
	client   *mongo.Client
	database *mongo.Database
}

// NewMongoDB creates a new MongoDB connection
func NewMongoDB(ctx context.Context, uri, dbName string) (*MongoDB, error) {
	// Set client options
	clientOpts := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	database := client.Database(dbName)

	return &MongoDB{
		client:   client,
		database: database,
	}, nil
}

// Close closes the MongoDB connection
func (m *MongoDB) Close(ctx context.Context) error {
	if m.client != nil {
		return m.client.Disconnect(ctx)
	}
	return nil
}

// GetCollection returns a collection
func (m *MongoDB) GetCollection(name string) *mongo.Collection {
	return m.database.Collection(name)
}

// GetDatabase returns the database
func (m *MongoDB) GetDatabase() *mongo.Database {
	return m.database
}

// CreateIndexes creates indexes for collections
func (m *MongoDB) CreateIndexes(ctx context.Context) error {
	// Create user collection indexes
	userCollection := m.GetCollection("users")
	userIndexModel := mongo.IndexModel{
		Keys: options.Index().SetUnique(true),
	}

	if _, err := userCollection.Indexes().CreateOne(ctx, userIndexModel); err != nil {
		return fmt.Errorf("failed to create user email index: %w", err)
	}

	// Create product collection indexes
	productCollection := m.GetCollection("products")
	productIndexModel := mongo.IndexModel{
		Keys: options.Index().SetUnique(true),
	}

	if _, err := productCollection.Indexes().CreateOne(ctx, productIndexModel); err != nil {
		return fmt.Errorf("failed to create product sku index: %w", err)
	}

	return nil
}
