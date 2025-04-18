package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func ProductAttributeValueRoutes(r *gin.Engine, h *handlers.ProductAttributeValueHandler) {
	// Admin routes
	admin := r.Group("/api/admin/products")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.GET("/:id/attributes", h.GetByProduct)
	admin.POST("/:id/attributes", h.Add)
	admin.DELETE("/:id/attributes", h.Delete)
}
