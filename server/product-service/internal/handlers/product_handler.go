package handlers

import (
	"net/http"
	"strconv"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	Service services.ProductServiceInterface
}

func NewProductHandler(service services.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch products"})
		return
	}

	var responses []dto.ProductResponse
	for _, p := range products {
		var imageURLs []string
		for _, img := range p.ProductImage {
			imageURLs = append(imageURLs, img.URL)
		}

		responses = append(responses, dto.ProductResponse{
			ID:          p.ID.String(),
			Name:        p.Name,
			Slug:        p.Slug,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			Sold:        p.Sold,
			IsFeatured:  p.IsFeatured,
			IsActive:    p.IsActive,
			Images:      imageURLs,
			Category:    p.Category,
			Subcategory: p.Subcategory,
		})
	}

	c.JSON(http.StatusOK, gin.H{"products": responses})
}

func (h *ProductHandler) GetProductBySlug(c *gin.Context) {
	slug := c.Param("slug")
	product, err := h.Service.GetBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	var imageURLs []string
	for _, img := range product.ProductImage {
		imageURLs = append(imageURLs, img.URL)
	}

	response := dto.ProductResponse{
		ID:          product.ID.String(),
		Name:        product.Name,
		Slug:        product.Slug,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Sold:        product.Sold,
		IsFeatured:  product.IsFeatured,
		IsActive:    product.IsActive,
		Images:      imageURLs,
		Category:    product.Category,
		Subcategory: product.Subcategory,
	}

	c.JSON(http.StatusOK, gin.H{"product": response})
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	categoryID, _ := uuid.Parse(c.PostForm("categoryId"))
	var subcategoryID *uuid.UUID
	if subID := c.PostForm("subcategoryId"); subID != "" {
		id, err := uuid.Parse(subID)
		if err == nil {
			subcategoryID = &id
		}
	}

	name := c.PostForm("name")

	req := dto.CreateProductRequest{
		Name:          name,
		Slug:          utils.GenerateSlug(name),
		Description:   c.PostForm("description"),
		Price:         price,
		Stock:         stock,
		CategoryID:    categoryID,
		SubcategoryID: subcategoryID,
		IsFeatured:    c.PostForm("isFeatured") == "true",
	}

	form, err := c.MultipartForm()
	if err != nil || form.File["images"] == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Product image(s) required"})
		return
	}

	imageURLs := []string{}
	for _, fileHeader := range form.File["images"] {
		if err := utils.ValidateImageFile(fileHeader); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer file.Close()

		url, err := utils.UploadToCloudinary(file)
		if err == nil {
			imageURLs = append(imageURLs, url)
		}
	}

	if len(imageURLs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "At least 1 valid image required"})
		return
	}

	if err := h.Service.CreateWithImages(req, imageURLs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
		return
	}

	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	categoryID, _ := uuid.Parse(c.PostForm("categoryId"))
	var subcategoryID *uuid.UUID
	if subID := c.PostForm("subcategoryId"); subID != "" {
		id, err := uuid.Parse(subID)
		if err == nil {
			subcategoryID = &id
		}
	}

	name := c.PostForm("name")

	req := dto.UpdateProductRequest{
		Name:          name,
		Slug:          utils.GenerateSlug(name),
		Description:   c.PostForm("description"),
		Price:         price,
		Stock:         stock,
		CategoryID:    categoryID,
		SubcategoryID: subcategoryID,
		IsFeatured:    c.PostForm("isFeatured") == "true",
	}

	var imageURLs []string
	form, err := c.MultipartForm()
	if err == nil && form.File["images"] != nil {
		for _, fileHeader := range form.File["images"] {
			if err := utils.ValidateImageFile(fileHeader); err != nil {
				continue
			}

			file, err := fileHeader.Open()
			if err != nil {
				continue
			}
			defer file.Close()

			url, err := utils.UploadToCloudinary(file)
			if err == nil {
				imageURLs = append(imageURLs, url)
			}
		}
	}

	if err := h.Service.UpdateWithImages(id, req, imageURLs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
		return
	}

	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
