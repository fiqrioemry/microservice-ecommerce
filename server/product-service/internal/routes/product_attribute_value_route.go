package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func ProductAttributeValueRoutes(r *gin.Engine, h *handlers.ProductAttributeValueHandler) {

	productAttr := r.Group("/api/products/:productId/attributes")
	productAttr.GET("", h.GetByProduct)
	admin := productAttr.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.POST("", h.Add)
	admin.DELETE("/:id", h.Delete)
}
