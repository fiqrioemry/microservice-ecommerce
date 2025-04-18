package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/handlers"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine, h *handlers.CartHandler) {
	cart := r.Group("/api/cart")
	cart.Use(middleware.AuthRequired())

	cart.GET("", h.GetCart)
	cart.POST("", h.AddToCart)
	cart.PUT("/items/:id", h.UpdateCartItem)
	cart.DELETE("/items/:id", h.RemoveCartItem)
	cart.DELETE("", h.ClearCart)
}
