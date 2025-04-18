package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindBySlug(slug string) (*models.Product, error)
	FindByID(id uuid.UUID) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id uuid.UUID) error
	SaveImages(images []models.ProductImage) error
	DeleteImagesByProductID(productID uuid.UUID) error
	FindImagesByProductID(productID uuid.UUID) ([]models.ProductImage, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Category").Preload("Subcategory").Preload("ProductImage").Find(&products).Error
	return products, err
}

func (r *productRepo) FindBySlug(slug string) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Category").Preload("Subcategory").Preload("ProductImage").Preload("ProductVariant").Preload("ProductAttributeValue").
		Where("slug = ?", slug).
		First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) FindByID(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, "id = ?", id).Error
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
