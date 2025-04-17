package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, handler *handlers.ProductHandler) {
	product := r.Group("/api/products")
	{
		product.GET("", handler.GetAllProducts)
		product.GET("/:slug", handler.GetProductBySlug)

		product.Use(middleware.AuthRequired(), middleware.AdminOnly())
		{
			product.POST("", handler.CreateProduct)
			product.PUT("/:id", handler.UpdateProduct)
			product.DELETE("/:id", handler.DeleteProduct)
		}
	}
}
