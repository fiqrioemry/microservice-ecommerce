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

type CreateSizeRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateSizeRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateProductVariantRequest struct {
	ProductID string  `json:"productId" binding:"required"`
	ColorID   *uint   `json:"colorId"`
	SizeID    *uint   `json:"sizeId"`
	SKU       string  `json:"sku" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	Stock     int     `json:"stock" binding:"required"`
	IsActive  bool    `json:"isActive"`
}

type UpdateProductVariantRequest struct {
	ColorID  *uint   `json:"colorId"`
	SizeID   *uint   `json:"sizeId"`
	Price    float64 `json:"price" binding:"required"`
	Stock    int     `json:"stock" binding:"required"`
	IsActive bool    `json:"isActive"`
}

type CreateAttributeRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAttributeRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateAttributeValueRequest struct {
	AttributeID uint   `json:"attributeId" binding:"required"`
	Value       string `json:"value" binding:"required"`
}
