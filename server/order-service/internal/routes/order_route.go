package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/handlers"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine, handler *handlers.OrderHandler) {
	order := router.Group("/api/orders")
	order.Use(middleware.AuthRequired())

	// customer
	order.POST("", handler.Checkout)
	order.GET("", handler.GetUserOrders)
	order.POST("/check-shipping", handler.CheckShippingCost)

	// admin
	shipment := router.Group("/api/shipments")
	shipment.Use(middleware.AdminOnly())
	shipment.POST("", handler.CreateShipment)
	shipment.GET(":orderId", handler.GetShipment)
	shipment.PUT(":orderId", handler.UpdateShipmentStatus)

	admin := router.Group("/api/admin/orders")
	admin.Use(middleware.AdminOnly())
	admin.GET("", handler.GetAllOrders)

	router.POST("/api/payments/midtrans-notify", handler.HandleMidtransNotification)
}
