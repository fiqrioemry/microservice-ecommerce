package handlers

import (
	"net/http"

	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/services"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartHandler struct {
	Service services.CartService
}

func NewCartHandler(service services.CartService) *CartHandler {
	return &CartHandler{Service: service}
}

// GET /api/cart
func (h *CartHandler) GetCart(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	cart, err := h.Service.GetUserCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch cart"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

// POST /api/cart
func (h *CartHandler) AddToCart(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// Simulasi ambil snapshot dari product-service
	// Real-nya: HTTP call ke product-service
	product := dto.ProductSnapshot{
		Name:     "Contoh Produk",
		ImageURL: "https://example.com/image.jpg",
		Price:    120000,
	}

	if err := h.Service.AddToCart(userID, req, product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add to cart", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Item added to cart"})
}

// PUT /api/cart/items/:id
func (h *CartHandler) UpdateCartItem(c *gin.Context) {
	itemID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid item ID"})
		return
	}

	var req dto.UpdateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	if err := h.Service.UpdateCartItem(itemID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update cart item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item updated"})
}

// DELETE /api/cart/items/:id
func (h *CartHandler) RemoveCartItem(c *gin.Context) {
	itemID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid item ID"})
		return
	}

	if err := h.Service.RemoveCartItem(itemID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to remove item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item removed"})
}

func (h *CartHandler) ClearCart(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	if err := h.Service.ClearUserCart(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to clear cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart cleared"})
}
