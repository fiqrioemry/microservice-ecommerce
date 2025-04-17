package services

import (
	"errors"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/repositories"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthServiceInterface {
	return &AuthService{Repo: repo}
}

type AuthServiceInterface interface {
	RegisterUser(req dto.RegisterRequest) error
	GetUserByID(userID string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	LoginUser(req dto.LoginRequest) (*models.User, error)
	ResetPassword(email string, newPassword string) error
	ChangePassword(userID string, oldPassword, newPassword string) error
	GetAllUsersPaginated(search string, page int, limit int) ([]models.User, int64, error)
}

func (s *AuthService) GetUserByID(userID string) (*models.User, error) {
	return s.Repo.GetUserByID(userID)
}

func (s *AuthService) RegisterUser(req dto.RegisterRequest) error {
	if existing, _ := s.Repo.FindUserByEmail(req.Email); existing != nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return errors.New("failed to hash password")
	}

	profile := models.Profile{
		Fullname: req.Fullname,
		Avatar:   utils.RandomUserAvatar(),
	}

	user := models.User{
		ID:       uuid.New(),
		Email:    req.Email,
		Password: string(hashedPassword),
		Profile:  profile,
	}

	return s.Repo.CreateUser(&user)
}

func (s *AuthService) LoginUser(req dto.LoginRequest) (*models.User, error) {
	user, err := s.Repo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("Wrong password")
	}

	return user, nil
}

func (s *AuthService) FindUserByEmail(email string) (*models.User, error) {
	return s.Repo.FindUserByEmail(email)
}

func (s *AuthService) ResetPassword(email string, newPassword string) error {
	user, err := s.Repo.FindUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), 10)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user.Password = string(hashed)
	return s.Repo.UpdateUser(user)
}

func (s *AuthService) ChangePassword(userID string, oldPassword, newPassword string) error {
	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("old password is incorrect")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), 10)
	if err != nil {
		return errors.New("failed to hash new password")
	}

	user.Password = string(hashed)
	return s.Repo.UpdateUser(user)
}

func (s *AuthService) GetAllUsersPaginated(search string, page int, limit int) ([]models.User, int64, error) {
	return s.Repo.GetAllUsersPaginated(search, page, limit)
}
