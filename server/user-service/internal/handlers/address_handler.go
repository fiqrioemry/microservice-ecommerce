package handlers

import (
	"net/http"

	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/services"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	Service services.AddressServiceInterface
}

func NewAddressHandler(service services.AddressServiceInterface) *AddressHandler {
	return &AddressHandler{Service: service}
}

func (ctrl *AddressHandler) GetAddresses(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	addresses, err := ctrl.Service.GetAddresses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get addresses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"addresses": addresses})
}

func (ctrl *AddressHandler) AddAddress(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.AddressRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := ctrl.Service.AddAddress(userID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add address"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Address added successfully"})
}

func (ctrl *AddressHandler) UpdateAddress(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	addressID := c.Param("id")

	var req dto.AddressRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := ctrl.Service.UpdateAddress(userID, addressID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address updated successfully"})
}

func (ctrl *AddressHandler) DeleteAddress(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	addressID := c.Param("id")

	if err := ctrl.Service.DeleteAddress(userID, addressID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})
}

func (ctrl *AddressHandler) SetMainAddress(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	addressID := c.Param("id")

	if err := ctrl.Service.SetMainAddress(userID, addressID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to set main address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Main address updated"})
}
