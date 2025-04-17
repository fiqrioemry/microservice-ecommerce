package handlers

import (
	"net/http"
	"strconv"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/services"
	"github.com/gin-gonic/gin"
)

type SizeHandler struct {
	Service services.SizeService
}

func NewSizeHandler(service services.SizeService) *SizeHandler {
	return &SizeHandler{Service: service}
}

func (h *SizeHandler) GetAll(c *gin.Context) {
	sizes, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch sizes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sizes": sizes})
}

func (h *SizeHandler) Create(c *gin.Context) {
	var req dto.CreateSizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Create(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create size"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Size created"})
}

func (h *SizeHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var req dto.UpdateSizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Update(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update size"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Size updated"})
}

func (h *SizeHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete size"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Size deleted"})
}
