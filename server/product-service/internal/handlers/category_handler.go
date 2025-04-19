package handlers

import (
	"net/http"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	Service services.CategoryServiceInterface
}

func NewCategoryHandler(service services.CategoryServiceInterface) *CategoryHandler {
	return &CategoryHandler{Service: service}
}
func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func (h *CategoryHandler) Create(c *gin.Context) {
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "All fields are required"})
		return
	}

	imageURL := ""
	fileHeader, _ := c.FormFile("image")
	if fileHeader != nil {
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

		imageURL, err = utils.UploadToCloudinary(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Upload failed"})
			return
		}
	}

	req := dto.CategoryRequest{
		Name:  name,
		Slug:  utils.GenerateSlug(name),
		Image: imageURL,
	}

	if err := h.Service.Create(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Category created"})
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid category ID"})
		return
	}

	category, err := h.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		return
	}

	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "All fields are required"})
		return
	}

	imageURL := category.Image
	fileHeader, _ := c.FormFile("image")
	if fileHeader != nil {
		if err := utils.ValidateImageFile(fileHeader); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if imageURL != "" {
			_ = utils.DeleteFromCloudinary(imageURL)
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

	req := dto.CategoryRequest{
		Name:  name,
		Slug:  utils.GenerateSlug(name),
		Image: imageURL,
	}

	if err := h.Service.Update(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category updated"})
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid category ID"})
		return
	}

	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}

func (h *CategoryHandler) CreateSubcategory(c *gin.Context) {
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

	if err := h.Service.CreateSubcategory(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Subcategory created"})
}

func (h *CategoryHandler) UpdateSubcategory(c *gin.Context) {
	id, err := uuid.Parse(c.Param("subId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid subcategory ID"})
		return
	}

	sub, err := h.Service.GetSubcategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Subcategory not found"})
		return
	}

	name := c.PostForm("name")
	categoryID, _ := uuid.Parse(c.PostForm("categoryId"))
	imageURL := sub.Image

	fileHeader, _ := c.FormFile("image")
	if fileHeader != nil {
		if err := utils.ValidateImageFile(fileHeader); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if imageURL != "" {
			_ = utils.DeleteFromCloudinary(imageURL)
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

	if err := h.Service.UpdateSubcategory(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subcategory updated"})
}

func (h *CategoryHandler) DeleteSubcategory(c *gin.Context) {
	id, err := uuid.Parse(c.Param("subId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid subcategory ID"})
		return
	}

	if err := h.Service.DeleteSubcategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subcategory deleted"})
}
