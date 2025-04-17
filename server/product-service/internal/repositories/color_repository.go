package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"

	"gorm.io/gorm"
)

type ColorRepository interface {
	FindAll() ([]models.Color, error)
	FindByID(id uint) (*models.Color, error)
	Create(color *models.Color) error
	Update(color *models.Color) error
	Delete(id uint) error
}

type colorRepo struct {
	db *gorm.DB
}

func NewColorRepository(db *gorm.DB) ColorRepository {
	return &colorRepo{db: db}
}

func (r *colorRepo) FindAll() ([]models.Color, error) {
	var colors []models.Color
	err := r.db.Find(&colors).Error
	return colors, err
}

func (r *colorRepo) FindByID(id uint) (*models.Color, error) {
	var color models.Color
	err := r.db.First(&color, id).Error
	if err != nil {
		return nil, err
	}
	return &color, nil
}

func (r *colorRepo) Create(color *models.Color) error {
	return r.db.Create(color).Error
}

func (r *colorRepo) Update(color *models.Color) error {
	return r.db.Save(color).Error
}

func (r *colorRepo) Delete(id uint) error {
	return r.db.Delete(&models.Color{}, id).Error
}
