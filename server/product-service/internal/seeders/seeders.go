package seeders

import (
	"log"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

func SeedInitialData(db *gorm.DB) {
	log.Println("Seeding initial product data...")

	// Dummy image
	dummyImage := "https://placehold.co/400x400"

	// ====== VARIANT TYPES ======
	sizeType := models.VariantOptionType{Name: "Size"}
	colorType := models.VariantOptionType{Name: "Color"}
	db.Create(&sizeType)
	db.Create(&colorType)

	// ====== VARIANT VALUES ======
	small := models.VariantOptionValue{TypeID: sizeType.ID, Value: "S"}
	medium := models.VariantOptionValue{TypeID: sizeType.ID, Value: "M"}
	red := models.VariantOptionValue{TypeID: colorType.ID, Value: "Red"}
	blue := models.VariantOptionValue{TypeID: colorType.ID, Value: "Blue"}
	db.Create(&small)
	db.Create(&medium)
	db.Create(&red)
	db.Create(&blue)

	// ====== CATEGORY & SUBCATEGORY ======
	catFashion := models.Category{
		ID:    uuid.New(),
		Name:  "Fashion",
		Slug:  "fashion",
		Image: dummyImage,
	}
	db.Create(&catFashion)

	subMenClothing := models.Subcategory{
		ID:         uuid.New(),
		Name:       "Men Clothing",
		Slug:       "men-clothing",
		CategoryID: catFashion.ID,
		Image:      dummyImage,
	}
	db.Create(&subMenClothing)

	// ====== PRODUCT ======
	product := models.Product{
		ID:            uuid.New(),
		Name:          "Cool T-Shirt",
		Slug:          "cool-t-shirt",
		Description:   "High quality cotton T-Shirt",
		CategoryID:    catFashion.ID,
		SubcategoryID: &subMenClothing.ID,
		IsFeatured:    true,
		IsActive:      true,
		Weight:        0.5,
		Length:        20,
		Width:         15,
		Height:        5,
	}
	db.Create(&product)

	// ====== PRODUCT IMAGE ======
	images := []models.ProductImage{
		{ID: uuid.New(), ProductID: product.ID, URL: dummyImage, IsPrimary: true},
		{ID: uuid.New(), ProductID: product.ID, URL: dummyImage, IsPrimary: false},
	}
	db.Create(&images)

	// ====== PRODUCT VARIANTS ======
	variant1 := models.ProductVariant{
		ID:        uuid.New(),
		ProductID: product.ID,
		SKU:       "TSHIRT-RED-S",
		Price:     19.99,
		Stock:     10,
		IsActive:  true,
		ImageURL:  dummyImage,
	}
	variant2 := models.ProductVariant{
		ID:        uuid.New(),
		ProductID: product.ID,
		SKU:       "TSHIRT-BLUE-M",
		Price:     21.99,
		Stock:     5,
		IsActive:  true,
		ImageURL:  dummyImage,
	}
	db.Create(&variant1)
	db.Create(&variant2)

	// ====== PRODUCT VARIANT OPTIONS (link to VariantOptionValue) ======
	db.Create(&models.ProductVariantOption{
		ProductVariantID: variant1.ID,
		OptionValueID:    red.ID,
	})
	db.Create(&models.ProductVariantOption{
		ProductVariantID: variant1.ID,
		OptionValueID:    small.ID,
	})
	db.Create(&models.ProductVariantOption{
		ProductVariantID: variant2.ID,
		OptionValueID:    blue.ID,
	})
	db.Create(&models.ProductVariantOption{
		ProductVariantID: variant2.ID,
		OptionValueID:    medium.ID,
	})

	log.Println("Seeding completed.")
}
