package services

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/repositories"
)

type AttributeService interface {
	GetAll() ([]models.Attribute, error)
	Create(req dto.CreateAttributeRequest) error
	Update(id uint, req dto.UpdateAttributeRequest) error
	Delete(id uint) error
	CreateValue(req dto.CreateAttributeValueRequest) error
	GetValues(attrID uint) ([]models.AttributeValue, error)
}

type attributeService struct {
	attrRepo repositories.AttributeRepository
	valRepo  repositories.AttributeValueRepository
}

func NewAttributeService(a repositories.AttributeRepository, v repositories.AttributeValueRepository) AttributeService {
	return &attributeService{a, v}
}

func (s *attributeService) GetAll() ([]models.Attribute, error) {
	return s.attrRepo.GetAll()
}

func (s *attributeService) Create(req dto.CreateAttributeRequest) error {
	return s.attrRepo.Create(&models.Attribute{Name: req.Name})
}

func (s *attributeService) Update(id uint, req dto.UpdateAttributeRequest) error {
	attr, err := s.attrRepo.FindByID(id)
	if err != nil {
		return err
	}
	attr.Name = req.Name
	return s.attrRepo.Update(attr)
}

func (s *attributeService) Delete(id uint) error {
	return s.attrRepo.Delete(id)
}

func (s *attributeService) CreateValue(req dto.CreateAttributeValueRequest) error {
	return s.valRepo.Create(&models.AttributeValue{
		AttributeID: req.AttributeID,
		Value:       req.Value,
	})
}

func (s *attributeService) GetValues(attrID uint) ([]models.AttributeValue, error) {
	return s.valRepo.GetByAttribute(attrID)
}
