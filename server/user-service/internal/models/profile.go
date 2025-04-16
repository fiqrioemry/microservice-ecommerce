package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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
