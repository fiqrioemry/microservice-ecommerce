package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID                    uuid.UUID               `gorm:"type:char(36);primaryKey" json:"id"`
	CategoryID            uuid.UUID               `gorm:"type:char(36);not null" json:"-"`
	SubcategoryID         *uuid.UUID              `gorm:"type:char(36)" json:"-"`
	Name                  string                  `gorm:"type:varchar(255);not null" json:"name"`
	Slug                  string                  `gorm:"type:varchar(255);uniqueIndex" json:"slug"`
	Description           string                  `gorm:"type:text" json:"description"`
	Price                 float64                 `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock                 int                     `gorm:"not null;default:0" json:"stock"`
	Sold                  int                     `gorm:"default:0" json:"sold"`
	IsFeatured            bool                    `gorm:"default:false" json:"isFeatured"`
	IsActive              bool                    `gorm:"default:true" json:"isActive"`
	CreatedAt             time.Time               `json:"createdAt"`
	UpdatedAt             time.Time               `json:"updatedAt"`
	DeletedAt             gorm.DeletedAt          `gorm:"index" json:"-"`
	Category              Category                `json:"category"`
	Subcategory           *Subcategory            `json:"subcategory"`
	ProductImage          []ProductImage          `gorm:"foreignKey:ProductID" json:"images"`
	ProductVariant        []ProductVariant        `gorm:"foreignKey:ProductID" json:"variants"`
	ProductAttributeValue []ProductAttributeValue `gorm:"foreignKey:ProductID" json:"attributes"`
}

type ProductImage struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	ProductID uuid.UUID `gorm:"type:char(36);not null" json:"-"`
	URL       string    `gorm:"type:varchar(255);not null" json:"url"`
	IsPrimary bool      `gorm:"default:false" json:"isPrimary"`
}

type Category struct {
	ID            uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Name          string         `gorm:"type:varchar(100);not null;unique" json:"name"`
	Slug          string         `gorm:"type:varchar(100);uniqueIndex" json:"slug"`
	Image         string         `gorm:"type:varchar(255)" json:"image"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-" `
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	Subcategories []Subcategory  `gorm:"foreignKey:CategoryID" json:"subCategories"`
}

type Subcategory struct {
	ID         uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Name       string         `gorm:"type:varchar(100);not null" json:"name"`
	Slug       string         `gorm:"type:varchar(100);uniqueIndex" json:"slug"`
	CategoryID uuid.UUID      `gorm:"type:char(36);not null" json:"-"`
	Image      string         `gorm:"type:varchar(255)" json:"image"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

type Color struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(50);not null;unique" json:"name"`
	Hex  string `gorm:"type:varchar(7)" json:"hex"`
}

type Size struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(50);not null;unique" json:"name"`
}

type ProductVariant struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	ProductID uuid.UUID `gorm:"type:char(36);not null" json:"-"`
	ColorID   *uint     `gorm:"index" json:"-"`
	SizeID    *uint     `gorm:"index" json:"-"`
	SKU       string    `gorm:"type:varchar(100);uniqueIndex" json:"sku"`
	Price     float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock     int       `gorm:"default:0" json:"stock"`
	IsActive  bool      `gorm:"default:true" json:"isActive"`

	Color Color `gorm:"foreignKey:colorId" json:"color"`
	Size  Size  `gorm:"foreignKey:sizeId" json:"size"`
}

type Attribute struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);unique;not null" json:"name"`

	AttributeValue []AttributeValue `gorm:"foreignKey:attributeId" json:"values"`
}

type AttributeValue struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	AttributeID uint   `gorm:"not null" json:"-"`
	Value       string `gorm:"type:varchar(100);not null" json:"value"`
}

type ProductAttributeValue struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	ProductID        uuid.UUID `gorm:"type:char(36);not null" json:"-"`
	AttributeID      uint      `gorm:"not null" json:"attributeId"`
	AttributeValueID uint      `gorm:"not null" json:"attributeValueId"`

	Attribute []Attribute `gorm:"foreignKey:attributeId" json:"attribute"`
}

func (p *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}

func (p *Subcategory) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}

func (a *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return
}

func (p *ProductImage) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
