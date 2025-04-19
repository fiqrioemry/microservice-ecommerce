package services

import (
	"errors"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/repositories"
)

type VariantService interface {
	GetAllTypes() ([]models.VariantOptionType, error)
	CreateType(req dto.VariantTypeRequest) error
	UpdateType(id uint, req dto.VariantTypeRequest) error
	DeleteType(id uint) error

	AddValue(req dto.VariantValueRequest) error
	UpdateValue(id uint, value string) error
	DeleteValue(id uint) error

	MapToCategory(req dto.CategoryVariantTypeRequest) error
	MapToSubcategory(req dto.SubcategoryVariantTypeRequest) error
}

type variantService struct {
	repo repositories.VariantRepository
}

func NewVariantService(repo repositories.VariantRepository) VariantService {
	return &variantService{repo: repo}
}

func (s *variantService) GetAllTypes() ([]models.VariantOptionType, error) {
	return s.repo.GetAllTypesWithValues()
}

func (s *variantService) CreateType(req dto.VariantTypeRequest) error {
	return s.repo.CreateType(&models.VariantOptionType{Name: req.Name})
}

func (s *variantService) UpdateType(id uint, req dto.VariantTypeRequest) error {
	return s.repo.UpdateType(&models.VariantOptionType{ID: id, Name: req.Name})
}

func (s *variantService) DeleteType(id uint) error {
	return s.repo.DeleteType(id)
}

func (s *variantService) AddValue(req dto.VariantValueRequest) error {
	if s.repo.IsValueExist(req.TypeID, req.Value) {
		return errors.New("value already exists")
	}
	return s.repo.CreateValue(&models.VariantOptionValue{
		TypeID: req.TypeID,
		Value:  req.Value,
	})
}

func (s *variantService) UpdateValue(id uint, value string) error {
	return s.repo.UpdateValue(&models.VariantOptionValue{ID: id, Value: value})
}

func (s *variantService) DeleteValue(id uint) error {
	return s.repo.DeleteValue(id)
}

func (s *variantService) MapToCategory(req dto.CategoryVariantTypeRequest) error {
	return s.repo.MapToCategory(req.CategoryID, req.VariantTypeID)
}

func (s *variantService) MapToSubcategory(req dto.SubcategoryVariantTypeRequest) error {
	return s.repo.MapToSubcategory(req.SubcategoryID, req.VariantTypeID)
}
