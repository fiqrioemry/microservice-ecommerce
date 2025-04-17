package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, handler *handlers.AuthHandler) {
	auth := router.Group("/api/auth")
	auth.POST("/login", handler.Login)
	auth.POST("/logout", handler.Logout)
	auth.POST("/register", handler.Register)
	auth.POST("/reset-password", handler.ResetPassword)
	auth.POST("/forgot-password", handler.ForgotPassword)
	auth.GET("/me", middleware.AuthRequired(), handler.Me)
	auth.PUT("/change-password", middleware.AuthRequired(), handler.ChangePassword)

	// Admin routes
	admin := auth.Group("/admin")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.GET("/user", handler.GetAllUsers)
	admin.GET("/user/:id", handler.GetUserByIDAdmin)
}
