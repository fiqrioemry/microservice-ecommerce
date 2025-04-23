package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func BannerRoutes(r *gin.Engine, h *handlers.BannerHandler) {
	group := r.Group("/api/banners")
	group.GET("", h.GetAllBanners)
	group.GET("/:position", h.GetBanner)
	group.POST("", middleware.AuthRequired(), middleware.AdminOnly(), h.UploadBanner)
	group.PUT("/:id", middleware.AuthRequired(), middleware.AdminOnly(), h.UpdateBanner)
	group.DELETE("/:id", middleware.AuthRequired(), middleware.AdminOnly(), h.DeleteBanner)
}
