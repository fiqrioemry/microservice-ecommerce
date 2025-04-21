package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:text;not null" json:"-"`
	Role      string    `gorm:"type:varchar(20);not null;default:'customer'" json:"role"`
	CreatedAt time.Time `json:"joinedAt"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Profile   Profile   `gorm:"foreignKey:UserID" json:"profile"`
	Addresses []Address `gorm:"foreignKey:UserID" json:"addresses"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

type Address struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID     uuid.UUID `gorm:"type:char(36);not null;index" json:"-"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	IsMain     bool      `gorm:"default:false" json:"isMain"`
	Address    string    `gorm:"type:text;not null" json:"address"`
	ProvinceID uint      `gorm:"not null" json:"province_id"`
	CityID     uint      `gorm:"not null" json:"city_id"`
	Province   string    `gorm:"type:varchar(255);not null" json:"province"`
	City       string    `gorm:"type:varchar(255);not null" json:"city"`
	Zipcode    string    `gorm:"type:varchar(10);not null" json:"zipcode"`
	Phone      string    `gorm:"type:varchar(20);not null" json:"phone"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

type Province struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(100);not null" json:"name"`
	Cities []City `gorm:"foreignKey:ProvinceID" json:"-"`
}

type City struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	ProvinceID uint     `gorm:"not null" json:"province_id"`
	Province   Province `gorm:"foreignKey:ProvinceID" json:"-"`
	Type       string   `gorm:"type:varchar(20)" json:"type"`
	Name       string   `gorm:"type:varchar(100)" json:"name"`
	PostalCode string   `gorm:"type:varchar(20)" json:"postal_code"`
}

func (a *Address) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return
}

type Profile struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:char(36);not null;index" json:"-"`
	Fullname  string         `gorm:"type:varchar(255)" json:"fullname"`
	Birthday  string         `gorm:"type:varchar(255)" json:"birthday"`
	Gender    string         `gorm:"type:varchar(10)" json:"gender"`
	Avatar    string         `gorm:"type:varchar(255)" json:"avatar"`
	Phone     string         `gorm:"type:varchar(20)" json:"phone"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *Profile) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}

func (p *Profile) BeforeSave(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}

	if p.Gender != "" && p.Gender != "male" && p.Gender != "female" {
		return errors.New("gender must be 'male' or 'female'")
	}

	return
}
