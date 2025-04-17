package services

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/repositories"
)

type SizeService interface {
	GetAll() ([]models.Size, error)
	Create(req dto.CreateSizeRequest) error
	Update(id uint, req dto.UpdateSizeRequest) error
	Delete(id uint) error
}

type sizeService struct {
	repo repositories.SizeRepository
}

func NewSizeService(repo repositories.SizeRepository) SizeService {
	return &sizeService{repo}
}

func (s *sizeService) GetAll() ([]models.Size, error) {
	return s.repo.GetAll()
}

func (s *sizeService) Create(req dto.CreateSizeRequest) error {
	return s.repo.Create(&models.Size{Name: req.Name})
}

func (s *sizeService) Update(id uint, req dto.UpdateSizeRequest) error {
	size, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	size.Name = req.Name
	return s.repo.Update(size)
}

func (s *sizeService) Delete(id uint) error {
	return s.repo.Delete(id)
}
