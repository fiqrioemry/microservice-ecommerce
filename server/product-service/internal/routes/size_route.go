package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SizeRoutes(r *gin.Engine, handler *handlers.SizeHandler) {
	// Public
	r.GET("/api/sizes", handler.GetAll)

	// Admin
	admin := r.Group("/api/admin/sizes")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.POST("", handler.Create)
	admin.PUT("/:id", handler.Update)
	admin.DELETE("/:id", handler.Delete)
}
