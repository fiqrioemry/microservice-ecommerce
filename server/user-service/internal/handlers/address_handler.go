package handlers

import (
	"net/http"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/services"
	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	Service services.AddressServiceInterface
}

func NewAddressHandler(service services.AddressServiceInterface) *AddressHandler {
	return &AddressHandler{Service: service}
}

func (h *AddressHandler) AddAddress(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.AddressRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	err := h.Service.AddAddressWithLocation(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to add address"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "address added successfully"})
}

func (h *AddressHandler) UpdateAddress(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	addressID := c.Param("id")

	var req dto.AddressRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	err := h.Service.UpdateAddressWithLocation(userID, addressID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "address updated successfully"})
}

func (h *AddressHandler) GetAddresses(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	addresses, err := h.Service.GetAddresses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get addresses"})
		return
	}

	c.JSON(http.StatusOK, addresses)
}

func (ctrl *AddressHandler) DeleteAddress(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	addressID := c.Param("id")

	if err := ctrl.Service.DeleteAddress(userID, addressID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete address", "error": err.Error()})
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
