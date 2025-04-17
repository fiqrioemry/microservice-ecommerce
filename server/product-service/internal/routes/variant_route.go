package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func VariantRoutes(r *gin.Engine, h *handlers.ProductVariantHandler) {
	v := r.Group("/api/products/:productId/variants")
	v.GET("", h.GetByProduct)

	vAdmin := r.Group("/variants")
	vAdmin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	vAdmin.POST("", h.Create)
	vAdmin.PUT("/:id", h.Update)
	vAdmin.DELETE("/:id", h.Delete)
}
