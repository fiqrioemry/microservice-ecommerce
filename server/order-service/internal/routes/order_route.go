package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/handlers"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine, handler *handlers.OrderHandler) {
	order := router.Group("/api/orders")
	order.Use(middleware.AuthRequired())

	order.GET("/:id", handler.GetOrder)
	order.POST("", handler.Checkout)
	order.POST("/check-shipping", handler.CheckShippingCost)

	router.POST("/api/payments/midtrans-notify", handler.HandleMidtransNotification)
}
