package services

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/repositories"
	"github.com/google/uuid"
)

type ProductAttributeValueService interface {
	GetByProduct(slug string) ([]models.ProductAttributeValue, error)
	Add(req dto.AddProductAttributeValueRequest) error
	Delete(id uint) error
}

type pavService struct {
	repo repositories.ProductAttributeValueRepository
}

func NewProductAttributeValueService(r repositories.ProductAttributeValueRepository) ProductAttributeValueService {
	return &pavService{r}
}

func (s *pavService) GetByProduct(slug string) ([]models.ProductAttributeValue, error) {
	return s.repo.GetByProduct(slug)
}

func (s *pavService) Add(req dto.AddProductAttributeValueRequest) error {
	return s.repo.Add(&models.ProductAttributeValue{
		ProductID:        uuid.MustParse(req.ProductID),
		AttributeID:      req.AttributeID,
		AttributeValueID: req.AttributeValueID,
	})
}

func (s *pavService) Delete(id uint) error {
	return s.repo.Delete(id)
}
