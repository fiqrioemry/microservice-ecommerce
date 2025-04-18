package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func ProductAttributeValueRoutes(r *gin.Engine, h *handlers.ProductAttributeValueHandler) {
	// Public routes
	r.GET("/api/products/:slug/attributes", h.GetByProduct)

	// Admin routes
	admin := r.Group("/api/admin/product-attributes")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.POST("", h.Add)
	admin.DELETE("/:id", h.Delete)
}
