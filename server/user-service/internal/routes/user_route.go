package routes

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/handlers"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, controller *handlers.ProfileController, addHandler *handlers.AddressController) {
	user := router.Group("/api/user")
	user.Use(middleware.AuthRequired())

	// Profile
	user.GET("/profile", handler.GetProfile)
	user.PUT("/profile", handler.UpdateProfile)

	// Address
	user.GET("/addresses", addHandler.GetAddresses)
	user.POST("/addresses", addHandler.AddAddress)
	user.PUT("/addresses/:id", addHandler.UpdateAddress)
	user.DELETE("/addresses/:id", addHandler.DeleteAddress)
	user.PUT("/addresses/:id/set-main", addHandler.SetMainAddress)
}
