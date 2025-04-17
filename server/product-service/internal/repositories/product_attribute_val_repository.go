package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"gorm.io/gorm"
)

type ProductAttributeValueRepository interface {
	GetByProduct(productID string) ([]models.ProductAttributeValue, error)
	Add(pav *models.ProductAttributeValue) error
	Delete(id uint) error
}

type pavRepo struct{ db *gorm.DB }

func NewProductAttributeValueRepository(db *gorm.DB) ProductAttributeValueRepository {
	return &pavRepo{db}
}

func (r *pavRepo) GetByProduct(productID string) ([]models.ProductAttributeValue, error) {
	var result []models.ProductAttributeValue
	err := r.db.Where("product_id = ?", productID).Find(&result).Error
	return result, err
}

func (r *pavRepo) Add(pav *models.ProductAttributeValue) error {
	return r.db.Create(pav).Error
}

func (r *pavRepo) Delete(id uint) error {
	return r.db.Delete(&models.ProductAttributeValue{}, id).Error
}
