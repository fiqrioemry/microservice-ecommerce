package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"gorm.io/gorm"
)

type ProductVariantRepository interface {
	GetByProduct(productId string) ([]models.ProductVariant, error)
	Create(v *models.ProductVariant) error
	Update(v *models.ProductVariant) error
	Delete(id string) error
	FindByID(id string) (*models.ProductVariant, error)
}

type variantRepo struct{ db *gorm.DB }

func NewVariantRepository(db *gorm.DB) ProductVariantRepository {
	return &variantRepo{db}
}

func (r *variantRepo) GetByProduct(productId string) ([]models.ProductVariant, error) {
	var result []models.ProductVariant
	err := r.db.Where("productId = ?", productId).Find(&result).Error
	return result, err
}

func (r *variantRepo) Create(v *models.ProductVariant) error {
	return r.db.Create(v).Error
}

func (r *variantRepo) Update(v *models.ProductVariant) error {
	return r.db.Save(v).Error
}

func (r *variantRepo) Delete(id string) error {
	return r.db.Delete(&models.ProductVariant{}, "id = ?", id).Error
}

func (r *variantRepo) FindByID(id string) (*models.ProductVariant, error) {
	var v models.ProductVariant
	err := r.db.First(&v, "id = ?", id).Error
	return &v, err
}
