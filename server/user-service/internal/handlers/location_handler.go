package handlers

import (
	"net/http"
	"strconv"

	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/services"
	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	Service services.LocationService
}

func NewLocationHandler(service services.LocationService) *LocationHandler {
	return &LocationHandler{Service: service}
}

func (h *LocationHandler) GetProvinces(c *gin.Context) {
	data, err := h.Service.GetAllProvinces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get provinces"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *LocationHandler) GetCitiesByProvince(c *gin.Context) {
	provinceIDStr := c.Param("id")
	provinceID, err := strconv.Atoi(provinceIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid province ID"})
		return
	}

	data, err := h.Service.GetCitiesByProvince(uint(provinceID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get cities"})
		return
	}
	c.JSON(http.StatusOK, data)
}
