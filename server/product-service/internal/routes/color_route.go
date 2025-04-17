package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func ColorRoutes(r *gin.Engine, handler *handlers.ColorHandler) {
	color := r.Group("/api/colors")
	{
		color.GET("", handler.GetAll)
		color.Use(middleware.AuthRequired(), middleware.AdminOnly())
		{
			color.POST("", handler.Create)
			color.PUT("/:id", handler.Update)
			color.DELETE("/:id", handler.Delete)
		}
	}
}
