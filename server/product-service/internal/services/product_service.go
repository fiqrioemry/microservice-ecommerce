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
	Create(req dto.CreateProductRequest) error
	GetBySlug(slug string) (*models.Product, error)
	Update(id uuid.UUID, req dto.UpdateProductRequest) error
	CreateWithImages(req dto.CreateProductRequest, imageURLs []string) error
	UpdateWithImages(id uuid.UUID, req dto.UpdateProductRequest, imageURLs []string) error
}

type ProductService struct {
	Repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductServiceInterface {
	return &ProductService{Repo: repo}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.Repo.FindAll()
}

func (s *ProductService) GetBySlug(slug string) (*models.Product, error) {
	return s.Repo.FindBySlug(slug)
}

func (s *ProductService) Create(req dto.CreateProductRequest) error {
	product := models.Product{
		Name:          req.Name,
		Slug:          req.Slug,
		Description:   req.Description,
		Price:         req.Price,
		Stock:         req.Stock,
		CategoryID:    req.CategoryID,
		SubcategoryID: req.SubcategoryID,
		IsActive:      true,
		IsFeatured:    req.IsFeatured,
	}
	return s.Repo.Create(&product)
}

func (s *ProductService) Update(id uuid.UUID, req dto.UpdateProductRequest) error {
	product, err := s.Repo.FindByID(id)
	if err != nil {
		return errors.New("product not found")
	}

	product.Name = req.Name
	product.Slug = req.Slug
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	product.CategoryID = req.CategoryID
	product.SubcategoryID = req.SubcategoryID
	product.IsFeatured = req.IsFeatured

	return s.Repo.Update(product)
}

func (s *ProductService) Delete(id uuid.UUID) error {
	images, _ := s.Repo.FindImagesByProductID(id)
	for _, img := range images {
		_ = utils.DeleteFromCloudinary(img.URL)
	}
	_ = s.Repo.DeleteImagesByProductID(id)

	return s.Repo.Delete(id)
}

func (s *ProductService) CreateWithImages(req dto.CreateProductRequest, imageURLs []string) error {
	product := models.Product{
		Name:          req.Name,
		Slug:          req.Slug,
		Description:   req.Description,
		Price:         req.Price,
		Stock:         req.Stock,
		CategoryID:    req.CategoryID,
		SubcategoryID: req.SubcategoryID,
		IsActive:      true,
		IsFeatured:    req.IsFeatured,
	}

	err := s.Repo.Create(&product)
	if err != nil {
		return err
	}

	images := []models.ProductImage{}
	for i, url := range imageURLs {
		images = append(images, models.ProductImage{
			ProductID: product.ID,
			URL:       url,
			IsPrimary: i == 0,
		})
	}

	return s.Repo.SaveImages(images)
}

func (s *ProductService) UpdateWithImages(id uuid.UUID, req dto.UpdateProductRequest, imageURLs []string) error {
	product, err := s.Repo.FindByID(id)
	if err != nil {
		return errors.New("product not found")
	}

	product.Name = req.Name
	product.Slug = req.Slug
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	product.CategoryID = req.CategoryID
	product.SubcategoryID = req.SubcategoryID
	product.IsFeatured = req.IsFeatured

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

	return nil
}
