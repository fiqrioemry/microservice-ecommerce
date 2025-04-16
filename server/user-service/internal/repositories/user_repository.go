package repositories

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	GetUserByID(userID string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	GetAllUsersPaginated(search string, page int, limit int) ([]models.User, int64, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func NewUserRepositoryWithDB(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := config.DB.Preload("Profile").Preload("Addresses").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	if err := config.DB.Preload("Profile").Preload("Addresses").First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) UpdateUser(user *models.User) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(user).Error
}

func (r *userRepo) GetAllUsersPaginated(search string, page, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	offset := (page - 1) * limit
	query := r.db.Model(&models.User{}).Preload("Profile").Preload("Addresses")

	if search != "" {
		query = query.Where("email LIKE ? OR id LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&users).Error
	return users, total, err
}
