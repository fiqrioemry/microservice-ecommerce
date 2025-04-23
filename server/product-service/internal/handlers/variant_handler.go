package handlers

import (
	"net/http"
	"strconv"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/services"
	"github.com/gin-gonic/gin"
)

type VariantHandler struct {
	Service services.VariantService
}

func NewVariantHandler(service services.VariantService) *VariantHandler {
	return &VariantHandler{Service: service}
}

func (h *VariantHandler) GetAllTypes(c *gin.Context) {
	types, err := h.Service.GetAllTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, types)
}

func (h *VariantHandler) CreateType(c *gin.Context) {
	var req dto.VariantTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.CreateType(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Variant type created"})
}

func (h *VariantHandler) UpdateType(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req dto.VariantTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.UpdateType(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Variant type updated"})
}

func (h *VariantHandler) DeleteType(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteType(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Variant type deleted"})
}

func (h *VariantHandler) AddValue(c *gin.Context) {
	var req dto.VariantValueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.AddValue(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Variant value added"})
}

func (h *VariantHandler) UpdateValue(c *gin.Context) {
	valueID, _ := strconv.Atoi(c.Param("valueId"))
	var body struct {
		Value string `json:"value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.UpdateValue(uint(valueID), body.Value); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Variant value updated"})
}

func (h *VariantHandler) DeleteValue(c *gin.Context) {
	valueID, _ := strconv.Atoi(c.Param("valueId"))
	if err := h.Service.DeleteValue(uint(valueID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Variant value deleted"})
}
