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
	Name          string `json:"name" binding:"required"`
	Address       string `json:"address" binding:"required"`
	ProvinceID    uint   `json:"provinceId" binding:"required"`
	CityID        uint   `json:"cityId" binding:"required"`
	DistrictID    uint   `json:"districtId" binding:"required"`
	SubdistrictID uint   `json:"subdistrictId" binding:"required"`
	PostalCodeID  uint   `json:"postalCodeId" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
	IsMain        bool   `json:"isMain"`
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

type SearchCityRequest struct {
	Query string `form:"q" binding:"required"`
}
