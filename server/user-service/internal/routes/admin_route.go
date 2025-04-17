package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine, handler *handlers.AuthHandler) {
	admin := router.Group("/api/admin")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	admin.GET("/users", handler.GetAllUsers)
	admin.GET("/user/:id", handler.GetUserByIDAdmin)
}
