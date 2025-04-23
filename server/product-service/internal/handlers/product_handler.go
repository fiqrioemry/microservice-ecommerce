package handlers

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

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

func (h *ProductHandler) SearchProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	minPrice, _ := strconv.ParseFloat(c.DefaultQuery("minPrice", "0"), 64)
	maxPrice, _ := strconv.ParseFloat(c.DefaultQuery("maxPrice", "0"), 64)

	sort := c.DefaultQuery("sort", "name:asc")
	parts := strings.Split(sort, ":")
	if len(parts) != 2 || (parts[1] != "asc" && parts[1] != "desc") {
		sort = "name:asc" // fallback jika format sort tidak valid
	}

	params := dto.SearchParams{
		Query:       c.Query("q"),
		Category:    c.Query("category"),
		Subcategory: c.Query("subcategory"),
		InStock:     c.Query("stock") == "true",
		Sort:        sort,
		Page:        page,
		Limit:       limit,
		MinPrice:    minPrice,
		MaxPrice:    maxPrice,
	}

	products, total, err := h.Service.Search(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Search failed"})
		return
	}
	var response []dto.ProductMinimal
	for _, p := range products {
		item := dto.ProductMinimal{
			ID:          p.ID.String(),
			Name:        p.Name,
			Slug:        p.Slug,
			Price:       0,
			Description: p.Description,
			IsFeatured:  p.IsFeatured,
			IsActive:    p.IsActive,
			Category: dto.CategoryMinimal{
				ID:   p.Category.ID,
				Name: p.Category.Name,
				Slug: p.Category.Slug,
			},
		}

		if p.Subcategory != nil {
			item.Subcategory = &dto.CategoryMinimal{
				ID:   p.Subcategory.ID,
				Name: p.Subcategory.Name,
				Slug: p.Subcategory.Slug,
			}
		}

		for _, img := range p.ProductImage {
			item.Images = append(item.Images, img.URL)
		}

		if len(p.ProductVariant) > 0 {
			item.Price = p.ProductVariant[0].Price
		}

		response = append(response, item)
	}

	if response == nil {
		response = []dto.ProductMinimal{}
	}

	c.JSON(http.StatusOK, gin.H{
		"page":    params.Page,
		"limit":   params.Limit,
		"total":   total,
		"pages":   (total + int64(params.Limit) - 1) / int64(params.Limit),
		"results": response,
	})
}


func (h *ProductHandler) CreateProduct(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid form data"})
		return
	}

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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	products, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch products"})
		return
	}
	total := len(products)
	end := min(offset + limit, total)

	var response []dto.ProductMinimal
	for _, p := range products[offset:end] {
		item := dto.ProductMinimal{
			ID:          p.ID.String(),
			Name:        p.Name,
			Slug:        p.Slug,
			Price:       0,
			Description: p.Description,
			IsFeatured:  p.IsFeatured,
			IsActive:    p.IsActive,
			Category: dto.CategoryMinimal{
				ID:   p.Category.ID,
				Name: p.Category.Name,
				Slug: p.Category.Slug,
			},
		}

		if p.Subcategory != nil {
			item.Subcategory = &dto.CategoryMinimal{
				ID:   p.Subcategory.ID,
				Name: p.Subcategory.Name,
				Slug: p.Subcategory.Slug,
			}
		}

		for _, img := range p.ProductImage {
			item.Images = append(item.Images, img.URL)
		}

		if len(p.ProductVariant) > 0 {
			item.Price = p.ProductVariant[0].Price
		}

		response = append(response, item)
	}

	c.JSON(http.StatusOK, gin.H{
		"page":    page,
		"limit":   limit,
		"total":   total,
		"pages":   (total + limit - 1) / limit,
		"results": response,
	})
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
		Discount:    product.Discount,
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
			Sold:     v.Sold,
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

func (h *ProductHandler) DeleteVariantProduct(c *gin.Context) {
	id := c.Param("id")

	variantId, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
		return
	}

	if err := h.Service.DeleteVariantProduct(variantId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product Variant deleted successfully"})
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

func (h *ProductHandler) UploadLocalImage(c *gin.Context) {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image file is required"})
		return
	}
	if err := utils.ValidateImageFile(fileHeader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to open file"})
		return
	}
	defer file.Close()

	path, err := utils.UploadToLocal(file, fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"localPath": path})
}

func (h *ProductHandler) DownloadImage(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
		return
	}

	product, err := h.Service.GetProductByID(id)
	if err != nil || len(product.ProductImage) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Image not found"})
		return
	}

	firstImgPath := product.ProductImage[0].URL
	c.FileAttachment(firstImgPath, filepath.Base(firstImgPath))
}
