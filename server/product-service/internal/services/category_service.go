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
	Create(req dto.CategoryRequest) error
	Update(id uuid.UUID, req dto.CategoryRequest) error
	GetByID(id uuid.UUID) (*models.Category, error)
	Delete(id uuid.UUID) error
}

type CategoryService struct {
	Repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryServiceInterface {
	return &CategoryService{Repo: repo}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.Repo.FindAll()
}

func (s *CategoryService) Create(req dto.CategoryRequest) error {
	category := models.Category{
		Name:  req.Name,
		Slug:  req.Slug,
		Image: req.Image,
	}
	return s.Repo.Create(&category)
}

func (s *CategoryService) Update(id uuid.UUID, req dto.CategoryRequest) error {
	category, err := s.Repo.FindByID(id)
	if err != nil {
		return errors.New("category not found")
	}

	category.Name = req.Name
	category.Slug = req.Slug
	category.Image = req.Image

	return s.Repo.Update(category)
}

func (s *CategoryService) Delete(id uuid.UUID) error {
	category, err := s.Repo.FindByID(id)
	if err != nil {
		return errors.New("category not found")
	}

	if category.Image != "" {
		_ = utils.DeleteFromCloudinary(category.Image)
	}

	return s.Repo.Delete(id)
}

func (s *CategoryService) GetByID(id uuid.UUID) (*models.Category, error) {
	return s.Repo.FindByID(id)
}
