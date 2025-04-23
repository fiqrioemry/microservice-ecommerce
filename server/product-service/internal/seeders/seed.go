package seeders

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
)

func SeedShoesAndAccessoriesProducts(db *gorm.DB) {
	log.Println("\n>>> Seeding products for category: Shoes & Accessories")

	// Ambil kategori dan subkategori
	var category models.Category
	db.First(&category, "slug = ?", "shoes-&-accessories")

	// --- Subcategories ---
	subcategories := map[string]models.Subcategory{}
	subs := []string{"sandals", "walking-style-shoes", "dress-shoes-&-oxford"}
	for _, slug := range subs {
		var sub models.Subcategory
		db.First(&sub, "slug = ?", slug)
		subcategories[slug] = sub
	}

	// --- Variant Option Value Mapping (shoes size, color) ---
	optionValueMap := map[string]models.VariantOptionValue{}
	var values []models.VariantOptionValue
	db.Find(&values)
	for _, v := range values {
		optionValueMap[v.Value] = v
	}

	// --- Data Produk ---
	products := []struct {
		Name        string
		Slug        string
		SubSlug     string
		Images      []string
		Price       float64
		Variants    []struct {
			Sizes []string
			Color string
			Stock []int
			Image string
		}
	}{
		{
			Name:    "Top Quality Wholesale Men Buckle Straps Cork Sole Sandals with Cow Leather Foot Bed",
			Slug:    "top-quality-wholesale-men-buck",
			SubSlug: "sandals",
			Images: []string{
				"https://s.alicdn.com/@sc04/kf/H0d8f82bdcad148e5a241b21d2cd5ffcfZ.jpg_720x720q50.jpg",
				"https://s.alicdn.com/@sc04/kf/H09bb8bc4065d48e493a1f4ba4674d0adZ.jpg_720x720q50.jpg",
				"https://s.alicdn.com/@sc04/kf/H0ce35e6c2ab24fb9b00938d5991bd2adp.jpg_720x720q50.jpg",
			},
			Price: 137000,
			Variants: []struct {
				Sizes []string
				Color string
				Stock []int
				Image string
			}{
				{
					Sizes: []string{"38", "39", "40"},
					Stock: []int{10, 15, 20},
					Image: "https://placehold.co/400x400",
				},
			},
		},
		// ... Tambah produk lainnya di sini (sesuai struktur yang sama)
	}

	for _, p := range products {
		sub := subcategories[p.SubSlug]
		prod := models.Product{
			ID:            uuid.New(),
			CategoryID:    category.ID,
			SubcategoryID: &sub.ID,
			Name:          p.Name,
			Slug:          p.Slug,
			IsActive:      true,
			IsFeatured:    true,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}
		if err := db.Create(&prod).Error; err != nil {
			log.Println("gagal insert produk", p.Name, err)
			continue
		}

		// Insert Images
		for i, img := range p.Images {
			imgData := models.ProductImage{
				ID:        uuid.New(),
				ProductID: prod.ID,
				URL:       img,
				IsPrimary: i == 0,
			}
			db.Create(&imgData)
		}

		// Insert Variants
		for _, v := range p.Variants {
			for i, sz := range v.Sizes {
				variant := models.ProductVariant{
					ID:        uuid.New(),
					ProductID: prod.ID,
					SKU:       uuid.NewString(),
					Price:     p.Price,
					Stock:     v.Stock[i],
					IsActive:  true,
					ImageURL:  v.Image,
				}
				if err := db.Create(&variant).Error; err != nil {
					log.Println("gagal insert variant", p.Name, err)
					continue
				}

				// Insert Variant Option
				if val, ok := optionValueMap[sz]; ok {
					db.Create(&models.ProductVariantOption{
						ProductVariantID: variant.ID,
						OptionValueID:     val.ID,
					})
				}
				if v.Color != "" {
					if val, ok := optionValueMap[v.Color]; ok {
						db.Create(&models.ProductVariantOption{
							ProductVariantID: variant.ID,
							OptionValueID:     val.ID,
						})
					}
				}
			}
		}
	}
}
