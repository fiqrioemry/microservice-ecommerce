package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey"`
	Name      string         `gorm:"type:varchar(100);not null;unique" json:"name"`
	Slug      string         `gorm:"type:varchar(100);uniqueIndex" json:"slug"`
	Image     string         `gorm:"type:varchar(255)" json:"image"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Subcategories []Subcategory `gorm:"foreignKey:CategoryID"`
}

type Subcategory struct {
	ID         uuid.UUID      `gorm:"type:char(36);primaryKey"`
	Name       string         `gorm:"type:varchar(100);not null;unique" json:"name"`
	Slug       string         `gorm:"type:varchar(100);uniqueIndex" json:"slug"`
	CategoryID uuid.UUID      `gorm:"type:char(36);not null" json:"-"`
	Image      string         `gorm:"type:varchar(255)"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

type Product struct {
	ID            uuid.UUID  `gorm:"type:char(36);primaryKey"`
	CategoryID    uuid.UUID  `gorm:"type:char(36);not null"`
	SubcategoryID *uuid.UUID `gorm:"type:char(36)"`
	Name          string     `gorm:"type:varchar(255);not null"`
	Slug          string     `gorm:"type:varchar(255);uniqueIndex"`
	Description   string     `gorm:"type:text"`
	IsFeatured    bool       `gorm:"default:false"`
	IsActive      bool       `gorm:"default:true"`
	Weight        float64    `gorm:"default:0" json:"weight"`
	Length        float64    `gorm:"default:0" json:"length"`
	Width         float64    `gorm:"default:0" json:"width"`
	Height        float64    `gorm:"default:0" json:"height"`
	Discount      *float64   `gorm:"type:decimal(10,2);default:0" json:"discount"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Category              Category                `gorm:"foreignKey:CategoryID"`
	Subcategory           *Subcategory            `gorm:"foreignKey:SubcategoryID"`
	ProductImage          []ProductImage          `gorm:"foreignKey:ProductID"`
	ProductVariant        []ProductVariant        `gorm:"foreignKey:ProductID"`
	ProductAttributeValue []ProductAttributeValue `gorm:"foreignKey:ProductID"`
}

type ProductImage struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	ProductID uuid.UUID `gorm:"type:char(36);not null"`
	URL       string    `gorm:"type:varchar(255);not null"`
	IsPrimary bool      `gorm:"default:false"`
}

type VariantOptionType struct {
	ID     uint                 `gorm:"primaryKey"`
	Name   string               `gorm:"type:varchar(100);unique;not null"`
	Values []VariantOptionValue `gorm:"foreignKey:TypeID"`
}

type VariantOptionValue struct {
	ID     uint   `gorm:"primaryKey"`
	TypeID uint   `gorm:"not null"`
	Value  string `gorm:"type:varchar(100);not null"`

	Type VariantOptionType `gorm:"references:ID;foreignKey:TypeID"`
}

type ProductVariant struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	ProductID uuid.UUID `gorm:"type:char(36);not null"`
	SKU       string
	Price     float64
	Stock     int
	Sold      int `gorm:"default:0"`
	ImageURL  string

	VariantValues []ProductVariantOption `gorm:"foreignKey:ProductVariantID"`
}

type ProductVariantOption struct {
	ID               uint      `gorm:"primaryKey"`
	ProductVariantID uuid.UUID `gorm:"type:char(36);not null"`
	OptionValueID    uint      `gorm:"not null"`

	OptionValue VariantOptionValue `gorm:"foreignKey:OptionValueID"`
}

type Attribute struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100);unique;not null"`

	AttributeValue []AttributeValue `gorm:"foreignKey:AttributeID"`
}

type AttributeValue struct {
	ID          uint   `gorm:"primaryKey"`
	AttributeID uint   `gorm:"not null"`
	Value       string `gorm:"type:varchar(100);not null"`
}

type ProductAttributeValue struct {
	ID               uint      `gorm:"primaryKey"`
	ProductID        uuid.UUID `gorm:"type:char(36);not null"`
	AttributeID      uint      `gorm:"not null"`
	AttributeValueID uint      `gorm:"not null"`

	Attribute      Attribute      `gorm:"foreignKey:AttributeID"`
	AttributeValue AttributeValue `gorm:"foreignKey:AttributeValueID"`
}

func (m *Category) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *Subcategory) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *Product) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *ProductImage) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (m *ProductVariant) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

type Banner struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	Position  string    `gorm:"type:varchar(50);not null"`
	ImageURL  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *Banner) BeforeCreate(tx *gorm.DB) error {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return nil
}
