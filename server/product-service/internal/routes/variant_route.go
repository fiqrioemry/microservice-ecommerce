package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func VariantRoutes(r *gin.Engine, h *handlers.VariantHandler) {
	route := r.Group("/api/variants")

	route.GET("", h.GetAllTypes)

	admin := route.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.POST("", h.CreateType)
	admin.PUT("/:id", h.UpdateType)
	admin.DELETE("/:id", h.DeleteType)

	admin.POST("/:id/values", h.AddValue)
	admin.PUT("/values/:valueId", h.UpdateValue)
	admin.DELETE("/values/:valueId", h.DeleteValue)

}
