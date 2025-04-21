package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/models"

	"gorm.io/gorm"
)

type AddressRepository interface {
	GetAddressesByUserID(userID string) ([]models.Address, error)
	CreateAddress(address *models.Address) error
	UpdateAddress(address *models.Address) error
	DeleteAddress(addressID string, userID string) error
	SetMainAddress(addressID string, userID string) error
	UnsetAllMain(userID string) error
	GetAddressByID(addressID string) (*models.Address, error)
}

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepo{db}
}

func (r *addressRepo) GetAddressesByUserID(userID string) ([]models.Address, error) {
	var addresses []models.Address
	err := r.db.Where("user_id = ?", userID).Order("is_main DESC").Find(&addresses).Error
	return addresses, err
}

func (r *addressRepo) CreateAddress(address *models.Address) error {
	return r.db.Create(address).Error
}

func (r *addressRepo) UpdateAddress(address *models.Address) error {
	return r.db.Save(address).Error
}

func (r *addressRepo) DeleteAddress(addressID string, userID string) error {
	return r.db.Where("id = ? AND user_id = ?", addressID, userID).Delete(&models.Address{}).Error
}

func (r *addressRepo) SetMainAddress(addressID string, userID string) error {
	return r.db.Model(&models.Address{}).Where("id = ? AND user_id = ?", addressID, userID).Update("is_main", true).Error
}

func (r *addressRepo) UnsetAllMain(userID string) error {
	return r.db.Model(&models.Address{}).Where("user_id = ?", userID).Update("is_main", false).Error
}

func (r *addressRepo) GetAddressByID(addressID string) (*models.Address, error) {
	var addr models.Address
	err := r.db.Where("id = ? ", addressID).Find(&addr).Error
	return &addr, err
}
