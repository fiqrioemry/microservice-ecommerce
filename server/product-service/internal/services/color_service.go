package services

import (
	"errors"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/repositories"
)

type ColorServiceInterface interface {
	GetAll() ([]models.Color, error)
	Create(req dto.ColorRequest) error
	Update(id uint, req dto.ColorRequest) error
	Delete(id uint) error
}

type ColorService struct {
	Repo repositories.ColorRepository
}

func NewColorService(repo repositories.ColorRepository) ColorServiceInterface {
	return &ColorService{Repo: repo}
}

func (s *ColorService) GetAll() ([]models.Color, error) {
	return s.Repo.FindAll()
}

func (s *ColorService) Create(req dto.ColorRequest) error {
	color := models.Color{
		Name: req.Name,
		Hex:  req.Hex,
	}
	return s.Repo.Create(&color)
}

func (s *ColorService) Update(id uint, req dto.ColorRequest) error {
	color, err := s.Repo.FindByID(id)
	if err != nil {
		return errors.New("color not found")
	}
	color.Name = req.Name
	color.Hex = req.Hex
	return s.Repo.Update(color)
}

func (s *ColorService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
