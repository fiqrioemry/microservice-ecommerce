package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID `gorm:"type:char(36);not null;index"`
	AddressID uuid.UUID `gorm:"type:char(36);not null"`
	Status    string    `gorm:"type:varchar(50);default:'pending'"`

	TotalPrice   float64 `gorm:"type:decimal(10,2);not null"`
	ShippingCost float64 `gorm:"type:decimal(10,2);default:0"`
	AmountToPay  float64 `gorm:"type:decimal(10,2);not null"`

	Note      string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Items []OrderItem `gorm:"foreignKey:OrderID"`

	// Snapshot shipping fields
	CourierName     string `gorm:"type:varchar(255)"`
	ShippingName    string `gorm:"type:varchar(255)"`
	ShippingAddress string `gorm:"type:text"`
	City            string `gorm:"type:varchar(255)"`
	Province        string `gorm:"type:varchar(255)"`
	Zipcode         string `gorm:"type:varchar(20)"`
	Phone           string `gorm:"type:varchar(20)"`
}

type OrderItem struct {
	ID          uuid.UUID  `gorm:"type:char(36);primaryKey"`
	OrderID     uuid.UUID  `gorm:"type:char(36);not null;index"`
	ProductID   uuid.UUID  `gorm:"type:char(36);not null"`
	VariantID   *uuid.UUID `gorm:"type:char(36)"`
	ProductName string     `gorm:"type:varchar(255);not null"`
	ImageURL    string     `gorm:"type:varchar(255)"`
	Price       float64    `gorm:"type:decimal(10,2);not null"`
	Quantity    int        `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Payment struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	OrderID   uuid.UUID `gorm:"type:char(36);unique;not null"`
	Method    string    `gorm:"type:varchar(50);not null"` // e.g. midtrans, bank_transfer, cod
	Status    string    `gorm:"type:varchar(50);not null"` // e.g. pending, success, failed
	PaidAt    *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Address struct {
	Name     string
	Address  string
	City     string
	Province string
	Zipcode  string
	Phone    string
}

type Cart struct {
	Items []CartItem
}

type CartItem struct {
	ProductID   uuid.UUID
	VariantID   *uuid.UUID
	ProductName string
	ImageURL    string
	Price       float64
	Quantity    int
	IsChecked   bool
}

func (p *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}

func (a *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return
}

func (p *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
