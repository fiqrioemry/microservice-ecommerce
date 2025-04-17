package handlers

import (
	"net/http"
	"strconv"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/services"

	"github.com/gin-gonic/gin"
)

type ColorHandler struct {
	Service services.ColorServiceInterface
}

func NewColorHandler(service services.ColorServiceInterface) *ColorHandler {
	return &ColorHandler{Service: service}
}

func (h *ColorHandler) GetAll(c *gin.Context) {
	colors, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch colors"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"colors": colors})
}

func (h *ColorHandler) Create(c *gin.Context) {
	var req dto.ColorRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}
	if err := h.Service.Create(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Color created successfully"})
}

func (h *ColorHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid color ID"})
		return
	}

	var req dto.ColorRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.Service.Update(uint(id), req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Color updated successfully"})
}

func (h *ColorHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid color ID"})
		return
	}

	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete color"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Color deleted successfully"})
}
