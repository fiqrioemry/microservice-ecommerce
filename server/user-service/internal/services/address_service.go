package services

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/repositories"

	"github.com/google/uuid"
)

type AddressServiceInterface interface {
	GetAddresses(userID string) ([]models.Address, error)
	AddAddress(userID string, req dto.AddressRequest) error
	UpdateAddress(userID string, addressID string, req dto.AddressRequest) error
	DeleteAddress(userID string, addressID string) error
	SetMainAddress(userID string, addressID string) error
}

type AddressService struct {
	Repo repositories.AddressRepository
}

func NewAddressService(repo repositories.AddressRepository) AddressServiceInterface {
	return &AddressService{Repo: repo}
}

func (s *AddressService) GetAddresses(userID string) ([]models.Address, error) {
	return s.Repo.GetAddressesByUserID(userID)
}

func (s *AddressService) AddAddress(userID string, req dto.AddressRequest) error {
	addr := models.Address{
		ID:       uuid.New(),
		UserID:   uuid.MustParse(userID),
		Name:     req.Name,
		Address:  req.Address,
		Province: req.Province,
		City:     req.City,
		Zipcode:  req.Zipcode,
		Phone:    req.Phone,
		IsMain:   req.IsMain,
	}

	if req.IsMain {
		_ = s.Repo.UnsetAllMain(userID)
	}

	return s.Repo.CreateAddress(&addr)
}

func (s *AddressService) UpdateAddress(userID string, addressID string, req dto.AddressRequest) error {
	addr, err := s.Repo.GetAddressByID(addressID, userID)
	if err != nil {
		return err
	}

	addr.Name = req.Name
	addr.Address = req.Address
	addr.Province = req.Province
	addr.City = req.City
	addr.Zipcode = req.Zipcode
	addr.Phone = req.Phone
	addr.IsMain = req.IsMain

	if req.IsMain {
		_ = s.Repo.UnsetAllMain(userID)
	}

	return s.Repo.UpdateAddress(addr)
}

func (s *AddressService) DeleteAddress(userID string, addressID string) error {
	return s.Repo.DeleteAddress(addressID, userID)
}

func (s *AddressService) SetMainAddress(userID string, addressID string) error {
	if err := s.Repo.UnsetAllMain(userID); err != nil {
		return err
	}
	return s.Repo.SetMainAddress(addressID, userID)
}
