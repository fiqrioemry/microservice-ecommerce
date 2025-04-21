package services

import (
	"errors"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/repositories"
	"github.com/google/uuid"
)

type ProductServiceInterface interface {
	Delete(id uuid.UUID) error
	GetAll() ([]models.Product, error)
	GetBySlug(slug string) (*models.Product, error)
	GetProductByID(id uuid.UUID) (*models.Product, error)
	CreateFullProduct(req dto.CreateFullProductRequest, imageURLs []string) error
	UpdateWithImages(id uuid.UUID, req dto.UpdateProductRequest, imageURLs []string) error
	Search(params dto.SearchParams) ([]models.Product, int64, error)
}

type ProductService struct {
	Repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductServiceInterface {
	return &ProductService{Repo: repo}
}

func (s *ProductService) CreateFullProduct(req dto.CreateFullProductRequest, imageURLs []string) error {
	product := models.Product{
		Name:          req.Name,
		Slug:          utils.GenerateSlug(req.Name),
		Description:   req.Description,
		CategoryID:    req.CategoryID,
		SubcategoryID: req.SubcategoryID,
		IsFeatured:    req.IsFeatured,
		IsActive:      true,
		Weight:        req.Weight,
		Length:        req.Length,
		Width:         req.Width,
		Height:        req.Height,
		Discount:      req.Discount,
	}
	if err := s.Repo.Create(&product); err != nil {
		return err
	}

	var images []models.ProductImage
	for i, url := range imageURLs {
		images = append(images, models.ProductImage{
			ProductID: product.ID,
			URL:       url,
			IsPrimary: i == 0,
		})
	}
	if err := s.Repo.SaveImages(images); err != nil {
		return err
	}

	for _, v := range req.Variants {
		variant := models.ProductVariant{
			ID:        uuid.New(),
			ProductID: product.ID,
			SKU:       v.SKU,
			Price:     v.Price,
			Stock:     v.Stock,
			Sold:      v.Sold,
			IsActive:  v.IsActive,
			ImageURL:  v.ImageURL,
		}
		if err := s.Repo.CreateVariant(&variant); err != nil {
			return err
		}

		for typeName, value := range v.Options {
			optionValue, err := s.Repo.FindVariantOptionValue(typeName, value)
			if err != nil {
				return err
			}
			variantOption := models.ProductVariantOption{
				ProductVariantID: variant.ID,
				OptionValueID:    optionValue.ID,
			}
			if err := s.Repo.CreateVariantOption(&variantOption); err != nil {
				return err
			}
		}
	}

	for _, attr := range req.Attributes {
		pav := models.ProductAttributeValue{
			ProductID:        product.ID,
			AttributeID:      attr.AttributeID,
			AttributeValueID: attr.AttributeValueID,
		}
		if err := s.Repo.CreateProductAttributeValue(&pav); err != nil {
			return err
		}
	}

	return nil
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.Repo.FindAllWithPreload()
}

func (s *ProductService) GetBySlug(slug string) (*models.Product, error) {
	return s.Repo.FindBySlug(slug)
}

func (s *ProductService) UpdateWithImages(id uuid.UUID, req dto.UpdateProductRequest, imageURLs []string) error {
	product, err := s.Repo.FindByID(id)
	if err != nil {
		return errors.New("product not found")
	}

	product.Name = req.Name
	product.Slug = utils.GenerateSlug(req.Name)
	product.Description = req.Description
	product.CategoryID = req.CategoryID
	product.SubcategoryID = req.SubcategoryID
	product.IsFeatured = req.IsFeatured
	product.Weight = req.Weight
	product.Length = req.Length
	product.Width = req.Width
	product.Height = req.Height
	product.Discount = req.Discount

	if err := s.Repo.Update(product); err != nil {
		return err
	}

	if len(imageURLs) > 0 {
		oldImages, _ := s.Repo.FindImagesByProductID(product.ID)
		for _, img := range oldImages {
			_ = utils.DeleteFromCloudinary(img.URL)
		}
		_ = s.Repo.DeleteImagesByProductID(product.ID)

		var newImages []models.ProductImage
		for i, url := range imageURLs {
			newImages = append(newImages, models.ProductImage{
				ProductID: product.ID,
				URL:       url,
				IsPrimary: i == 0,
			})
		}
		return s.Repo.SaveImages(newImages)
	}

	variants, _ := s.Repo.FindVariantsByProductID(product.ID)
	for _, v := range variants {
		if v.ImageURL != "" {
			_ = utils.DeleteFromCloudinary(v.ImageURL)
		}
	}

	return nil
}

func (s *ProductService) Delete(id uuid.UUID) error {
	images, _ := s.Repo.FindImagesByProductID(id)
	for _, img := range images {
		_ = utils.DeleteFromCloudinary(img.URL)
	}
	_ = s.Repo.DeleteImagesByProductID(id)

	variants, _ := s.Repo.FindVariantsByProductID(id)
	for _, v := range variants {
		if v.ImageURL != "" {
			_ = utils.DeleteFromCloudinary(v.ImageURL)
		}
	}

	return s.Repo.Delete(id)
}

func (s *ProductService) GetProductByID(id uuid.UUID) (*models.Product, error) {
	return s.Repo.FindByID(id)
}

func (s *ProductService) Search(params dto.SearchParams) ([]models.Product, int64, error) {
	return s.Repo.SearchProducts(params)
}
