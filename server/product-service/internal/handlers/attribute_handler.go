package handlers

import (
	"net/http"
	"strconv"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/services"
	"github.com/gin-gonic/gin"
)

type AttributeHandler struct {
	Service services.AttributeService
}

func NewAttributeHandler(service services.AttributeService) *AttributeHandler {
	return &AttributeHandler{Service: service}
}

// GET /api/attributes
func (h *AttributeHandler) GetAll(c *gin.Context) {
	attrs, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, attrs)
}

// POST /api/attributes
func (h *AttributeHandler) Create(c *gin.Context) {
	var req dto.AttributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.Create(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Attribute created successfully"})
}

// PUT /api/attributes/:id
func (h *AttributeHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req dto.AttributeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.Update(uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Attribute updated successfully"})
}

// DELETE /api/attributes/:id
func (h *AttributeHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Attribute deleted successfully"})
}

// POST /api/attributes/:id/values
func (h *AttributeHandler) AddValue(c *gin.Context) {
	var req dto.AttributeValueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.AddValue(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Attribute value added"})
}

// PUT /api/attributes/values/:valueId
func (h *AttributeHandler) UpdateValue(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{"message": "Attribute value updated"})
}

// DELETE /api/attributes/values/:valueId
func (h *AttributeHandler) DeleteValue(c *gin.Context) {
	valueID, _ := strconv.Atoi(c.Param("valueId"))
	if err := h.Service.DeleteValue(uint(valueID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Attribute value deleted"})
}
