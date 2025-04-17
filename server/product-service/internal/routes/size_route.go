package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SizeRoutes(r *gin.Engine, handler *handlers.SizeHandler) {
	sizes := r.Group("/api/sizes")
	sizes.GET("", handler.GetAll)

	// Admin-only routes
	sizes.Use(middleware.AuthRequired(), middleware.AdminOnly())
	sizes.POST("", handler.Create)
	sizes.PUT("/:id", handler.Update)
	sizes.DELETE("/:id", handler.Delete)
}
