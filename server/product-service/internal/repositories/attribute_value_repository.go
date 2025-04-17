package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"gorm.io/gorm"
)

type AttributeValueRepository interface {
	GetByAttribute(attrID uint) ([]models.AttributeValue, error)
	Create(val *models.AttributeValue) error
}

type attrValueRepo struct{ db *gorm.DB }

func NewAttributeValueRepository(db *gorm.DB) AttributeValueRepository {
	return &attrValueRepo{db}
}

func (r *attrValueRepo) GetByAttribute(attrID uint) ([]models.AttributeValue, error) {
	var values []models.AttributeValue
	err := r.db.Where("attribute_id = ?", attrID).Find(&values).Error
	return values, err
}

func (r *attrValueRepo) Create(val *models.AttributeValue) error {
	return r.db.Create(val).Error
}
