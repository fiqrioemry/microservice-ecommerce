package seeders

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
)

func SeedProductData(db *gorm.DB) {
	// Dummy images
	dummyImage := "https://placehold.co/400x400"

	// ===== CATEGORY =====
	cat1 := models.Category{
		ID:    uuid.New(),
		Name:  "Fashion",
		Slug:  "fashion",
		Image: dummyImage,
	}
	cat2 := models.Category{
		ID:    uuid.New(),
		Name:  "Electronics",
		Slug:  "electronics",
		Image: dummyImage,
	}

	// ===== SUBCATEGORY =====
	sub1 := models.Subcategory{
		ID:         uuid.New(),
		Name:       "Men Clothing",
		Slug:       "men-clothing",
		CategoryID: cat1.ID,
		Image:      dummyImage,
	}
	sub2 := models.Subcategory{
		ID:         uuid.New(),
		Name:       "Women Clothing",
		Slug:       "women-clothing",
		CategoryID: cat1.ID,
		Image:      dummyImage,
	}
	sub3 := models.Subcategory{
		ID:         uuid.New(),
		Name:       "Smartphones",
		Slug:       "smartphones",
		CategoryID: cat2.ID,
		Image:      dummyImage,
	}
	sub4 := models.Subcategory{
		ID:         uuid.New(),
		Name:       "Laptops",
		Slug:       "laptops",
		CategoryID: cat2.ID,
		Image:      dummyImage,
	}

	db.Create(&cat1)
	db.Create(&cat2)
	db.Create(&sub1)
	db.Create(&sub2)
	db.Create(&sub3)
	db.Create(&sub4)

	// ===== PRODUCTS =====
	products := []models.Product{
		{
			ID:            uuid.New(),
			Name:          "Men T-Shirt",
			Slug:          "men-tshirt",
			Description:   "Comfortable men t-shirt",
			Price:         120000,
			Stock:         100,
			CategoryID:    cat1.ID,
			SubcategoryID: &sub1.ID,
			IsFeatured:    true,
			IsActive:      true,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			ProductImage: []models.ProductImage{
				{ID: uuid.New(), URL: dummyImage, IsPrimary: true},
			},
		},
		{
			ID:            uuid.New(),
			Name:          "Women Blouse",
			Slug:          "women-blouse",
			Description:   "Trendy women blouse",
			Price:         150000,
			Stock:         60,
			CategoryID:    cat1.ID,
			SubcategoryID: &sub2.ID,
			IsFeatured:    true,
			IsActive:      true,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			ProductImage: []models.ProductImage{
				{ID: uuid.New(), URL: dummyImage, IsPrimary: true},
			},
		},
		{
			ID:            uuid.New(),
			Name:          "Smartphone X",
			Slug:          "smartphone-x",
			Description:   "Latest smartphone with great features",
			Price:         4500000,
			Stock:         40,
			CategoryID:    cat2.ID,
			SubcategoryID: &sub3.ID,
			IsFeatured:    true,
			IsActive:      true,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			ProductImage: []models.ProductImage{
				{ID: uuid.New(), URL: dummyImage, IsPrimary: true},
			},
		},
		{
			ID:            uuid.New(),
			Name:          "Laptop Pro 15",
			Slug:          "laptop-pro-15",
			Description:   "Powerful laptop for professionals",
			Price:         9500000,
			Stock:         30,
			CategoryID:    cat2.ID,
			SubcategoryID: &sub4.ID,
			IsFeatured:    false,
			IsActive:      true,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			ProductImage: []models.ProductImage{
				{ID: uuid.New(), URL: dummyImage, IsPrimary: true},
			},
		},
	}

	for _, p := range products {
		if err := db.Create(&p).Error; err != nil {
			log.Println("failed to seed product:", p.Name, err)
		}
	}

	log.Println("product, Category, and Subcategory seeding completed")
}

func SeedProductOptions(db *gorm.DB) {
	// === Colors ===
	colors := []models.Color{
		{Name: "Red", Hex: "#FF0000"},
		{Name: "Blue", Hex: "#0000FF"},
		{Name: "Green", Hex: "#00FF00"},
	}

	// === Sizes ===
	sizes := []models.Size{
		{Name: "S"},
		{Name: "M"},
		{Name: "L"},
		{Name: "XL"},
	}

	// === Attributes ===
	attributes := []models.Attribute{
		{Name: "Material"},
		{Name: "Brand"},
	}

	// === Attribute Values ===
	attributeValues := []models.AttributeValue{
		{AttributeID: 1, Value: "Cotton"},
		{AttributeID: 1, Value: "Polyester"},
		{AttributeID: 2, Value: "Brand A"},
		{AttributeID: 2, Value: "Brand B"},
	}

	for _, c := range colors {
		db.FirstOrCreate(&c, models.Color{Name: c.Name})
	}
	for _, s := range sizes {
		db.FirstOrCreate(&s, models.Size{Name: s.Name})
	}
	for _, a := range attributes {
		db.FirstOrCreate(&a, models.Attribute{Name: a.Name})
	}
	for _, av := range attributeValues {
		db.FirstOrCreate(&av, models.AttributeValue{AttributeID: av.AttributeID, Value: av.Value})
	}

	log.Println("✅ Color, Size, and Attribute seeding completed")
}

func SeedVariantsAndAttributes(db *gorm.DB) {
	var products []models.Product
	var colors []models.Color
	var sizes []models.Size
	var brandAttr models.Attribute
	var materialAttr models.Attribute
	var brandAVal, materialCottonVal models.AttributeValue

	// Ambil semua data yang dibutuhkan
	db.Find(&products)
	db.Find(&colors)
	db.Find(&sizes)
	db.Where("name = ?", "Brand").First(&brandAttr)
	db.Where("name = ?", "Material").First(&materialAttr)
	db.Where("value = ?", "Brand A").First(&brandAVal)
	db.Where("value = ?", "Cotton").First(&materialCottonVal)

	for i, product := range products {
		// Tambahkan 2 varian per produk (jika ada cukup size & color)
		if len(colors) > 0 && len(sizes) > 0 {
			variant1 := models.ProductVariant{
				ID:        uuid.New(),
				ProductID: product.ID,
				ColorID:   &colors[0].ID,
				SizeID:    &sizes[0].ID,
				SKU:       product.Slug + "-R-S",
				Price:     product.Price,
				Stock:     product.Stock / 2,
				IsActive:  true,
			}
			variant2 := models.ProductVariant{
				ID:        uuid.New(),
				ProductID: product.ID,
				ColorID:   &colors[len(colors)-1].ID,
				SizeID:    &sizes[len(sizes)-1].ID,
				SKU:       product.Slug + "-G-XL",
				Price:     product.Price + 10000,
				Stock:     product.Stock / 2,
				IsActive:  true,
			}
			db.Create(&variant1)
			db.Create(&variant2)
		}

		// Tambahkan attribute values ke produk
		attr1 := models.ProductAttributeValue{
			ProductID:        product.ID,
			AttributeID:      brandAttr.ID,
			AttributeValueID: brandAVal.ID,
		}
		attr2 := models.ProductAttributeValue{
			ProductID:        product.ID,
			AttributeID:      materialAttr.ID,
			AttributeValueID: materialCottonVal.ID,
		}
		db.Create(&attr1)
		db.Create(&attr2)

		log.Printf("✅ Variant & attributes added for product #%d: %s\n", i+1, product.Name)
	}
}
