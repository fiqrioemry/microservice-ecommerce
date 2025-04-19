package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UnifiedCategoryRepository interface {
	// Category
	FindAll() ([]models.Category, error)
	FindByID(id uuid.UUID) (*models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(id uuid.UUID) error

	// Subcategory
	CreateSubcategory(subcategory *models.Subcategory) error
	UpdateSubcategory(subcategory *models.Subcategory) error
	DeleteSubcategory(id uuid.UUID) error
	FindSubcategoryByID(id uuid.UUID) (*models.Subcategory, error)
}

type unifiedCategoryRepo struct {
	db *gorm.DB
}

func NewUnifiedCategoryRepository(db *gorm.DB) UnifiedCategoryRepository {
	return &unifiedCategoryRepo{db: db}
}

func (r *unifiedCategoryRepo) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Preload("Subcategories").Find(&categories).Error
	return categories, err
}

func (r *unifiedCategoryRepo) FindByID(id uuid.UUID) (*models.Category, error) {
	var category models.Category
	err := r.db.Preload("Subcategories").First(&category, "id = ?", id).Error
	return &category, err
}

func (r *unifiedCategoryRepo) CreateCategory(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *unifiedCategoryRepo) UpdateCategory(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *unifiedCategoryRepo) DeleteCategory(id uuid.UUID) error {
	return r.db.Delete(&models.Category{}, "id = ?", id).Error
}

func (r *unifiedCategoryRepo) CreateSubcategory(subcategory *models.Subcategory) error {
	return r.db.Create(subcategory).Error
}

func (r *unifiedCategoryRepo) UpdateSubcategory(subcategory *models.Subcategory) error {
	return r.db.Save(subcategory).Error
}

func (r *unifiedCategoryRepo) DeleteSubcategory(id uuid.UUID) error {
	return r.db.Delete(&models.Subcategory{}, "id = ?", id).Error
}

func (r *unifiedCategoryRepo) FindSubcategoryByID(id uuid.UUID) (*models.Subcategory, error) {
	var subcategory models.Subcategory
	err := r.db.First(&subcategory, "id = ?", id).Error
	return &subcategory, err
}
