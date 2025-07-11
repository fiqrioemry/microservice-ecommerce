package handlers

import (
	"net/http"

	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/services"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/grpc"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartHandler struct {
	Service    services.CartService
	GRPCClient grpc.ProductGRPCClient
}

func NewCartHandler(service services.CartService, grpcClient grpc.ProductGRPCClient) *CartHandler {
	return &CartHandler{
		Service:    service,
		GRPCClient: grpcClient,
	}
}

// GET /api/cart
func (h *CartHandler) GetCart(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	cart, err := h.Service.GetUserCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch cart"})
		return
	}

	// mapping hasil ke response lengkap
	items := []dto.CartItemResponse{}
	for _, item := range cart.Items {
		var variantIDStr *string
		if item.VariantID != nil {
			id := item.VariantID.String()
			variantIDStr = &id
		}

		items = append(items, dto.CartItemResponse{
			ID:          item.ID.String(),
			ProductID:   item.ProductID.String(),
			VariantID:   variantIDStr,
			ProductName: item.ProductName,
			ImageURL:    item.ImageURL,
			Price:       item.Price,
			Quantity:    item.Quantity,
			IsChecked:   item.IsChecked,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"cart": gin.H{
			"id":     cart.ID.String(),
			"userId": cart.UserID.String(),
			"items":  items,
		},
	})
}

// POST /api/cart
func (h *CartHandler) AddToCart(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	var variantIDStr string
	if req.VariantID != nil {
		variantIDStr = req.VariantID.String()
	}

	productResp, err := h.GRPCClient.GetProductSnapshot(req.ProductID.String(), variantIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch product info", "error": err.Error()})
		return
	}

	snapshot := dto.ProductSnapshot{
		Name:     productResp.Name,
		ImageURL: productResp.ImageUrl,
		Price:    productResp.Price,
		Stock:    int(productResp.Stock),
	}

	if err := h.Service.AddToCart(userID, req, snapshot); err != nil {
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

	var variantIDStr string
	if req.VariantID != nil {
		variantIDStr = req.VariantID.String()
	}

	productResp, err := h.GRPCClient.GetProductSnapshot(req.ProductID.String(), variantIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch product info"})
		return
	}

	snapshot := dto.ProductSnapshot{
		Name:     productResp.Name,
		ImageURL: productResp.ImageUrl,
		Price:    productResp.Price,
		Stock:    int(productResp.Stock),
	}

	if err := h.Service.UpdateCartItem(itemID, req, snapshot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
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
