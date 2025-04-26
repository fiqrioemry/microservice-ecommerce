package dto

import "github.com/google/uuid"

type AddToCartRequest struct {
	ProductID uuid.UUID  `json:"productId" binding:"required"`
	VariantID *uuid.UUID `json:"variantId"`
	Quantity  int        `json:"quantity" binding:"required"`
}

type UpdateCartItemRequest struct {
	ProductID uuid.UUID  `json:"productId" binding:"required"`
	Quantity  int        `json:"quantity"`
	IsChecked bool       `json:"isChecked"`
	VariantID *uuid.UUID `json:"variantId"`
}

type CartItemResponse struct {
	ID          string  `json:"id"`
	ProductID   string  `json:"productId"`
	ProductName string  `json:"name"`
	ImageURL    string  `json:"imageUrl"`
	VariantID   *string `json:"variantId,omitempty"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	IsChecked   bool    `json:"isChecked"`
}

type ProductSnapshot struct {
	Name     string  `json:"name"`
	ImageURL string  `json:"imageUrl"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}
