package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func ProductVariantRoutes(r *gin.Engine, h *handlers.ProductVariantHandler) {

	// Admin routes
	admin := r.Group("/api/admin/products/:id/variants")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.GET("", h.GetByProduct)
	admin.POST("", h.Create)
	admin.PUT("/:id", h.Update)
	admin.DELETE("/:id", h.Delete)
}
