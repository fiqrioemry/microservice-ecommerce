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
	GetMainAddress(userID string) (*models.Address, error)
}

type AddressService struct {
	Repo         repositories.AddressRepository
	LocationRepo repositories.LocationRepository
}

func NewAddressService(repo repositories.AddressRepository, locationRepo repositories.LocationRepository) AddressServiceInterface {
	return &AddressService{Repo: repo, LocationRepo: locationRepo}
}

func (s *AddressService) GetMainAddress(userID string) (*models.Address, error) {
	return s.Repo.GetMainAddress(userID)
}

func (s *AddressService) GetAddresses(userID string) ([]models.Address, error) {
	return s.Repo.GetAddressesByUserID(userID)
}

func (s *AddressService) AddAddressWithLocation(userID string, req dto.AddressRequest) error {
	// 1. Ambil Province
	provinces, err := s.LocationRepo.GetAllProvinces()
	if err != nil {
		return errors.New("failed to fetch provinces")
	}
	var province models.Province
	for _, p := range provinces {
		if p.ID == req.ProvinceID {
			province = p
			break
		}
	}
	if province.ID == 0 {
		return errors.New("invalid province ID")
	}

	// 2. Ambil City
	cities, err := s.LocationRepo.GetCitiesByProvinceID(req.ProvinceID)
	if err != nil {
		return errors.New("failed to fetch cities")
	}
	var city models.City
	for _, c := range cities {
		if c.ID == req.CityID {
			city = c
			break
		}
	}
	if city.ID == 0 {
		return errors.New("invalid city ID")
	}

	// 3. Ambil District
	districts, err := s.LocationRepo.GetDistrictsByCityID(req.CityID)
	if err != nil {
		return errors.New("failed to fetch districts")
	}
	var district models.District
	for _, d := range districts {
		if d.ID == req.DistrictID {
			district = d
			break
		}
	}
	if district.ID == 0 {
		return errors.New("invalid district ID")
	}

	// 4. Ambil Subdistrict
	subdistricts, err := s.LocationRepo.GetSubdistrictsByDistrictID(req.DistrictID)
	if err != nil {
		return errors.New("failed to fetch subdistricts")
	}
	var subdistrict models.Subdistrict
	for _, sd := range subdistricts {
		if sd.ID == req.SubdistrictID {
			subdistrict = sd
			break
		}
	}
	if subdistrict.ID == 0 {
		return errors.New("invalid subdistrict ID")
	}

	// 5. Ambil Postal Code
	postalCodes, err := s.LocationRepo.GetPostalCodesBySubdistrictID(req.SubdistrictID)
	if err != nil {
		return errors.New("failed to fetch postal codes")
	}
	var postalCode models.PostalCode
	for _, pc := range postalCodes {
		if pc.ID == req.PostalCodeID {
			postalCode = pc
			break
		}
	}
	if postalCode.ID == 0 {
		return errors.New("invalid postal code ID")
	}

	// 6. Buat Address
	addr := models.Address{
		ID:            uuid.New(),
		UserID:        uuid.MustParse(userID),
		Name:          req.Name,
		Address:       req.Address,
		ProvinceID:    req.ProvinceID,
		CityID:        req.CityID,
		DistrictID:    req.DistrictID,
		SubdistrictID: req.SubdistrictID,
		PostalCodeID:  req.PostalCodeID,
		Province:      province.Name,
		City:          city.Name,
		District:      district.Name,
		Subdistrict:   subdistrict.Name,
		PostalCode:    postalCode.PostalCode,
		Phone:         req.Phone,
		IsMain:        req.IsMain,
	}

	if req.IsMain {
		_ = s.Repo.UnsetAllMain(userID)
	}

	return s.Repo.CreateAddress(&addr)
}

func (s *AddressService) UpdateAddressWithLocation(userID string, addressID string, req dto.AddressRequest) error {
	addr, err := s.Repo.GetAddressByID(addressID)
	if err != nil {
		return err
	}

	// 1. Cek Province
	provinces, err := s.LocationRepo.GetAllProvinces()
	if err != nil {
		return errors.New("failed to fetch provinces")
	}
	var province models.Province
	for _, p := range provinces {
		if p.ID == req.ProvinceID {
			province = p
			break
		}
	}
	if province.ID == 0 {
		return errors.New("invalid province ID")
	}

	// 2. Cek City
	cities, err := s.LocationRepo.GetCitiesByProvinceID(req.ProvinceID)
	if err != nil {
		return errors.New("failed to fetch cities")
	}
	var city models.City
	for _, c := range cities {
		if c.ID == req.CityID {
			city = c
			break
		}
	}
	if city.ID == 0 {
		return errors.New("invalid city ID")
	}

	// 3. Cek District
	districts, err := s.LocationRepo.GetDistrictsByCityID(req.CityID)
	if err != nil {
		return errors.New("failed to fetch districts")
	}
	var district models.District
	for _, d := range districts {
		if d.ID == req.DistrictID {
			district = d
			break
		}
	}
	if district.ID == 0 {
		return errors.New("invalid district ID")
	}

	// 4. Cek Subdistrict
	subdistricts, err := s.LocationRepo.GetSubdistrictsByDistrictID(req.DistrictID)
	if err != nil {
		return errors.New("failed to fetch subdistricts")
	}
	var subdistrict models.Subdistrict
	for _, sd := range subdistricts {
		if sd.ID == req.SubdistrictID {
			subdistrict = sd
			break
		}
	}
	if subdistrict.ID == 0 {
		return errors.New("invalid subdistrict ID")
	}

	// 5. Cek PostalCode
	postalCodes, err := s.LocationRepo.GetPostalCodesBySubdistrictID(req.SubdistrictID)
	if err != nil {
		return errors.New("failed to fetch postal codes")
	}
	var postalCode models.PostalCode
	for _, pc := range postalCodes {
		if pc.ID == req.PostalCodeID {
			postalCode = pc
			break
		}
	}
	if postalCode.ID == 0 {
		return errors.New("invalid postal code ID")
	}

	// Update field Address
	addr.Name = req.Name
	addr.Address = req.Address
	addr.ProvinceID = req.ProvinceID
	addr.CityID = req.CityID
	addr.DistrictID = req.DistrictID
	addr.SubdistrictID = req.SubdistrictID
	addr.PostalCodeID = req.PostalCodeID
	addr.Province = province.Name
	addr.City = city.Name
	addr.District = district.Name
	addr.Subdistrict = subdistrict.Name
	addr.PostalCode = postalCode.PostalCode
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
