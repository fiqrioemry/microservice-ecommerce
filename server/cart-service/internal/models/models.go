package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID        uuid.UUID  `gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID  `gorm:"type:char(36);index"`
	Items     []CartItem `gorm:"foreignKey:CartID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CartItem struct {
	ID          uuid.UUID  `gorm:"type:char(36);primaryKey"`
	CartID      uuid.UUID  `gorm:"type:char(36);index"`
	ProductID   uuid.UUID  `gorm:"type:char(36)"`
	VariantID   *uuid.UUID `gorm:"type:char(36)"`
	ProductName string     `gorm:"type:varchar(255)"`
	ImageURL    string     `gorm:"type:varchar(255)"`
	Price       float64    `gorm:"not null"`
	Quantity    int        `gorm:"default:1"`
	IsChecked   bool       `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
