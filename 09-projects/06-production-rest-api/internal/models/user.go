package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the system
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name" validate:"required,min=2"`
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	Password  string             `bson:"password" json:"-"` // Never expose password
	Phone     string             `bson:"phone" json:"phone" validate:"omitempty,phone"`
	Address   Address            `bson:"address" json:"address"`
	Status    string             `bson:"status" json:"status"` // active, inactive, suspended
	Role      string             `bson:"role" json:"role"`     // admin, user, moderator
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

// Address represents a user address
type Address struct {
	Street  string `bson:"street" json:"street"`
	City    string `bson:"city" json:"city"`
	State   string `bson:"state" json:"state"`
	ZipCode string `bson:"zip_code" json:"zip_code"`
	Country string `bson:"country" json:"country"`
}

// CreateUserRequest is the request body for creating a user
type CreateUserRequest struct {
	Name     string  `json:"name" validate:"required,min=2"`
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,min=8"`
	Phone    string  `json:"phone" validate:"omitempty,phone"`
	Address  Address `json:"address"`
}

// UpdateUserRequest is the request body for updating a user
type UpdateUserRequest struct {
	Name    string   `json:"name" validate:"omitempty,min=2"`
	Email   string   `json:"email" validate:"omitempty,email"`
	Phone   string   `json:"phone" validate:"omitempty,phone"`
	Address *Address `json:"address"`
	Status  string   `json:"status" validate:"omitempty,oneof=active inactive suspended"`
}

// UserResponse is the response body for user endpoints
type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   Address   `json:"address"`
	Status    string    `json:"status"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToResponse converts User to UserResponse
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID.Hex(),
		Name:      u.Name,
		Email:     u.Email,
		Phone:     u.Phone,
		Address:   u.Address,
		Status:    u.Status,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
