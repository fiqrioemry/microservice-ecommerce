package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"gorm.io/gorm"
)

type SizeRepository interface {
	GetAll() ([]models.Size, error)
	Create(size *models.Size) error
	Update(size *models.Size) error
	Delete(id uint) error
	FindByID(id uint) (*models.Size, error)
}

type sizeRepo struct {
	db *gorm.DB
}

func NewSizeRepository(db *gorm.DB) SizeRepository {
	return &sizeRepo{db}
}

func (r *sizeRepo) GetAll() ([]models.Size, error) {
	var sizes []models.Size
	err := r.db.Find(&sizes).Error
	return sizes, err
}

func (r *sizeRepo) Create(size *models.Size) error {
	return r.db.Create(size).Error
}

func (r *sizeRepo) Update(size *models.Size) error {
	return r.db.Save(size).Error
}

func (r *sizeRepo) Delete(id uint) error {
	return r.db.Delete(&models.Size{}, id).Error
}

func (r *sizeRepo) FindByID(id uint) (*models.Size, error) {
	var size models.Size
	err := r.db.First(&size, id).Error
	return &size, err
}
