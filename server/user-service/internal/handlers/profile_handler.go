package handlers

import (
	"net/http"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/services"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	Service services.ProfileServiceInterface
}

func NewProfileHandler(service services.ProfileServiceInterface) *ProfileHandler {
	return &ProfileHandler{Service: service}
}

func (ctrl *ProfileHandler) GetProfile(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	user, err := ctrl.Service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": user.Profile})
}

func (ctrl *ProfileHandler) UpdateProfile(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.ProfileRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}
	file, _ := c.FormFile("avatar")

	if err := ctrl.Service.UpdateProfile(userID, req, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update profile", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
