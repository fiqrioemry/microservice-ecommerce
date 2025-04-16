package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Email     string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"type:text;not null" json:"-"`
	Role      string         `gorm:"type:varchar(20);not null;default:'customer'" json:"role"`
	CreatedAt time.Time      `json:"joinedAt"`
	UpdatedAt time.Time      `json:"-"`
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
