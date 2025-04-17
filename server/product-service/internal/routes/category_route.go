package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine, handler *handlers.CategoryHandler) {
	category := r.Group("/api/categories")
	{
		// Public
		category.GET("", handler.GetAll)

		// Admin only
		category.Use(middleware.AuthRequired(), middleware.AdminOnly())
		{
			category.POST("", handler.Create)
			category.PUT("/:id", handler.Update)
			category.DELETE("/:id", handler.Delete)
		}
	}
}
