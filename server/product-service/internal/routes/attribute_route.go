package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AttributeRoutes(r *gin.Engine, h *handlers.AttributeHandler) {
	// Admin
	admin := r.Group("/api/admin/attributes")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.GET("", h.GetAll)
	admin.GET("/:id/values", h.GetValues)
	admin.POST("", h.Create)
	admin.PUT("/:id", h.Update)
	admin.DELETE("/:id", h.Delete)
	admin.POST("/values", h.CreateValue)
}
