package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"gorm.io/gorm"
)

type AttributeRepository interface {
	GetAll() ([]models.Attribute, error)
	Create(attr *models.Attribute) error
	Update(attr *models.Attribute) error
	Delete(id uint) error
	FindByID(id uint) (*models.Attribute, error)
}

type attributeRepo struct{ db *gorm.DB }

func NewAttributeRepository(db *gorm.DB) AttributeRepository {
	return &attributeRepo{db}
}

func (r *attributeRepo) GetAll() ([]models.Attribute, error) {
	var data []models.Attribute
	err := r.db.Find(&data).Error
	return data, err
}

func (r *attributeRepo) Create(attr *models.Attribute) error {
	return r.db.Create(attr).Error
}

func (r *attributeRepo) Update(attr *models.Attribute) error {
	return r.db.Save(attr).Error
}

func (r *attributeRepo) Delete(id uint) error {
	return r.db.Delete(&models.Attribute{}, id).Error
}

func (r *attributeRepo) FindByID(id uint) (*models.Attribute, error) {
	var attr models.Attribute
	err := r.db.First(&attr, id).Error
	return &attr, err
}
