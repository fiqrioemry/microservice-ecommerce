package services

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/repositories"
	"github.com/google/uuid"
)

type ProductVariantService interface {
	GetByProduct(productID string) ([]models.ProductVariant, error)
	Create(req dto.CreateProductVariantRequest) error
	Update(id string, req dto.UpdateProductVariantRequest) error
	Delete(id string) error
}

type variantService struct {
	repo repositories.ProductVariantRepository
}

func NewProductVariantService(r repositories.ProductVariantRepository) ProductVariantService {
	return &variantService{r}
}

func (s *variantService) GetByProduct(productID string) ([]models.ProductVariant, error) {
	return s.repo.GetByProduct(productID)
}

func (s *variantService) Create(req dto.CreateProductVariantRequest) error {
	data := models.ProductVariant{
		ID:        uuid.New(),
		ProductID: uuid.MustParse(req.ProductID),
		ColorID:   req.ColorID,
		SizeID:    req.SizeID,
		SKU:       req.SKU,
		Price:     req.Price,
		Stock:     req.Stock,
		IsActive:  req.IsActive,
	}
	return s.repo.Create(&data)
}

func (s *variantService) Update(id string, req dto.UpdateProductVariantRequest) error {
	data, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	data.ColorID = req.ColorID
	data.SizeID = req.SizeID
	data.Price = req.Price
	data.Stock = req.Stock
	data.IsActive = req.IsActive
	return s.repo.Update(data)
}

func (s *variantService) Delete(id string) error {
	return s.repo.Delete(id)
}
