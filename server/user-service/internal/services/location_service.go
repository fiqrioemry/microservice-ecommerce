package services

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/repositories"
)

type LocationService interface {
	GetAllProvinces() ([]models.Province, error)
	GetCitiesByProvince(provinceID uint) ([]models.City, error)
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
