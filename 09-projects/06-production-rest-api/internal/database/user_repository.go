package database

import (
	"context"
	"fmt"

	"github.com/SrujanKumar/go-learning-journey/09-projects/06-production-rest-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserRepository handles user database operations
type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *MongoDB) *UserRepository {
	return &UserRepository{
		collection: db.GetCollection("users"),
	}
}

// Create inserts a new user into the database
func (r *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, fmt.Errorf("user with email already exists")
		}
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %w", err)
	}

	var user models.User
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// List retrieves all users with pagination
func (r *UserRepository) List(ctx context.Context, skip, limit int64) ([]*models.User, error) {
	opts := options.Find().SetSkip(skip).SetLimit(limit)
	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	defer cursor.Close(ctx)

	var users []*models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, fmt.Errorf("failed to decode users: %w", err)
	}

	return users, nil
}

// Update updates a user
func (r *UserRepository) Update(ctx context.Context, id string, update *models.User) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid user ID format: %w", err)
	}

	updateData := bson.M{
		"name":       update.Name,
		"email":      update.Email,
		"phone":      update.Phone,
		"address":    update.Address,
		"status":     update.Status,
		"updated_at": update.UpdatedAt,
	}

	result, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": updateData},
	)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("email already in use")
		}
		return fmt.Errorf("failed to update user: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// Delete deletes a user (soft delete)
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid user ID format: %w", err)
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// Count returns the total number of users
func (r *UserRepository) Count(ctx context.Context) (int64, error) {
	count, err := r.collection.EstimatedDocumentCount(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to count users: %w", err)
	}
	return count, nil
}
