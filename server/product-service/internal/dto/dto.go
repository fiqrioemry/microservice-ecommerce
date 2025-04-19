package dto

import (
	"mime/multipart"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/google/uuid"
)

type ProductResponse struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Slug        string                 `json:"slug"`
	Description string                 `json:"description"`
	IsFeatured  bool                   `json:"isFeatured"`
	IsActive    bool                   `json:"isActive"`
	Weight      float64                `json:"weight"`
	Length      float64                `json:"length"`
	Width       float64                `json:"width"`
	Height      float64                `json:"height"`
	Images      []string               `json:"images"`
	Category    models.Category        `json:"category"`
	Subcategory *models.Subcategory    `json:"subcategory,omitempty"`
	Variants    []ProductVariantOutput `json:"variants,omitempty"`
	Attributes  []AttributeOutput      `json:"attributes,omitempty"`
}

type ProductVariantOutput struct {
	ID       string            `json:"id"`
	SKU      string            `json:"sku"`
	Price    float64           `json:"price"`
	Stock    int               `json:"stock"`
	IsActive bool              `json:"isActive"`
	ImageURL string            `json:"imageUrl"`
	Options  map[string]string `json:"options"`
}

type AttributeOutput struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CreateFullProductRequest struct {
	Name          string                            `form:"name" binding:"required"`
	Description   string                            `form:"description"`
	CategoryID    uuid.UUID                         `form:"categoryId" binding:"required"`
	SubcategoryID *uuid.UUID                        `form:"subcategoryId"`
	IsFeatured    bool                              `form:"isFeatured"`
	Weight        float64                           `json:"weight"`
	Length        float64                           `json:"length"`
	Width         float64                           `json:"width"`
	Height        float64                           `json:"height"`
	Variants      []CreateVariantRequest            `json:"variants"`
	Attributes    []AddProductAttributeValueRequest `json:"attributes"`
}

type CreateVariantRequest struct {
	SKU      string            `json:"sku" binding:"required"`
	Price    float64           `json:"price" binding:"required"`
	Stock    int               `json:"stock" binding:"required"`
	IsActive bool              `json:"isActive"`
	ImageURL string            `json:"imageUrl"`
	Options  map[string]string `json:"options"`
}

type AddProductAttributeValueRequest struct {
	AttributeID      uint `json:"attributeId" binding:"required"`
	AttributeValueID uint `json:"attributeValueId" binding:"required"`
}

type CreateFullProductWithImages struct {
	Data          CreateFullProductRequest
	ProductImages []*multipart.FileHeader
	VariantImages []*multipart.FileHeader
}

type UpdateProductRequest struct {
	Name          string                            `form:"name" binding:"required"`
	Description   string                            `form:"description"`
	CategoryID    uuid.UUID                         `form:"categoryId" binding:"required"`
	SubcategoryID *uuid.UUID                        `form:"subcategoryId"`
	IsFeatured    bool                              `form:"isFeatured"`
	Weight        float64                           `json:"weight"`
	Length        float64                           `json:"length"`
	Width         float64                           `json:"width"`
	Height        float64                           `json:"height"`
	Variants      []CreateVariantRequest            `json:"variants"`
	Attributes    []AddProductAttributeValueRequest `json:"attributes"`
}

type CategoryRequest struct {
	Name  string `form:"name" binding:"required"`
	Slug  string `json:"slug"`
	Image string `json:"image"`
}

type AttributeRequest struct {
	Name string `json:"name" binding:"required"`
}

type AttributeValueRequest struct {
	AttributeID uint   `json:"attributeId" binding:"required"`
	Value       string `json:"value" binding:"required"`
}
type VariantTypeRequest struct {
	Name string `json:"name" binding:"required"`
}

type VariantValueRequest struct {
	TypeID uint   `json:"typeId" binding:"required"`
	Value  string `json:"value" binding:"required"`
}

type CategoryVariantTypeRequest struct {
	CategoryID    uuid.UUID `json:"categoryId" binding:"required"`
	VariantTypeID uint      `json:"variantTypeId" binding:"required"`
}

type SubcategoryVariantTypeRequest struct {
	SubcategoryID uuid.UUID `json:"subcategoryId" binding:"required"`
	VariantTypeID uint      `json:"variantTypeId" binding:"required"`
}

type SubcategoryRequest struct {
	Name       string    `form:"name" binding:"required"`
	Slug       string    `json:"slug"`
	CategoryID uuid.UUID `form:"categoryId" binding:"required"`
	Image      string    `form:"image"`
}

type CategoryWithSubResponse struct {
	ID            uuid.UUID             `json:"id"`
	Name          string                `json:"name"`
	Slug          string                `json:"slug"`
	Image         string                `json:"image"`
	Subcategories []SubcategoryResponse `json:"subcategories"`
}

type SubcategoryResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Slug  string    `json:"slug"`
	Image string    `json:"image"`
}
