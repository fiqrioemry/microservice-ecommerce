package services

import (
	"errors"

	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/repositories"
	"github.com/google/uuid"
)

type AddressServiceInterface interface {
	GetAddresses(userID string) ([]models.Address, error)
	AddAddressWithLocation(userID string, req dto.AddressRequest) error
	UpdateAddressWithLocation(userID string, addressID string, req dto.AddressRequest) error
	DeleteAddress(userID string, addressID string) error
	SetMainAddress(userID string, addressID string) error
}

type AddressService struct {
	Repo         repositories.AddressRepository
	LocationRepo repositories.LocationRepository
}

func NewAddressService(repo repositories.AddressRepository, locationRepo repositories.LocationRepository) AddressServiceInterface {
	return &AddressService{Repo: repo, LocationRepo: locationRepo}
}

func (s *AddressService) GetAddresses(userID string) ([]models.Address, error) {
	return s.Repo.GetAddressesByUserID(userID)
}

func (s *AddressService) AddAddressWithLocation(userID string, req dto.AddressRequest) error {
	province, err := s.LocationRepo.GetProvinceByID(req.ProvinceID)
	if err != nil {
		return errors.New("invalid province ID")
	}
	city, err := s.LocationRepo.GetCityByID(req.CityID)
	if err != nil {
		return errors.New("invalid city ID")
	}

	addr := models.Address{
		ID:         uuid.New(),
		UserID:     uuid.MustParse(userID),
		Name:       req.Name,
		Address:    req.Address,
		ProvinceID: req.ProvinceID,
		CityID:     req.CityID,
		Province:   province.Name,
		City:       city.Name,
		Zipcode:    req.Zipcode,
		Phone:      req.Phone,
		IsMain:     req.IsMain,
	}

	if req.IsMain {
		_ = s.Repo.UnsetAllMain(userID)
	}

	return s.Repo.CreateAddress(&addr)
}

func (s *AddressService) UpdateAddressWithLocation(userID string, addressID string, req dto.AddressRequest) error {
	addr, err := s.Repo.GetAddressByID(addressID, userID)
	if err != nil {
		return err
	}

	province, err := s.LocationRepo.GetProvinceByID(req.ProvinceID)
	if err != nil {
		return errors.New("invalid province ID")
	}
	city, err := s.LocationRepo.GetCityByID(req.CityID)
	if err != nil {
		return errors.New("invalid city ID")
	}

	addr.Name = req.Name
	addr.Address = req.Address
	addr.ProvinceID = req.ProvinceID
	addr.CityID = req.CityID
	addr.Province = province.Name
	addr.City = city.Name
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
