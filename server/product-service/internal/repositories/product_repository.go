package repositories

import (
	"fmt"
	"strings"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindBySlug(slug string) (*models.Product, error)
	FindByID(id uuid.UUID) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id uuid.UUID) error

	// search
	FindAllWithPreload() ([]models.Product, error)
	SearchProducts(params dto.SearchParams) ([]models.Product, int64, error)

	// Images
	SaveImages(images []models.ProductImage) error
	DeleteImagesByProductID(productID uuid.UUID) error
	FindImagesByProductID(productID uuid.UUID) ([]models.ProductImage, error)

	// Variants
	DeleteVariantProduct(variantId uuid.UUID) error
	CreateVariant(variant *models.ProductVariant) error
	UpdateVariant(variant *models.ProductVariant) error
	CreateVariantOption(option *models.ProductVariantOption) error
	FindVariantOptionValue(typeName string, value string) (*models.VariantOptionValue, error)

	// Attributes
	CreateProductAttributeValue(pav *models.ProductAttributeValue) error
	FindVariantsByProductID(productID uuid.UUID) ([]models.ProductVariant, error)
	FindVariantByID(id uuid.UUID) (*models.ProductVariant, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) UpdateVariant(v *models.ProductVariant) error {
	return r.db.Save(v).Error
}

func (r *productRepo) SearchProducts(params dto.SearchParams) ([]models.Product, int64, error) {
	query := r.db.Model(&models.Product{}).
		Preload("Category").
		Preload("Subcategory").
		Preload("ProductImage").
		Preload("ProductVariant")

	if params.Query != "" {
		query = query.Where("products.name LIKE ?", "%"+params.Query+"%")
	}
	if params.Category != "" {
		query = query.Joins("JOIN categories ON categories.id = products.category_id").
			Where("categories.slug = ?", params.Category)
	}
	if params.Subcategory != "" {
		query = query.Joins("JOIN subcategories ON subcategories.id = products.subcategory_id").
			Where("subcategories.slug = ?", params.Subcategory)
	}
	if params.InStock {
		query = query.Joins("JOIN product_variants ON product_variants.product_id = products.id").
			Where("product_variants.stock > 0")
	}
	if params.MinPrice > 0 {
		query = query.Joins("JOIN product_variants pv1 ON pv1.product_id = products.id").
			Where("pv1.price >= ?", params.MinPrice)
	}
	if params.MaxPrice > 0 {
		query = query.Joins("JOIN product_variants pv2 ON pv2.product_id = products.id").
			Where("pv2.price <= ?", params.MaxPrice)
	}
	if params.Sort != "" {
		parts := strings.Split(params.Sort, ":")
		if len(parts) == 2 {
			query = query.Order(fmt.Sprintf("%s %s", parts[0], parts[1]))
		}
	}

	var total int64
	query.Count(&total)

	var products []models.Product
	offset := (params.Page - 1) * params.Limit
	err := query.Offset(offset).Limit(params.Limit).Find(&products).Error
	return products, total, err
}

func (r *productRepo) FindAllWithPreload() ([]models.Product, error) {
	var products []models.Product
	err := r.db.
		Preload("Category").
		Preload("Subcategory").
		Preload("ProductImage").
		Preload("ProductVariant").Find(&products).Error
	return products, err
}

func (r *productRepo) FindBySlug(slug string) (*models.Product, error) {
	var product models.Product
	err := r.db.
		Preload("Category").
		Preload("Subcategory").
		Preload("ProductImage").
		Preload("ProductAttributeValue.Attribute").
		Preload("ProductAttributeValue.AttributeValue").
		Preload("ProductVariant.VariantValues.OptionValue.Type").
		Where("slug = ?", slug).
		First(&product).Error

	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) FindByID(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("ProductImage").First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepo) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepo) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Product{}, "id = ?", id).Error
}

func (r *productRepo) DeleteVariantProduct(variantId uuid.UUID) error {
	return r.db.Delete(&models.ProductVariant{}, "id = ?", variantId).Error
}

func (r *productRepo) SaveImages(images []models.ProductImage) error {
	return r.db.Create(&images).Error
}

func (r *productRepo) DeleteImagesByProductID(productID uuid.UUID) error {
	return r.db.Where("product_id = ?", productID).Delete(&models.ProductImage{}).Error
}

func (r *productRepo) FindImagesByProductID(productID uuid.UUID) ([]models.ProductImage, error) {
	var images []models.ProductImage
	err := r.db.Where("product_id = ?", productID).Find(&images).Error
	return images, err
}

func (r *productRepo) CreateVariant(variant *models.ProductVariant) error {
	return r.db.Create(variant).Error
}

func (r *productRepo) CreateVariantOption(option *models.ProductVariantOption) error {
	return r.db.Create(option).Error
}

func (r *productRepo) FindVariantOptionValue(typeName string, value string) (*models.VariantOptionValue, error) {
	var result models.VariantOptionValue
	err := r.db.Joins("JOIN variant_option_types ON variant_option_types.id = variant_option_values.type_id").
		Where("variant_option_types.name = ? AND variant_option_values.value = ?", typeName, value).
		First(&result).Error
	return &result, err
}

func (r *productRepo) CreateProductAttributeValue(pav *models.ProductAttributeValue) error {
	return r.db.Create(pav).Error
}

func (r *productRepo) FindVariantsByProductID(productID uuid.UUID) ([]models.ProductVariant, error) {
	var variants []models.ProductVariant
	err := r.db.Where("product_id = ?", productID).Find(&variants).Error
	return variants, err
}

func (r *productRepo) FindVariantByID(id uuid.UUID) (*models.ProductVariant, error) {
	var variant models.ProductVariant
	err := r.db.First(&variant, "id = ?", id).Error
	return &variant, err
}
