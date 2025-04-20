package handlers

import (
	"encoding/json"
	"net/http"

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

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid form data"})
		return
	}

	// === Upload product images ===
	productImages := form.File["images"]
	if len(productImages) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "At least 1 product image is required"})
		return
	}

	uploadedProductImages := []string{}
	for _, fileHeader := range productImages {
		if err := utils.ValidateImageFile(fileHeader); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer file.Close()

		uploadedURL, err := utils.UploadToCloudinary(file)
		if err == nil {
			uploadedProductImages = append(uploadedProductImages, uploadedURL)
		}
	}

	// === Upload variant images (optional) ===
	variantImages := form.File["variantImages"]
	uploadedVariantImages := make([]string, len(variantImages))
	for i, fileHeader := range variantImages {
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
			uploadedVariantImages[i] = url
		}
	}

	// === Parse product JSON data ===
	data := c.PostForm("data")
	var req dto.CreateFullProductRequest
	if err := json.Unmarshal([]byte(data), &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON in 'data' field"})
		return
	}

	for i := range req.Variants {
		if i < len(uploadedVariantImages) {
			req.Variants[i].ImageURL = uploadedVariantImages[i]
		}
	}

	if len(req.Variants) == 0 {
		req.Variants = append(req.Variants, dto.CreateVariantRequest{
			SKU:      utils.GenerateSlug(req.Name),
			Price:    0,
			Stock:    0,
			IsActive: true,
		})
	}

	if err := h.Service.CreateFullProduct(req, uploadedProductImages); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch products"})
		return
	}

	var response []dto.ProductResponse
	for _, p := range products {
		resp := dto.ProductResponse{
			ID:          p.ID.String(),
			Name:        p.Name,
			Slug:        p.Slug,
			Description: p.Description,
			IsFeatured:  p.IsFeatured,
			IsActive:    p.IsActive,
			Category:    p.Category,
			Subcategory: p.Subcategory,
		}

		for _, img := range p.ProductImage {
			resp.Images = append(resp.Images, img.URL)
		}

		response = append(response, resp)
	}

	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) GetProductBySlug(c *gin.Context) {
	slug := c.Param("slug")
	product, err := h.Service.GetBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	resp := dto.ProductResponse{
		ID:          product.ID.String(),
		Name:        product.Name,
		Slug:        product.Slug,
		Description: product.Description,
		IsFeatured:  product.IsFeatured,
		IsActive:    product.IsActive,
		Weight:      product.Weight,
		Length:      product.Length,
		Width:       product.Width,
		Height:      product.Height,
		Category:    product.Category,
		Subcategory: product.Subcategory,
	}

	for _, img := range product.ProductImage {
		resp.Images = append(resp.Images, img.URL)
	}

	for _, v := range product.ProductVariant {
		variant := dto.ProductVariantOutput{
			ID:       v.ID.String(),
			SKU:      v.SKU,
			Price:    v.Price,
			Stock:    v.Stock,
			IsActive: v.IsActive,
			ImageURL: v.ImageURL,
			Options:  make(map[string]string),
		}

		for _, opt := range v.VariantValues {
			variant.Options[opt.OptionValue.Type.Name] = opt.OptionValue.Value
		}
		resp.Variants = append(resp.Variants, variant)
	}

	for _, attr := range product.ProductAttributeValue {
		resp.Attributes = append(resp.Attributes, dto.AttributeOutput{
			Name:  attr.Attribute.Name,
			Value: attr.AttributeValue.Value,
		})
	}

	c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid form data"})
		return
	}

	imageFiles := form.File["images"]
	uploadedImageURLs := []string{}
	for _, fileHeader := range imageFiles {
		if err := utils.ValidateImageFile(fileHeader); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer file.Close()

		uploadedURL, err := utils.UploadToCloudinary(file)
		if err == nil {
			uploadedImageURLs = append(uploadedImageURLs, uploadedURL)
		}
	}

	data := c.PostForm("data")
	var req dto.UpdateProductRequest
	if err := json.Unmarshal([]byte(data), &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON in 'data' field"})
		return
	}

	if err := h.Service.UpdateWithImages(id, req, uploadedImageURLs); err != nil {
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
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
