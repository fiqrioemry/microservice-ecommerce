package services

import (
	"errors"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/repositories"
	"github.com/google/uuid"
)

type CategoryServiceInterface interface {
	GetAll() ([]models.Category, error)
	GetByID(id uuid.UUID) (*models.Category, error)
	Create(req dto.CategoryRequest) error
	Update(id uuid.UUID, req dto.CategoryRequest) error
	Delete(id uuid.UUID) error

	// Subcategory
	CreateSubcategory(req dto.SubcategoryRequest) error
	UpdateSubcategory(id uuid.UUID, req dto.SubcategoryRequest) error
	DeleteSubcategory(id uuid.UUID) error
	GetSubcategoryByID(id uuid.UUID) (*models.Subcategory, error)
}

type CategoryService struct {
	Repo repositories.UnifiedCategoryRepository
}

func NewCategoryService(repo repositories.UnifiedCategoryRepository) CategoryServiceInterface {
	return &CategoryService{Repo: repo}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.Repo.FindAll()
}

func (s *CategoryService) GetByID(id uuid.UUID) (*models.Category, error) {
	return s.Repo.FindByID(id)
}

func (s *CategoryService) Create(req dto.CategoryRequest) error {
	category := models.Category{
		Name:  req.Name,
		Slug:  req.Slug,
		Image: req.Image,
	}
	return s.Repo.CreateCategory(&category)
}

func (s *CategoryService) Update(id uuid.UUID, req dto.CategoryRequest) error {
	category, err := s.Repo.FindByID(id)
	if err != nil {
		return errors.New("category not found")
	}

	category.Name = req.Name
	category.Slug = req.Slug
	category.Image = req.Image

	return s.Repo.UpdateCategory(category)
}

func (s *CategoryService) Delete(id uuid.UUID) error {
	category, err := s.Repo.FindByID(id)
	if err != nil {
		return errors.New("category not found")
	}

	if category.Image != "" {
		_ = utils.DeleteFromCloudinary(category.Image)
	}

	return s.Repo.DeleteCategory(id)
}

func (s *CategoryService) CreateSubcategory(req dto.SubcategoryRequest) error {
	subcat := models.Subcategory{
		Name:       req.Name,
		Slug:       req.Slug,
		Image:      req.Image,
		CategoryID: req.CategoryID,
	}
	return s.Repo.CreateSubcategory(&subcat)
}

func (s *CategoryService) UpdateSubcategory(id uuid.UUID, req dto.SubcategoryRequest) error {
	subcat, err := s.Repo.FindSubcategoryByID(id)
	if err != nil {
		return errors.New("subcategory not found")
	}

	subcat.Name = req.Name
	subcat.Slug = req.Slug
	subcat.Image = req.Image
	subcat.CategoryID = req.CategoryID

	return s.Repo.UpdateSubcategory(subcat)
}

func (s *CategoryService) DeleteSubcategory(id uuid.UUID) error {
	subcat, err := s.Repo.FindSubcategoryByID(id)
	if err != nil {
		return errors.New("subcategory not found")
	}

	if subcat.Image != "" {
		_ = utils.DeleteFromCloudinary(subcat.Image)
	}

	return s.Repo.DeleteSubcategory(id)
}

func (s *CategoryService) GetSubcategoryByID(id uuid.UUID) (*models.Subcategory, error) {
	return s.Repo.FindSubcategoryByID(id)
}
