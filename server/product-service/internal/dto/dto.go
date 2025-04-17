package dto

import "github.com/google/uuid"

type CreateProductRequest struct {
	Name          string
	Slug          string
	Description   string
	Price         float64
	Stock         int
	CategoryID    uuid.UUID
	SubcategoryID *uuid.UUID
	IsFeatured    bool
}

type UpdateProductRequest struct {
	Name          string
	Slug          string
	Description   string
	Price         float64
	Stock         int
	CategoryID    uuid.UUID
	SubcategoryID *uuid.UUID
	IsFeatured    bool
}

type CategoryRequest struct {
	Name  string `json:"name" binding:"required"`
	Image string `json:"image"`
	Slug  string
}

type SubcategoryRequest struct {
	Name       string
	Slug       string
	CategoryID uuid.UUID
	Image      string
}

type ColorRequest struct {
	Name string `form:"name" binding:"required"`
	Hex  string `form:"hex" binding:"required"`
}

type ProductResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Stock       int      `json:"stock"`
	Sold        int      `json:"sold"`
	IsFeatured  bool     `json:"isFeatured"`
	IsActive    bool     `json:"isActive"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
	Images      []string `json:"images"`
	Category    any      `json:"category"`
	Subcategory any      `json:"subcategory"`
}
