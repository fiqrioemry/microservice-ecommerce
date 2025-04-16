package services

import (
	"mime/multipart"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/dto"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/models"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/repositories"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
)

type ProfileServiceInterface interface {
	GetUserByID(userID string) (*models.User, error)
	UpdateProfile(userID string, req dto.ProfileRequest, file *multipart.FileHeader) error
}

type ProfileService struct {
	Repo repositories.UserRepository
}

func NewProfileService(repo repositories.UserRepository) ProfileServiceInterface {
	return &ProfileService{Repo: repo}
}

func (s *ProfileService) GetUserByID(userID string) (*models.User, error) {
	return s.Repo.GetUserByID(userID)
}

func (s *ProfileService) UpdateProfile(userID string, req dto.ProfileRequest, file *multipart.FileHeader) error {
	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Handle avatar jika ada upload file
	if file != nil {
		// validasi file
		if err := utils.ValidateImageFile(file); err != nil {
			return err
		}

		//  Upload file baru
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		newAvatarURL, err := utils.UploadToCloudinary(src)
		if err != nil {
			return err
		}

		// hapus gambar lama kalau berbeda dan bukan dari dicebear
		if user.Profile.Avatar != "" && user.Profile.Avatar != newAvatarURL && !isDiceBear(user.Profile.Avatar) {
			_ = utils.DeleteFromCloudinary(user.Profile.Avatar)
		}

		user.Profile.Avatar = newAvatarURL
	}

	user.Profile.Fullname = req.Fullname
	user.Profile.Birthday = req.Birthday
	user.Profile.Gender = req.Gender
	user.Profile.Phone = req.Phone

	return s.Repo.UpdateUser(user)
}

func isDiceBear(url string) bool {
	return url != "" && (len(url) > 0 && (url[:30] == "https://api.dicebear.com" || url[:31] == "https://avatars.dicebear.com"))
}
