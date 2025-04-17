package services

import (
	"errors"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/repositories"
	"github.com/google/uuid"
)

type SubcategoryServiceInterface interface {
	GetAll() ([]models.Subcategory, error)
	Create(req dto.SubcategoryRequest) error
	Update(id uuid.UUID, req dto.SubcategoryRequest) error
	Delete(id uuid.UUID) error
	GetByID(id uuid.UUID) (*models.Subcategory, error)
}

type SubcategoryService struct {
	Repo repositories.SubcategoryRepository
}

func NewSubcategoryService(repo repositories.SubcategoryRepository) SubcategoryServiceInterface {
	return &SubcategoryService{Repo: repo}
}

func (s *SubcategoryService) GetAll() ([]models.Subcategory, error) {
	return s.Repo.FindAll()
}

func (s *SubcategoryService) Create(req dto.SubcategoryRequest) error {
	subcategory := models.Subcategory{
		Name:       req.Name,
		Slug:       req.Slug,
		Image:      req.Image,
		CategoryID: req.CategoryID,
	}
	return s.Repo.Create(&subcategory)
}

func (s *SubcategoryService) Update(id uuid.UUID, req dto.SubcategoryRequest) error {
	subcategory, err := s.Repo.FindByID(id)
	if err != nil {
		return errors.New("subcategory not found")
	}

	subcategory.Name = req.Name
	subcategory.Slug = req.Slug
	subcategory.CategoryID = req.CategoryID

	return s.Repo.Update(subcategory)
}

func (s *SubcategoryService) Delete(id uuid.UUID) error {
	subcategory, err := s.Repo.FindByID(id)
	if err != nil {
		return errors.New("subcategory not found")
	}

	if subcategory.Image != "" {
		_ = utils.DeleteFromCloudinary(subcategory.Image)
	}

	return s.Repo.Delete(id)
}

func (s *SubcategoryService) GetByID(id uuid.UUID) (*models.Subcategory, error) {
	return s.Repo.FindByID(id)
}
