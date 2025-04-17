package handlers

import (
	"net/http"
	"strconv"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductAttributeValueHandler struct {
	Service services.ProductAttributeValueService
}

func NewProductAttributeValueHandler(s services.ProductAttributeValueService) *ProductAttributeValueHandler {
	return &ProductAttributeValueHandler{Service: s}
}

func (h *ProductAttributeValueHandler) GetByProduct(c *gin.Context) {
	productID := c.Param("productId")
	data, err := h.Service.GetByProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get product attribute values"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"attributes": data})
}

func (h *ProductAttributeValueHandler) Add(c *gin.Context) {
	var req dto.AddProductAttributeValueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Add(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Attribute value linked to product"})
}

func (h *ProductAttributeValueHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
