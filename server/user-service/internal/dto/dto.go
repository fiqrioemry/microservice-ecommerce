package dto

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Fullname string `json:"fullname" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type ProfileRequest struct {
	Fullname string `form:"fullname" binding:"required,min=6"`
	Birthday string `form:"birthday"`
	Gender   string `form:"gender"`
	Phone    string `form:"phone"`
}

type AddressRequest struct {
	Name       string `json:"name" binding:"required"`
	Address    string `json:"address" binding:"required"`
	ProvinceID uint   `json:"province_id" binding:"required"`
	CityID     uint   `json:"city_id" binding:"required"`
	Zipcode    string `json:"zipcode" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	IsMain     bool   `json:"isMain"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6"`
}
