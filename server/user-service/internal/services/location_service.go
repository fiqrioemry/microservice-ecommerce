package services

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/repositories"
)

type LocationService interface {
	GetAllProvinces() ([]models.Province, error)
	GetCitiesByProvince(provinceID uint) ([]models.City, error)
	GetDistrictsByCity(cityID uint) ([]models.District, error)
	GetSubDistrictsByDistricts(disctrictID uint) ([]models.Subdistrict, error)
	GetPostalCodesBySubdistricts(provinceID uint) ([]models.PostalCode, error)
}

type locationService struct {
	repo repositories.LocationRepository
}

func NewLocationService(repo repositories.LocationRepository) LocationService {
	return &locationService{repo}
}

func (s *locationService) GetAllProvinces() ([]models.Province, error) {
	return s.repo.GetAllProvinces()
}

func (s *locationService) GetCitiesByProvince(provinceID uint) ([]models.City, error) {
	return s.repo.GetCitiesByProvince(provinceID)
}

func (s *locationService) GetDistrictsByCity(cityID uint) ([]models.District, error) {
	return s.repo.GetDistrictsByCity(cityID)
}

func (s *locationService) GetSubDistrictsByDistricts(disctrictID uint) ([]models.Subdistrict, error) {
	return s.repo.GetSubDistrictsByDistricts(disctrictID)
}

func (s *locationService) GetPostalCodesBySubdistricts(provinceID uint) ([]models.PostalCode, error) {
	return s.repo.GetPostalCodesBySubdistricts(provinceID)
}
