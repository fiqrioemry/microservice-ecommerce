package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"gorm.io/gorm"
)

type AttributeRepository interface {
	GetAllAttributesWithValues() ([]models.Attribute, error)
	GetAttributeByID(id uint) (*models.Attribute, error)
	CreateAttribute(attr *models.Attribute) error
	UpdateAttribute(attr *models.Attribute) error
	DeleteAttribute(id uint) error

	CreateAttributeValue(val *models.AttributeValue) error
	UpdateAttributeValue(val *models.AttributeValue) error
	DeleteAttributeValue(id uint) error
	IsValueExist(attributeID uint, value string) bool
}
type attributeRepo struct {
	db *gorm.DB
}

func NewAttributeRepository(db *gorm.DB) AttributeRepository {
	return &attributeRepo{db: db}
}

func (r *attributeRepo) GetAllAttributesWithValues() ([]models.Attribute, error) {
	var attrs []models.Attribute
	err := r.db.Preload("AttributeValue").Find(&attrs).Error
	return attrs, err
}

func (r *attributeRepo) GetAttributeByID(id uint) (*models.Attribute, error) {
	var attr models.Attribute
	err := r.db.First(&attr, id).Error
	return &attr, err
}

func (r *attributeRepo) CreateAttribute(attr *models.Attribute) error {
	return r.db.Create(attr).Error
}

func (r *attributeRepo) UpdateAttribute(attr *models.Attribute) error {
	return r.db.Save(attr).Error
}

func (r *attributeRepo) DeleteAttribute(id uint) error {
	return r.db.Delete(&models.Attribute{}, id).Error
}

func (r *attributeRepo) CreateAttributeValue(val *models.AttributeValue) error {
	return r.db.Create(val).Error
}

func (r *attributeRepo) UpdateAttributeValue(val *models.AttributeValue) error {
	return r.db.Save(val).Error
}

func (r *attributeRepo) DeleteAttributeValue(id uint) error {
	return r.db.Delete(&models.AttributeValue{}, id).Error
}

func (r *attributeRepo) IsValueExist(attributeID uint, value string) bool {
	var count int64
	r.db.Model(&models.AttributeValue{}).
		Where("attribute_id = ? AND value = ?", attributeID, value).
		Count(&count)
	return count > 0
}
