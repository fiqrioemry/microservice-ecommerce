package handlers

import (
	"net/http"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SubcategoryHandler struct {
	Service services.SubcategoryServiceInterface
}

func NewSubcategoryHandler(service services.SubcategoryServiceInterface) *SubcategoryHandler {
	return &SubcategoryHandler{Service: service}
}

func (h *SubcategoryHandler) GetAll(c *gin.Context) {
	subcategories, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch subcategories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"subcategories": subcategories})
}

func (h *SubcategoryHandler) Create(c *gin.Context) {
	name := c.PostForm("name")
	categoryID, _ := uuid.Parse(c.PostForm("categoryId"))

	form, err := c.MultipartForm()
	if err != nil || form.File["image"] == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image is required"})
		return
	}

	fileHeader := form.File["image"][0]
	if err := utils.ValidateImageFile(fileHeader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to open image"})
		return
	}
	defer file.Close()

	imageURL, err := utils.UploadToCloudinary(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Upload failed"})
		return
	}

	req := dto.SubcategoryRequest{
		Name:       name,
		Slug:       utils.GenerateSlug(name),
		CategoryID: categoryID,
		Image:      imageURL,
	}

	if err := h.Service.Create(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Subcategory created successfully"})
}

func (h *SubcategoryHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid subcategory ID"})
		return
	}

	oldSubcat, err := h.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Subcategory not found"})
		return
	}

	name := c.PostForm("name")
	categoryID, _ := uuid.Parse(c.PostForm("categoryId"))

	imageURL := oldSubcat.Image

	fileHeader, _ := c.FormFile("image")
	if fileHeader != nil {
		if err := utils.ValidateImageFile(fileHeader); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if oldSubcat.Image != "" {
			_ = utils.DeleteFromCloudinary(oldSubcat.Image)
		}

		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to open image"})
			return
		}
		defer file.Close()

		imageURL, err = utils.UploadToCloudinary(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Upload failed"})
			return
		}
	}

	req := dto.SubcategoryRequest{
		Name:       name,
		Slug:       utils.GenerateSlug(name),
		CategoryID: categoryID,
		Image:      imageURL,
	}

	if err := h.Service.Update(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subcategory updated successfully"})
}

func (h *SubcategoryHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid subcategory ID"})
		return
	}

	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete subcategory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subcategory deleted successfully"})
}
