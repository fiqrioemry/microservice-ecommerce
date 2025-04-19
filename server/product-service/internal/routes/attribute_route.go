package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AttributeRoutes(r *gin.Engine, h *handlers.AttributeHandler) {
	route := r.Group("/api/attributes")

	// Public route
	route.GET("", h.GetAll)

	// Admin-only routes
	admin := route.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.POST("", h.Create)
	admin.PUT("/:id", h.Update)
	admin.DELETE("/:id", h.Delete)

	admin.POST("/:id/values", h.AddValue)
	admin.PUT("/values/:valueId", h.UpdateValue)
	admin.DELETE("/values/:valueId", h.DeleteValue)
}
