package controllers

import (
	"net/http"
	"user-service/dto"
	"user-service/services"
	"user-service/utils"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	Service services.ProfileServiceInterface
}

func NewProfileController(service services.ProfileServiceInterface) *ProfileController {
	return &ProfileController{Service: service}
}

func (ctrl *ProfileController) GetProfile(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	user, err := ctrl.Service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": user.Profile})
}

func (ctrl *ProfileController) UpdateProfile(c *gin.Context) {
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
