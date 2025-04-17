package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, h *handlers.ProductHandler) {
	product := r.Group("/api/products")
	{
		product.GET("", h.GetAllProducts)
		product.GET("/:slug", h.GetProductBySlug)

		product.Use(middleware.AuthRequired(), middleware.AdminOnly())
		{
			product.POST("", h.CreateProduct)
			product.PUT("/:id", h.UpdateProduct)
			product.DELETE("/:id", h.DeleteProduct)
		}
	}
}
