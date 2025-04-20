package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, h *handlers.ProductHandler) {
	product := r.Group("/api/products")

	//  Admin routes
	admin := product.Group("/admin")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.POST("", h.CreateProduct)
	admin.PUT("/:id", h.UpdateProduct)
	admin.DELETE("/:id", h.DeleteProduct)
	admin.POST("/upload-local", h.UploadLocalImage)
	admin.GET("/:id/download", h.DownloadImage)

	//  Public routes
	product.GET("", h.GetAllProducts)
	product.GET("/:slug", h.GetProductBySlug)
}
