package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/handlers"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, controller *handlers.AuthController) {
	auth := router.Group("/api/auth")
	auth.POST("/login", controller.Login)
	auth.POST("/register", controller.Register)
	auth.POST("/reset-password", controller.ResetPassword)
	auth.POST("/forgot-password", controller.ForgotPassword)
	auth.GET("/me", middleware.AuthRequired(), controller.Me)
	auth.PUT("/change-password", middleware.AuthRequired(), controller.ChangePassword)

	// Admin routes
	admin := auth.Group("/admin")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.GET("/users", controller.GetAllUsers)
	admin.GET("/user/:id", controller.GetUserByIDAdmin)
}
