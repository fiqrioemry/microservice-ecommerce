// // package seeders
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

// 	log.Println("✅ Finished seeding product variants with Size and Color combinations")

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

// 	// Mapping AttributeValue for lookup
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
	"strings"
	"time"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"gorm.io/gorm"

	"slices"

	"github.com/google/uuid"
)

func SeedInitialData(db *gorm.DB) {
	log.Println("Seeding initial product data...")
	dummyImage := "https://placehold.co/400x400"

	// === Variant Types ===
	sizeType := models.VariantOptionType{Name: "Size"}
	colorType := models.VariantOptionType{Name: "Color"}
	capacityType := models.VariantOptionType{Name: "Capacity"}
	coverType := models.VariantOptionType{Name: "Cover Type"}
	db.Create(&sizeType)
	db.Create(&colorType)
	db.Create(&capacityType)
	db.Create(&coverType)

	sizeVals := []models.VariantOptionValue{{TypeID: sizeType.ID, Value: "S"}, {TypeID: sizeType.ID, Value: "M"}, {TypeID: sizeType.ID, Value: "L"}}
	colorVals := []models.VariantOptionValue{{TypeID: colorType.ID, Value: "Red"}, {TypeID: colorType.ID, Value: "Blue"}, {TypeID: colorType.ID, Value: "Black"}}
	capacityVals := []models.VariantOptionValue{{TypeID: capacityType.ID, Value: "64GB"}, {TypeID: capacityType.ID, Value: "128GB"}}
	coverVals := []models.VariantOptionValue{{TypeID: coverType.ID, Value: "Softcover"}, {TypeID: coverType.ID, Value: "Hardcover"}}
	db.Create(&sizeVals)
	db.Create(&colorVals)
	db.Create(&capacityVals)
	db.Create(&coverVals)

	valueMap := make(map[string]uint)
	var allValues []models.VariantOptionValue
	db.Find(&allValues)
	for _, v := range allValues {
		valueMap[v.Value] = v.ID
	}

	categoryData := []struct {
		Name        string
		Variants    []string
		SubSuffixes []string
	}{
		{"Fashion", []string{"Size", "Color"}, []string{"A1", "A2"}},
		{"Electronics", []string{"Capacity", "Color"}, []string{"A1", "A2"}},
		{"Books", []string{"Cover Type"}, []string{"A1", "A2"}},
		{"Sports", []string{"Size", "Color"}, []string{"A1", "A2"}},
	}

	for _, catData := range categoryData {
		cat := models.Category{
			ID:    uuid.New(),
			Name:  catData.Name,
			Slug:  strings.ToLower(catData.Name),
			Image: dummyImage,
		}
		db.Create(&cat)

		for _, variant := range catData.Variants {
			db.Create(&models.CategoryVariantType{
				CategoryID:    cat.ID,
				VariantTypeID: getVariantTypeIDByName(db, variant),
			})
		}

		for _, suffix := range catData.SubSuffixes {
			subName := catData.Name + " Sub " + suffix
			sub := models.Subcategory{
				ID:         uuid.New(),
				Name:       subName,
				Slug:       strings.ToLower(strings.ReplaceAll(subName, " ", "-")),
				CategoryID: cat.ID,
				Image:      dummyImage,
			}
			db.Create(&sub)

			prodName := catData.Name + " Product " + suffix
			product := models.Product{
				ID:            uuid.New(),
				Name:          prodName,
				Slug:          strings.ToLower(strings.ReplaceAll(prodName, " ", "-")),
				Description:   "Sample description for " + prodName,
				CategoryID:    cat.ID,
				SubcategoryID: &sub.ID,
				IsFeatured:    true,
				IsActive:      true,
				Weight:        1000.0,
				Length:        20,
				Width:         15,
				Height:        5,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			}
			db.Create(&product)

			images := []models.ProductImage{
				{ID: uuid.New(), ProductID: product.ID, URL: dummyImage, IsPrimary: true},
				{ID: uuid.New(), ProductID: product.ID, URL: dummyImage, IsPrimary: false},
			}
			db.Create(&images)

			// === Ambil semua kombinasi (max 3) ===
			valueSets := make([][]string, 0)
			for _, vName := range catData.Variants {
				var vals []string
				for _, v := range allValues {
					if getVariantTypeIDByName(db, vName) == v.TypeID {
						vals = append(vals, v.Value)
					}
				}
				valueSets = append(valueSets, vals)
			}
			combinations := cartesianProduct(valueSets)
			if len(combinations) > 3 {
				combinations = combinations[:3] // batasi 3 varian saja
			}

			for _, combo := range combinations {
				variant := models.ProductVariant{
					ID:        uuid.New(),
					ProductID: product.ID,
					SKU:       "SKU-" + strings.Join(combo, "-"),
					Price:     49.99,
					Stock:     10,
					IsActive:  true,
					ImageURL:  dummyImage,
				}
				db.Create(&variant)

				for _, val := range combo {
					pvOpt := models.ProductVariantOption{
						ProductVariantID: variant.ID,
						OptionValueID:    valueMap[val],
					}
					db.Create(&pvOpt)
				}
			}
		}
	}

	// Tambahan produk dummy tambahan (4 produk × 2 variant)
extraProducts := []string{"Test Product A", "Test Product B", "Test Product C", "Test Product D"}
extraCombos := [][]string{{"S", "Red"}, {"M", "Blue"}}
extraPrices := []float64{27433, 35628, 38161, 34852, 24607, 90838, 108607, 57456} // sudah diacak

priceIndex := 0
for _, name := range extraProducts {
	cat := models.Category{
		ID:    uuid.New(),
		Name:  name + " Category",
		Slug:  strings.ToLower(strings.ReplaceAll(name+" Category", " ", "-")),
		Image: dummyImage,
	}
	db.Create(&cat)

	sub := models.Subcategory{
		ID:         uuid.New(),
		Name:       name + " Sub",
		Slug:       strings.ToLower(strings.ReplaceAll(name+" Sub", " ", "-")),
		CategoryID: cat.ID,
		Image:      dummyImage,
	}
	db.Create(&sub)

	for _, variant := range []string{"Size", "Color"} {
		db.Create(&models.CategoryVariantType{
			CategoryID:    cat.ID,
			VariantTypeID: getVariantTypeIDByName(db, variant),
		})
	}

	product := models.Product{
		ID:            uuid.New(),
		Name:          name,
		Slug:          strings.ToLower(strings.ReplaceAll(name, " ", "-")),
		Description:   "Extra product " + name,
		CategoryID:    cat.ID,
		SubcategoryID: &sub.ID,
		IsFeatured:    true,
		IsActive:      true,
		Weight:        1000,
		Length:        10,
		Width:         5,
		Height:        2,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	db.Create(&product)

	db.Create(&models.ProductImage{
		ID:        uuid.New(),
		ProductID: product.ID,
		URL:       dummyImage,
		IsPrimary: true,
	})

	for _, combo := range extraCombos {
		variant := models.ProductVariant{
			ID:        uuid.New(),
			ProductID: product.ID,
			SKU:       "SKU-" + strings.Join(combo, "-"),
			Price:     extraPrices[priceIndex],
			Stock:     10,
			IsActive:  true,
			ImageURL:  dummyImage,
		}
		priceIndex++
		db.Create(&variant)

		for _, val := range combo {
			db.Create(&models.ProductVariantOption{
				ProductVariantID: variant.ID,
				OptionValueID:    valueMap[val],
			})
		}
	}
}

	log.Println("✅ Finished seeding all categories, subcategories, and products with multiple variants")
}

func getVariantTypeIDByName(db *gorm.DB, name string) uint {
	var v models.VariantOptionType
	db.First(&v, "name = ?", name)
	return v.ID
}

func cartesianProduct(sets [][]string) [][]string {
	if len(sets) == 0 {
		return [][]string{}
	}
	res := [][]string{{}}
	for _, set := range sets {
		var temp [][]string
		for _, r := range res {
			for _, v := range set {
				temp = append(temp, append(slices.Clone(r), v))
			}
		}
		res = temp
	}
	return res
}


