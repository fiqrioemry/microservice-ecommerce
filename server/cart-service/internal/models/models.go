package models

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID `gorm:"type:char(36);not null;index"` // didapat dari auth/session
	CreatedAt time.Time
	UpdatedAt time.Time
	Items     []CartItem `gorm:"foreignKey:CartID"`
}

type CartItem struct {
	ID        uuid.UUID  `gorm:"type:char(36);primaryKey"`
	CartID    uuid.UUID  `gorm:"type:char(36);not null;index"`
	ProductID uuid.UUID  `gorm:"type:char(36);not null"` // dari product-service
	VariantID *uuid.UUID `gorm:"type:char(36)"`          // opsional jika ada variasi
	Quantity  int        `gorm:"not null;default:1"`
	Price     float64    `gorm:"not null"`     // harga saat dimasukkan ke cart
	IsChecked bool       `gorm:"default:true"` // untuk keperluan checkout parsial
	CreatedAt time.Time
	UpdatedAt time.Time
}
