package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func VariantRoutes(r *gin.Engine, h *handlers.ProductVariantHandler) {
	// Public routes
	r.GET("/api/products/:productId/variants", h.GetByProduct)

	// Admin routes
	admin := r.Group("/api/admin/variants")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.POST("", h.Create)
	admin.PUT("/:id", h.Update)
	admin.DELETE("/:id", h.Delete)
}
