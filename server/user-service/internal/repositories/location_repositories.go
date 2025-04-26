package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/models"
	"gorm.io/gorm"
)

type LocationRepository interface {
	GetAllProvinces() ([]models.Province, error)
	GetCitiesByProvince(provinceID uint) ([]models.City, error)
	GetProvinceByID(id uint) (*models.Province, error)
	GetCityByID(id, provinceID uint) (*models.City, error)
}

type locationRepo struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepo{db}
}

func (r *locationRepo) GetAllProvinces() ([]models.Province, error) {
	var provinces []models.Province
	err := r.db.Order("name ASC").Find(&provinces).Error
	return provinces, err
}

func (r *locationRepo) GetAllProvinces() ([]models.Province, error) {
	var provinces []models.Province
	err := r.db.Order("name ASC").Find(&provinces).Error
	return provinces, err
}

func (r *locationRepo) GetCitiesByProvince(provinceID uint) ([]models.City, error) {
	var cities []models.City
	err := r.db.Where("province_id = ?", provinceID).Order("name ASC").Find(&cities).Error
	return cities, err
}

func (r *locationRepo) GetProvinceByID(id uint) (*models.Province, error) {
	var province models.Province
	err := r.db.First(&province, id).Error
	return &province, err
}

func (r *locationRepo) GetCityByID(id, provinceID uint) (*models.City, error) {
	var city models.City
	err := r.db.Where("province_id = ?", provinceID).First(&city, id).Error
	return &city, err
}
