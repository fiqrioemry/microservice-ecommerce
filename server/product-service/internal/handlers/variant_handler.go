package handlers

import (
	"net/http"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductVariantHandler struct {
	Service services.ProductVariantService
}

func NewVariantHandler(s services.ProductVariantService) *ProductVariantHandler {
	return &ProductVariantHandler{Service: s}
}

func (h *ProductVariantHandler) GetByProduct(c *gin.Context) {
	slug := c.Param("slug")
	data, err := h.Service.GetByProduct(slug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get product variants"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"variants": data})
}

func (h *ProductVariantHandler) Create(c *gin.Context) {
	var req dto.CreateProductVariantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Create(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Variant created"})
}

func (h *ProductVariantHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateProductVariantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Update(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Variant updated"})
}

func (h *ProductVariantHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Variant deleted"})
}
