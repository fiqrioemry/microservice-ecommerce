package handlers

import (
	"net/http"

	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/services"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderHandler struct {
	Service services.OrderServiceInterface
}

func NewOrderHandler(service services.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{Service: service}
}

// GET /api/orders/:id
func (h *OrderHandler) GetOrder(c *gin.Context) {
	orderID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	order, err := h.Service.GetOrderDetail(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Checkout(c *gin.Context) {
	userIDStr := utils.MustGetUserID(c)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req dto.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	addressID, err := uuid.Parse(req.AddressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID"})
		return
	}

	cart, err := h.Service.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve cart", "error": err.Error()})
		return
	}

	order, err := h.Service.CreateOrderFromCart(userID, addressID, cart, req.Note, req.ShippingCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	snapURL, err := h.Service.GenerateSnapTransaction(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Snap transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"order":    order,
		"snap_url": snapURL,
	})
}

func (h *OrderHandler) HandleMidtransNotification(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	orderIDStr, _ := payload["order_id"].(string)
	transactionStatus, _ := payload["transaction_status"].(string)
	paymentType, _ := payload["payment_type"].(string)
	statusCode, _ := payload["status_code"].(string)

	// ignore duplicate/invalid calls
	if orderIDStr == "" || transactionStatus == "" || statusCode != "200" {
		c.JSON(http.StatusOK, gin.H{"message": "ignored"})
		return
	}

	if err := h.Service.UpdatePaymentStatus(orderIDStr, transactionStatus, paymentType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment status updated"})
}

func (h *OrderHandler) CheckShippingCost(c *gin.Context) {
	var req dto.ShippingCostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	const originID = 278

	cost, err := utils.FetchShippingCost(originID, req.DestinationID, req.Weight, req.Courier)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error":  "Failed to fetch shipping cost",
			"detail": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"origin_id":     originID,
		"shipping_cost": cost,
	})
}
