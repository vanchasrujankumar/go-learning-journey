package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product represents a product in the system
type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name" validate:"required,min=3"`
	Description string             `bson:"description" json:"description" validate:"required,min=10"`
	Price       float64            `bson:"price" json:"price" validate:"required,gt=0"`
	Stock       int32              `bson:"stock" json:"stock" validate:"required,gte=0"`
	Category    string             `bson:"category" json:"category" validate:"required"`
	Sku         string             `bson:"sku" json:"sku" validate:"required,unique"`
	Status      string             `bson:"status" json:"status"`                    // available, out-of-stock, discontinued
	Tags        []string           `bson:"tags" json:"tags"`
	Images      []Image            `bson:"images" json:"images"`
	Metrics     ProductMetrics     `bson:"metrics" json:"metrics"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time         `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

// Image represents a product image
type Image struct {
	URL       string `bson:"url" json:"url"`
	Alt       string `bson:"alt" json:"alt"`
	Primary   bool   `bson:"primary" json:"primary"`
}

// ProductMetrics represents product metrics
type ProductMetrics struct {
	Views      int32   `bson:"views" json:"views"`
	Purchases  int32   `bson:"purchases" json:"purchases"`
	Rating     float64 `bson:"rating" json:"rating"`
	Reviews    int32   `bson:"reviews" json:"reviews"`
}

// CreateProductRequest is the request body for creating a product
type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=3"`
	Description string  `json:"description" validate:"required,min=10"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int32   `json:"stock" validate:"required,gte=0"`
	Category    string  `json:"category" validate:"required"`
	Sku         string  `json:"sku" validate:"required"`
	Tags        []string `json:"tags"`
	Images      []Image `json:"images"`
}

// UpdateProductRequest is the request body for updating a product
type UpdateProductRequest struct {
	Name        string   `json:"name" validate:"omitempty,min=3"`
	Description string   `json:"description" validate:"omitempty,min=10"`
	Price       *float64 `json:"price" validate:"omitempty,gt=0"`
	Stock       *int32   `json:"stock" validate:"omitempty,gte=0"`
	Category    string   `json:"category" validate:"omitempty"`
	Status      string   `json:"status" validate:"omitempty,oneof=available out-of-stock discontinued"`
	Tags        []string `json:"tags"`
	Images      []Image  `json:"images"`
}

// ProductResponse is the response body for product endpoints
type ProductResponse struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Stock       int32          `json:"stock"`
	Category    string         `json:"category"`
	Sku         string         `json:"sku"`
	Status      string         `json:"status"`
	Tags        []string       `json:"tags"`
	Images      []Image        `json:"images"`
	Metrics     ProductMetrics `json:"metrics"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// ToResponse converts Product to ProductResponse
func (p *Product) ToResponse() *ProductResponse {
	return &ProductResponse{
		ID:          p.ID.Hex(),
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		Category:    p.Category,
		Sku:         p.Sku,
		Status:      p.Status,
		Tags:        p.Tags,
		Images:      p.Images,
		Metrics:     p.Metrics,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
