package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine, handler *handlers.CategoryHandler) {
	// Public
	r.GET("/api/categories", handler.GetAll)

	// Admin Category
	admin := r.Group("/api/admin/categories")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())

	admin.POST("", handler.Create)
	admin.PUT("/:id", handler.Update)
	admin.DELETE("/:id", handler.Delete)

	// Admin Subcategory
	admin.POST("/:id/subcategories", handler.CreateSubcategory)
	admin.PUT("/subcategories/:subId", handler.UpdateSubcategory)
	admin.DELETE("/subcategories/:subId", handler.DeleteSubcategory)
}
