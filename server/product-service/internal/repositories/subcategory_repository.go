package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubcategoryRepository interface {
	FindAll() ([]models.Subcategory, error)
	FindByID(id uuid.UUID) (*models.Subcategory, error)
	Create(subcategory *models.Subcategory) error
	Update(subcategory *models.Subcategory) error
	Delete(id uuid.UUID) error
}

type subcategoryRepo struct {
	db *gorm.DB
}

func NewSubcategoryRepository(db *gorm.DB) SubcategoryRepository {
	return &subcategoryRepo{db: db}
}

func (r *subcategoryRepo) FindAll() ([]models.Subcategory, error) {
	var subcategories []models.Subcategory
	err := r.db.Find(&subcategories).Error
	return subcategories, err
}

func (r *subcategoryRepo) FindByID(id uuid.UUID) (*models.Subcategory, error) {
	var subcategory models.Subcategory
	err := r.db.First(&subcategory, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &subcategory, nil
}

func (r *subcategoryRepo) Create(subcategory *models.Subcategory) error {
	return r.db.Create(subcategory).Error
}

func (r *subcategoryRepo) Update(subcategory *models.Subcategory) error {
	return r.db.Save(subcategory).Error
}

func (r *subcategoryRepo) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Subcategory{}, "id = ?", id).Error
}
