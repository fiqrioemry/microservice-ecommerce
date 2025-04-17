package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SubcategoryRoutes(r *gin.Engine, handler *handlers.SubcategoryHandler) {
	route := r.Group("/api/subcategories")
	{
		// Public route
		route.GET("", handler.GetAll)

		// Admin-only routes
		route.Use(middleware.AuthRequired(), middleware.AdminOnly())
		{
			route.POST("", handler.Create)
			route.PUT("/:id", handler.Update)
			route.DELETE("/:id", handler.Delete)
		}
	}
}
