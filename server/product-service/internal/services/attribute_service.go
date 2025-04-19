package services

import (
	"errors"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/repositories"
)

type AttributeService interface {
	GetAll() ([]models.Attribute, error)
	GetByID(id uint) (*models.Attribute, error)
	Create(attr dto.AttributeRequest) error
	Update(id uint, attr dto.AttributeRequest) error
	Delete(id uint) error

	AddValue(req dto.AttributeValueRequest) error
	UpdateValue(id uint, newValue string) error
	DeleteValue(id uint) error
}

type attributeService struct {
	repo repositories.AttributeRepository
}

func NewAttributeService(repo repositories.AttributeRepository) AttributeService {
	return &attributeService{repo: repo}
}

func (s *attributeService) GetAll() ([]models.Attribute, error) {
	return s.repo.GetAllAttributesWithValues()
}

func (s *attributeService) GetByID(id uint) (*models.Attribute, error) {
	return s.repo.GetAttributeByID(id)
}

func (s *attributeService) Create(attr dto.AttributeRequest) error {
	newAttr := &models.Attribute{Name: attr.Name}
	return s.repo.CreateAttribute(newAttr)
}

func (s *attributeService) Update(id uint, attr dto.AttributeRequest) error {
	existing, err := s.repo.GetAttributeByID(id)
	if err != nil {
		return err
	}
	existing.Name = attr.Name
	return s.repo.UpdateAttribute(existing)
}

func (s *attributeService) Delete(id uint) error {
	return s.repo.DeleteAttribute(id)
}

func (s *attributeService) AddValue(req dto.AttributeValueRequest) error {
	if s.repo.IsValueExist(req.AttributeID, req.Value) {
		return errors.New("value already exists for this attribute")
	}
	return s.repo.CreateAttributeValue(&models.AttributeValue{
		AttributeID: req.AttributeID,
		Value:       req.Value,
	})
}

func (s *attributeService) UpdateValue(id uint, newValue string) error {
	val := &models.AttributeValue{
		ID:    id,
		Value: newValue,
	}
	return s.repo.UpdateAttributeValue(val)
}

func (s *attributeService) DeleteValue(id uint) error {
	return s.repo.DeleteAttributeValue(id)
}
