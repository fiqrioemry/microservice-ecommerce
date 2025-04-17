package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AttributeRoutes(r *gin.Engine, h *handlers.AttributeHandler) {
	attr := r.Group("/api/attributes")
	{
		attr.GET("", h.GetAll)
		attr.GET("/:id/values", h.GetValues)

		attr.Use(middleware.AuthRequired(), middleware.AdminOnly())
		{
			attr.POST("", h.Create)
			attr.PUT("/:id", h.Update)
			attr.DELETE("/:id", h.Delete)
			attr.POST("/values", h.CreateValue)

		}
	}
}
