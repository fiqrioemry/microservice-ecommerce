package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func ColorRoutes(r *gin.Engine, handler *handlers.ColorHandler) {
	// Public
	r.GET("/api/colors", handler.GetAll)

	// Admin only
	admin := r.Group("/api/admin/colors")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.POST("", handler.Create)
	admin.PUT("/:id", handler.Update)
	admin.DELETE("/:id", handler.Delete)
}
