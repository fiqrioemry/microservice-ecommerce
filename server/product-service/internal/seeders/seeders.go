// package seeders

// import (
// 	"log"
// 	"time"

// 	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
// 	"gorm.io/gorm"

// 	"github.com/google/uuid"
// )

// func SeedInitialData(db *gorm.DB) {
// 	log.Println("Seeding initial product data...")
// 	dummyImage := "https://placehold.co/400x400"

// 	// === Variant Types ===
// 	sizeType := models.VariantOptionType{Name: "Size"}
// 	colorType := models.VariantOptionType{Name: "Color"}
// 	db.Create(&sizeType)
// 	db.Create(&colorType)

// 	// === Variant Values ===
// 	sizes := []models.VariantOptionValue{
// 		{TypeID: sizeType.ID, Value: "L"},
// 		{TypeID: sizeType.ID, Value: "XL"},
// 	}
// 	colors := []models.VariantOptionValue{
// 		{TypeID: colorType.ID, Value: "Merah"},
// 		{TypeID: colorType.ID, Value: "Biru"},
// 		{TypeID: colorType.ID, Value: "Hijau"},
// 	}
// 	db.Create(&sizes)
// 	db.Create(&colors)

// 	// === Category & Subcategory ===
// 	catFashion := models.Category{
// 		ID:    uuid.New(),
// 		Name:  "Fashion",
// 		Slug:  "fashion",
// 		Image: dummyImage,
// 	}
// 	subMenClothing := models.Subcategory{
// 		ID:         uuid.New(),
// 		Name:       "Men Clothing",
// 		Slug:       "men-clothing",
// 		CategoryID: catFashion.ID,
// 		Image:      dummyImage,
// 	}
// 	db.Create(&catFashion)
// 	db.Create(&subMenClothing)

// 	// === Product ===
// 	product := models.Product{
// 		ID:            uuid.New(),
// 		Name:          "Cool T-Shirt",
// 		Slug:          "cool-t-shirt",
// 		Description:   "High quality cotton T-Shirt",
// 		CategoryID:    catFashion.ID,
// 		SubcategoryID: &subMenClothing.ID,
// 		IsFeatured:    true,
// 		IsActive:      true,
// 		Weight:        0.5,
// 		Length:        20,
// 		Width:         15,
// 		Height:        5,
// 		CreatedAt:     time.Now(),
// 		UpdatedAt:     time.Now(),
// 	}
// 	db.Create(&product)

// 	// === Product Images ===
// 	images := []models.ProductImage{
// 		{ID: uuid.New(), ProductID: product.ID, URL: dummyImage, IsPrimary: true},
// 		{ID: uuid.New(), ProductID: product.ID, URL: dummyImage, IsPrimary: false},
// 	}
// 	db.Create(&images)

// 	// === Mapping VariantOptionValue for lookup ===
// 	var valueMap = make(map[string]uint)
// 	var allValues []models.VariantOptionValue
// 	db.Find(&allValues)
// 	for _, v := range allValues {
// 		valueMap[v.Value] = v.ID
// 	}

// 	// === Product Variants ===
// 	variantConfigs := []struct {
// 		SKU   string
// 		Size  string
// 		Color string
// 		Stock int
// 		Price float64
// 	}{
// 		{"TSHIRT-L-MERAH", "L", "Merah", 2, 19.99},
// 		{"TSHIRT-L-BIRU", "L", "Biru", 4, 19.99},
// 		{"TSHIRT-L-HIJAU", "L", "Hijau", 5, 19.99},
// 		{"TSHIRT-XL-BIRU", "XL", "Biru", 7, 21.99},
// 	}

// 	for _, config := range variantConfigs {
// 		variant := models.ProductVariant{
// 			ID:        uuid.New(),
// 			ProductID: product.ID,
// 			SKU:       config.SKU,
// 			Price:     config.Price,
// 			Stock:     config.Stock,
// 			IsActive:  true,
// 			ImageURL:  dummyImage,
// 		}
// 		db.Create(&variant)

// 		// Link variant to option values
// 		optionValues := []string{config.Size, config.Color}
// 		for _, val := range optionValues {
// 			optID := valueMap[val]
// 			pvOpt := models.ProductVariantOption{
// 				ProductVariantID: variant.ID,
// 				OptionValueID:    optID,
// 			}
// 			db.Create(&pvOpt)
// 		}
// 	}

// 	// === Attributes ===
// 	brand := models.Attribute{Name: "Brand"}
// 	material := models.Attribute{Name: "Material"}
// 	db.Create(&brand)
// 	db.Create(&material)

// 	brandVals := []models.AttributeValue{
// 		{AttributeID: brand.ID, Value: "Uniqlo"},
// 		{AttributeID: brand.ID, Value: "H&M"},
// 	}
// 	materialVals := []models.AttributeValue{
// 		{AttributeID: material.ID, Value: "Cotton"},
// 		{AttributeID: material.ID, Value: "Polyester"},
// 	}
// 	db.Create(&brandVals)
// 	db.Create(&materialVals)

// 	// Mapping for lookup
// 	attrValueMap := make(map[string]uint)
// 	for _, v := range append(brandVals, materialVals...) {
// 		attrValueMap[v.Value] = v.ID
// 	}

// 	// === Link Attributes to Product ===
// 	attrLinks := []models.ProductAttributeValue{
// 		{
// 			ProductID:        product.ID,
// 			AttributeID:      brand.ID,
// 			AttributeValueID: attrValueMap["Uniqlo"],
// 		},
// 		{
// 			ProductID:        product.ID,
// 			AttributeID:      material.ID,
// 			AttributeValueID: attrValueMap["Cotton"],
// 		},
// 	}
// 	db.Create(&attrLinks)

// 	log.Println("✅ Finished seeding attributes and linking them to the product")

// }

package seeders

import (
	"log"
	"time"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

func SeedInitialData(db *gorm.DB) {
	log.Println("Seeding initial product data...")
	dummyImage := "https://placehold.co/400x400"

	// === Variant Types ===
	sizeType := models.VariantOptionType{Name: "Size"}
	colorType := models.VariantOptionType{Name: "Color"}
	db.Create(&sizeType)
	db.Create(&colorType)

	// === Variant Values ===
	sizes := []models.VariantOptionValue{
		{TypeID: sizeType.ID, Value: "L"},
		{TypeID: sizeType.ID, Value: "XL"},
	}
	colors := []models.VariantOptionValue{
		{TypeID: colorType.ID, Value: "Merah"},
		{TypeID: colorType.ID, Value: "Biru"},
		{TypeID: colorType.ID, Value: "Hijau"},
	}
	db.Create(&sizes)
	db.Create(&colors)

	// === Category & Subcategory ===
	catFashion := models.Category{
		ID:    uuid.New(),
		Name:  "Fashion",
		Slug:  "fashion",
		Image: dummyImage,
	}
	subMenClothing := models.Subcategory{
		ID:         uuid.New(),
		Name:       "Men Clothing",
		Slug:       "men-clothing",
		CategoryID: catFashion.ID,
		Image:      dummyImage,
	}
	db.Create(&catFashion)
	db.Create(&subMenClothing)

	// === Product ===
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
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	db.Create(&product)

	// === Product Images ===
	images := []models.ProductImage{
		{ID: uuid.New(), ProductID: product.ID, URL: dummyImage, IsPrimary: true},
		{ID: uuid.New(), ProductID: product.ID, URL: dummyImage, IsPrimary: false},
	}
	db.Create(&images)

	// === Mapping VariantOptionValue for lookup ===
	var valueMap = make(map[string]uint)
	var allValues []models.VariantOptionValue
	db.Find(&allValues)
	for _, v := range allValues {
		valueMap[v.Value] = v.ID
	}

	// === Product Variants ===
	variantConfigs := []struct {
		SKU   string
		Size  string
		Color string
		Stock int
		Price float64
	}{
		{"TSHIRT-L-MERAH", "L", "Merah", 2, 19.99},
		{"TSHIRT-L-BIRU", "L", "Biru", 4, 19.99},
		{"TSHIRT-L-HIJAU", "L", "Hijau", 5, 19.99},
		{"TSHIRT-XL-BIRU", "XL", "Biru", 7, 21.99},
	}

	for _, config := range variantConfigs {
		variant := models.ProductVariant{
			ID:        uuid.New(),
			ProductID: product.ID,
			SKU:       config.SKU,
			Price:     config.Price,
			Stock:     config.Stock,
			IsActive:  true,
			ImageURL:  dummyImage,
		}
		db.Create(&variant)

		// Link variant to option values
		optionValues := []string{config.Size, config.Color}
		for _, val := range optionValues {
			optID := valueMap[val]
			pvOpt := models.ProductVariantOption{
				ProductVariantID: variant.ID,
				OptionValueID:    optID,
			}
			db.Create(&pvOpt)
		}
	}

	log.Println("✅ Finished seeding product variants with Size and Color combinations")

	// === Attributes ===
	brand := models.Attribute{Name: "Brand"}
	material := models.Attribute{Name: "Material"}
	db.Create(&brand)
	db.Create(&material)

	brandVals := []models.AttributeValue{
		{AttributeID: brand.ID, Value: "Uniqlo"},
		{AttributeID: brand.ID, Value: "H&M"},
	}
	materialVals := []models.AttributeValue{
		{AttributeID: material.ID, Value: "Cotton"},
		{AttributeID: material.ID, Value: "Polyester"},
	}
	db.Create(&brandVals)
	db.Create(&materialVals)

	// Mapping AttributeValue for lookup
	attrValueMap := make(map[string]uint)
	for _, v := range append(brandVals, materialVals...) {
		attrValueMap[v.Value] = v.ID
	}

	// === Link Attributes to Product ===
	attrLinks := []models.ProductAttributeValue{
		{
			ProductID:        product.ID,
			AttributeID:      brand.ID,
			AttributeValueID: attrValueMap["Uniqlo"],
		},
		{
			ProductID:        product.ID,
			AttributeID:      material.ID,
			AttributeValueID: attrValueMap["Cotton"],
		},
	}
	db.Create(&attrLinks)

	log.Println("✅ Finished seeding attributes and linking them to the product")
}
