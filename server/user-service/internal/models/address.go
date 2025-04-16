package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36);not null;index" json:"-"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	IsMain    bool      `gorm:"default:false" json:"isMain"`
	Address   string    `gorm:"type:text;not null" json:"address"`
	Province  string    `gorm:"type:varchar(255);not null" json:"province"`
	City      string    `gorm:"type:varchar(255);not null" json:"city"`
	Zipcode   string    `gorm:"type:varchar(10);not null" json:"zipcode"`
	Phone     string    `gorm:"type:varchar(20);not null" json:"phone"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (a *Address) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return
}
