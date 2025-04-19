package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VariantRepository interface {
	GetAllTypesWithValues() ([]models.VariantOptionType, error)
	CreateType(*models.VariantOptionType) error
	UpdateType(*models.VariantOptionType) error
	DeleteType(id uint) error

	CreateValue(*models.VariantOptionValue) error
	UpdateValue(*models.VariantOptionValue) error
	DeleteValue(id uint) error
	IsValueExist(typeID uint, value string) bool

	MapToCategory(catID uuid.UUID, typeID uint) error
	MapToSubcategory(subID uuid.UUID, typeID uint) error
}

type variantRepo struct {
	db *gorm.DB
}

func NewVariantRepository(db *gorm.DB) VariantRepository {
	return &variantRepo{db: db}
}

func (r *variantRepo) GetAllTypesWithValues() ([]models.VariantOptionType, error) {
	var types []models.VariantOptionType
	err := r.db.Preload("VariantOptionValue").Find(&types).Error
	return types, err
}

func (r *variantRepo) CreateType(v *models.VariantOptionType) error {
	return r.db.Create(v).Error
}

func (r *variantRepo) UpdateType(v *models.VariantOptionType) error {
	return r.db.Save(v).Error
}

func (r *variantRepo) DeleteType(id uint) error {
	return r.db.Delete(&models.VariantOptionType{}, id).Error
}

func (r *variantRepo) CreateValue(val *models.VariantOptionValue) error {
	return r.db.Create(val).Error
}

func (r *variantRepo) UpdateValue(val *models.VariantOptionValue) error {
	return r.db.Save(val).Error
}

func (r *variantRepo) DeleteValue(id uint) error {
	return r.db.Delete(&models.VariantOptionValue{}, id).Error
}

func (r *variantRepo) IsValueExist(typeID uint, value string) bool {
	var count int64
	r.db.Model(&models.VariantOptionValue{}).
		Where("type_id = ? AND value = ?", typeID, value).
		Count(&count)
	return count > 0
}

func (r *variantRepo) MapToCategory(catID uuid.UUID, typeID uint) error {
	return r.db.Create(&models.CategoryVariantType{
		CategoryID:    catID,
		VariantTypeID: typeID,
	}).Error
}

func (r *variantRepo) MapToSubcategory(subID uuid.UUID, typeID uint) error {
	return r.db.Create(&models.SubcategoryVariantType{
		SubcategoryID: subID,
		VariantTypeID: typeID,
	}).Error
}
