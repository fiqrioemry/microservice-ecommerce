package services

import (
	"mime/multipart"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/repositories"
	"github.com/google/uuid"
)

type BannerService interface {
	Create(req dto.BannerRequest, file multipart.File) error
	Get(position string) ([]dto.BannerResponse, error)
	Delete(id uuid.UUID) error
	Update(id uuid.UUID, req dto.BannerRequest, file multipart.File) error
}

type bannerService struct {
	repo repositories.BannerRepository
}

func NewBannerService(repo repositories.BannerRepository) BannerService {
	return &bannerService{repo: repo}
}

func (s *bannerService) Create(req dto.BannerRequest, file multipart.File) error {
	url, err := utils.UploadToCloudinary(file)
	if err != nil {
		return err
	}
	banner := &models.Banner{
		Position: req.Position,
		ImageURL: url,
	}
	return s.repo.Create(banner)
}

func (s *bannerService) Get(position string) ([]dto.BannerResponse, error) {
	banners, err := s.repo.GetByPosition(position)
	if err != nil {
		return nil, err
	}
	var results []dto.BannerResponse
	for _, b := range banners {
		results = append(results, dto.BannerResponse{
			ID:       b.ID.String(),
			Position: b.Position,
			ImageURL: b.ImageURL,
		})
	}
	return results, nil
}

func (s *bannerService) Delete(id uuid.UUID) error {
	b, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	_ = utils.DeleteFromCloudinary(b.ImageURL)
	return s.repo.Delete(id)
}

func (s *bannerService) Update(id uuid.UUID, req dto.BannerRequest, file multipart.File) error {
	banner, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if file != nil {
		_ = utils.DeleteFromCloudinary(banner.ImageURL)
		url, err := utils.UploadToCloudinary(file)
		if err != nil {
			return err
		}
		banner.ImageURL = url
	}

	banner.Position = req.Position
	return s.repo.Update(banner)
}

