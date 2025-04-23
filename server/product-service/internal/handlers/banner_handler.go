package handlers

import (
	"mime/multipart"
	"net/http"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BannerHandler struct {
	Service services.BannerService
}

func NewBannerHandler(s services.BannerService) *BannerHandler {
	return &BannerHandler{Service: s}
}

func (h *BannerHandler) UploadBanner(c *gin.Context) {
	var req dto.BannerRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "position is required"})
		return
	}

	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image is required"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot open file"})
		return
	}
	defer file.Close()

	if err := h.Service.Create(req, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload banner"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "banner uploaded successfully"})
}

func (h *BannerHandler) GetBanner(c *gin.Context) {
	position := c.Param("position")
	results, err := h.Service.Get(position)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch banner"})
		return
	}
	c.JSON(http.StatusOK, results)
}

func (h *BannerHandler) DeleteBanner(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete banner"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "banner deleted"})
}

func (h *BannerHandler) UpdateBanner(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var req dto.BannerRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "position is required and must be one of: top, side1, side2, bottom"})
		return
	}

	var file multipart.File
	fileHeader, err := c.FormFile("image")
	if err == nil {
		file, err = fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot open file"})
			return
		}
		defer file.Close()
	}

	if err := h.Service.Update(id, req, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update banner"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "banner updated"})
}

