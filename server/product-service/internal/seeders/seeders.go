package seeders

import (
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedInitialData(db *gorm.DB) {
	dummyImage := "https://placehold.co/400x400"

	// === Categories and Subcategories ===
	fashionID := uuid.New()
	electronicsID := uuid.New()
	mensClothingID := uuid.New()
	mensShoesID := uuid.New()
	tvID := uuid.New()
	monitorID := uuid.New()

	fashion := models.Category{ID: fashionID, Name: "Fashion", Slug: "fashion", Image: dummyImage}
	electronics := models.Category{ID: electronicsID, Name: "Electronics", Slug: "electronics", Image: dummyImage}

	db.Create(&fashion)
	db.Create(&electronics)

	db.Create(&[]models.Subcategory{
		{ID: mensClothingID, Name: "Men's Clothing", Slug: "mens-clothing", CategoryID: fashionID, Image: dummyImage},
		{ID: mensShoesID, Name: "Men's Shoes", Slug: "mens-shoes", CategoryID: fashionID, Image: dummyImage},
		{ID: tvID, Name: "TV", Slug: "tv", CategoryID: electronicsID, Image: dummyImage},
		{ID: monitorID, Name: "Monitor", Slug: "monitor", CategoryID: electronicsID, Image: dummyImage},
	})

	// === Variant Types and Values ===
	sizeShoes := models.VariantOptionType{ID: 1, Name: "Size (Shoes)"}
	sizeClothing := models.VariantOptionType{ID: 2, Name: "Size (Clothing)"}
	color := models.VariantOptionType{ID: 3, Name: "Color"}
	storage := models.VariantOptionType{ID: 4, Name: "Storage"}
	db.Create(&[]models.VariantOptionType{sizeShoes, sizeClothing, color, storage})

	db.Create(&[]models.VariantOptionValue{
		{TypeID: 1, Value: "41"}, {TypeID: 1, Value: "42"}, {TypeID: 1, Value: "43"}, {TypeID: 1, Value: "44"},
		{TypeID: 2, Value: "S"}, {TypeID: 2, Value: "M"}, {TypeID: 2, Value: "L"}, {TypeID: 2, Value: "XL"}, {TypeID: 2, Value: "XXL"},
		{TypeID: 3, Value: "Red"}, {TypeID: 3, Value: "Blue"}, {TypeID: 3, Value: "Green"}, {TypeID: 3, Value: "White"},
		{TypeID: 4, Value: "64GB"}, {TypeID: 4, Value: "128GB"}, {TypeID: 4, Value: "256GB"}, {TypeID: 4, Value: "512GB"},
	})

	// === Attributes and Values ===
	brand := models.Attribute{ID: 1, Name: "Brand"}
	material := models.Attribute{ID: 2, Name: "Material"}
	typ := models.Attribute{ID: 3, Name: "Type"}
	feature := models.Attribute{ID: 4, Name: "Feature"}
	db.Create(&[]models.Attribute{brand, material, typ, feature})

	db.Create(&[]models.AttributeValue{
		{AttributeID: 1, Value: "Lenovo"}, {AttributeID: 1, Value: "HP"},
		{AttributeID: 2, Value: "Cotton"}, {AttributeID: 2, Value: "Polyester"},
		{AttributeID: 3, Value: "Casual"}, {AttributeID: 3, Value: "Formal"},
		{AttributeID: 4, Value: "Waterproof"}, {AttributeID: 4, Value: "Bluetooth"},
	})

	// === Dummy Products (simplified example) ===
	for i := 0; i < 4; i++ {
		productID := uuid.New()
		var catID uuid.UUID
		var subID uuid.UUID
		if i%2 == 0 {
			catID = fashionID
			subID = mensClothingID
		} else {
			catID = electronicsID
			subID = monitorID
		}
		product := models.Product{
			ID:            productID,
			Name:          "Product " + string(rune('A'+i)),
			Slug:          "product-" + string(rune('a'+i)),
			Description:   "Sample product",
			CategoryID:    catID,
			SubcategoryID: &subID,
			IsFeatured:    i%2 == 0,
			IsActive:      true,
			Weight:        1.0, Length: 1.0, Width: 1.0, Height: 1.0,
		}
		db.Create(&product)

		db.Create(&[]models.ProductImage{
			{ProductID: productID, URL: dummyImage, IsPrimary: true},
			{ProductID: productID, URL: dummyImage, IsPrimary: false},
		})

		variant1 := models.ProductVariant{ID: uuid.New(), ProductID: productID, SKU: "SKU-P" + string(rune('A'+i)) + "-1", Price: 100000, Stock: 10, IsActive: true, ImageURL: dummyImage}
		variant2 := models.ProductVariant{ID: uuid.New(), ProductID: productID, SKU: "SKU-P" + string(rune('A'+i)) + "-2", Price: 150000, Stock: 5, IsActive: true, ImageURL: dummyImage}
		db.Create(&[]models.ProductVariant{variant1, variant2})
	}
}
