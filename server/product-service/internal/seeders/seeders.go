// package seeders

// import (
// 	"log"
// 	"strings"

// 	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
// 	"gorm.io/gorm"

// 	"github.com/google/uuid"
// )


// func SeedBanner(db *gorm.DB) {
// 	banners := []models.Banner{
// 		// Top Banner
// 		{ID: uuid.New(), Position: "top", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383472/topbanner03_lgpcf5.webp"},
// 		{ID: uuid.New(), Position: "top", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383471/topbanner02_supj7d.webp"},
// 		{ID: uuid.New(), Position: "top", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383470/topbanner01_wvpc7l.webp"},

// 		// Bottom Banner
// 		{ID: uuid.New(), Position: "bottom", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383469/bottombanner02_kh2krk.webp"},
// 		{ID: uuid.New(), Position: "bottom", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383469/bottombanner01_k1lylg.webp"},

// 		// Side Banner 1
// 		{ID: uuid.New(), Position: "side1", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner01_gyfi00.webp"},
// 		{ID: uuid.New(), Position: "side1", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner04_bh6d5e.webp"},

// 		// Side Banner 2
// 		{ID: uuid.New(), Position: "side2", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner02_rdtezb.webp"},
// 		{ID: uuid.New(), Position: "side2", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner03_kraq61.webp"},
// 	}

// 	for _, b := range banners {
// 		if err := db.FirstOrCreate(&b, "image_url = ?", b.ImageURL).Error; err != nil {
// 			log.Printf("failed to seed banner: %v\n", err)
// 		}
// 	}
// }

// func SeedCategories(db *gorm.DB) {
// 	placeholder := "https://placehold.co/400x400"
// 	categories := map[string][]string{
// 		"Fashion & Apparel":         {"Men's Clothing", "Hats and Caps", "Women's Clothing"},
// 		"Shoes & Accessories":       {"Sandals", "Walking Style Shoes", "Dress Shoes & Oxford"},
// 		"Gadget & Electronics":      {"Mobile Phones", "Smart TV", "Digital Camera", "Earphones"},
// 		"Health & Care":             {"Collagen", "Vitamin", "Sport Nutritions"},
// 		"Food & Beverage":           {"Energy Drink", "Noodles", "Canned Food"},
// 		"Beauty & Skin Care":        {"Lip Gloss", "Hair Extention", "Make Up"},
// 		"Sport & Entertainment":     {"Cruise Bike", "Baseball", "Roller Wheels"},
// 	}

// 	for catName, subs := range categories {
// 		cat := models.Category{
// 			ID:    uuid.New(),
// 			Name:  catName,
// 			Slug:  strings.ToLower(strings.ReplaceAll(catName, " ", "-")),
// 			Image: placeholder,
// 		}

// 		err := db.Where("name = ?", cat.Name).FirstOrCreate(&cat).Error
// 		if err != nil {
// 			log.Println("failed to create category:", catName, err)
// 			continue
// 		}

// 		for _, subName := range subs {
// 			sub := models.Subcategory{
// 				ID:         uuid.New(),
// 				Name:       subName,
// 				Slug:       strings.ToLower(strings.ReplaceAll(subName, " ", "-")),
// 				CategoryID: cat.ID,
// 				Image:      placeholder,
// 			}

// 			db.Where("name = ? AND category_id = ?", sub.Name, cat.ID).FirstOrCreate(&sub)
// 		}
// 	}
// }

// func SeedVariantTypes(db *gorm.DB) {
// 	data := []struct {
// 		Name   string
// 		Values []string
// 	}{
// 		{"shoes size", []string{"38", "39", "40", "41", "42", "43", "44"}},
// 		{"clothing size", []string{"S", "M", "L", "XL"}},
// 		{"colors", []string{"red", "green", "gray", "pink", "blue", "black", "white", "brown", "orange"}},
// 		{"ram capacity", []string{"4gb", "6gb", "8gb", "12gb"}},
// 		{"memory capacity", []string{"64gb", "128gb", "256gb", "512gb"}},
// 		{"screen size", []string{"12\"", "14\"", "18\"", "20\"", "24\"", "30\""}},
// 	}

// 	for _, item := range data {
// 		typeModel := models.VariantOptionType{Name: item.Name}
// 		if err := db.Where("name = ?", item.Name).FirstOrCreate(&typeModel).Error; err != nil {
// 			log.Printf("failed to create variant type: %s", item.Name)
// 			continue
// 		}

// 		for _, val := range item.Values {
// 			valModel := models.VariantOptionValue{
// 				TypeID: typeModel.ID,
// 				Value:  val,
// 			}
// 			db.Where("type_id = ? AND value = ?", typeModel.ID, val).FirstOrCreate(&valModel)
// 		}
// 	}
// }

// func SeedShoesAndAccessoriesProducts(db *gorm.DB) {
// 	placeholder := "https://placehold.co/400x400"

// 	// Ambil category dan subcategory ID
// 	var shoesCategory models.Category
// 	db.Where("name = ?", "Shoes & Accessories").First(&shoesCategory)

// 	subcategories := map[string]uuid.UUID{}
// 	subNames := []string{"Sandals", "Walking Style Shoes", "Dress Shoes & Oxford"}
// 	for _, name := range subNames {
// 		var sub models.Subcategory
// 		db.Where("name = ? AND category_id = ?", name, shoesCategory.ID).First(&sub)
// 		subcategories[name] = sub.ID
// 	}

// 	// Ambil shoes size type
// 	var shoesSizeType models.VariantOptionType
// 	db.Where("name = ?", "shoes size").First(&shoesSizeType)
// 	shoesSizes := map[string]uint{}
// 	var shoesValues []models.VariantOptionValue
// 	db.Where("type_id = ?", shoesSizeType.ID).Find(&shoesValues)
// 	for _, v := range shoesValues {
// 		shoesSizes[v.Value] = v.ID
// 	}

// 	// Ambil colors
// 	var colorType models.VariantOptionType
// 	db.Where("name = ?", "colors").First(&colorType)
// 	colors := map[string]uint{}
// 	var colorValues []models.VariantOptionValue
// 	db.Where("type_id = ?", colorType.ID).Find(&colorValues)
// 	for _, v := range colorValues {
// 		colors[v.Value] = v.ID
// 	}

// 	// Helper: tambah produk + gambar + variant
// 	createProduct := func(name string, price float64, images []string, subID uuid.UUID, variants []map[string]interface{}) {
// 		product := models.Product{
// 			ID:            uuid.New(),
// 			Name:          name,
// 			Slug:          strings.ToLower(strings.ReplaceAll(name[:30], " ", "-")),
// 			CategoryID:    shoesCategory.ID,
// 			SubcategoryID: &subID,
// 			Description:   name,
// 			IsFeatured:    true,
// 			IsActive:      true,
// 		}
// 		db.Create(&product)
// 		for i, img := range images {
// 			db.Create(&models.ProductImage{
// 				ID:        uuid.New(),
// 				ProductID: product.ID,
// 				URL:       img,
// 				IsPrimary: i == 0,
// 			})
// 		}

// 		// Create variants
// 		for _, v := range variants {
// 			varID := uuid.New()
// 			db.Create(&models.ProductVariant{
// 				ID:        varID,
// 				ProductID: product.ID,
// 				SKU:       uuid.NewString(),
// 				Price:     price,
// 				Stock:     v["stock"].(int),
// 				IsActive:  true,
// 				ImageURL:  v["image"].(string),
// 			})

// 			// Create options (shoes size + optional color)
// 			if size, ok := v["size"].(string); ok {
// 				db.Create(&models.ProductVariantOption{
// 					ProductVariantID: varID,
// 					OptionValueID:    shoesSizes[size],
// 				})
// 			}
// 			if color, ok := v["color"].(string); ok {
// 				db.Create(&models.ProductVariantOption{
// 					ProductVariantID: varID,
// 					OptionValueID:    colors[color],
// 				})
// 			}
// 		}
// 	}

// 	createProduct(
// 		"Top Quality Whole Men Buckle Straps Cork Sole Sandals with Cow Leather Foot Bed",
// 		137000,
// 		[]string{
// 			"https://s.alicdn.com/@sc04/kf/H0d8f82bdcad148e5a241b21d2cd5ffcfZ.jpg_720x720q50.jpg",
// 			"https://s.alicdn.com/@sc04/kf/H09bb8bc4065d48e493a1f4ba4674d0adZ.jpg_720x720q50.jpg",
// 			"https://s.alicdn.com/@sc04/kf/H0ce35e6c2ab24fb9b00938d5991bd2adp.jpg_720x720q50.jpg",
// 		},
// 		subcategories["Sandals"],
// 		[]map[string]interface{}{
// 			{"size": "38", "stock": 10, "image": placeholder},
// 			{"size": "39", "stock": 15, "image": placeholder},
// 			{"size": "40", "stock": 20, "image": placeholder},
// 		},
// 	)
// 	createProduct("Top Quality Wholesale Men Buckle Straps Cork Sole Sandals with Cow Leather Foot Bed", 137000, []string{"https://s.alicdn.com/@sc04/kf/H0d8f82bdcad148e5a241b21d2cd5ffcfZ.jpg_720x720q50.jpg", "https://s.alicdn.com/@sc04/kf/H09bb8bc4065d48e493a1f4ba4674d0adZ.jpg_720x720q50.jpg", "https://s.alicdn.com/@sc04/kf/H0ce35e6c2ab24fb9b00938d5991bd2adp.jpg_720x720q50.jpg"}, subcategories["Sandals"], []map[string]interface{}{{"size": "38", "stock": 10, "image": placeholder}, {"size": "39", "stock": 15, "image": placeholder}, {"size": "40", "stock": 20, "image": placeholder}})
// 	createProduct("New Summer for Men Outdoor Casual Beach Fashion Sandals Ventilate Men Sandals New Style Driving Shoes", 165000, []string{"https://s.alicdn.com/@sc04/kf/H9e65cb5dc6164a94a0cbe2c8909349aaU.jpg_720x720q50.jpg", "https://s.alicdn.com/@sc04/kf/H1dffaa00338f4aceb9ef779f32f7f8aek.jpg_720x720q50.jpg", "https://s.alicdn.com/@sc04/kf/Hc778e6c76f7341bf868b11b09d6e72cbx.jpg_720x720q50.jpg"}, subcategories["Sandals"], []map[string]interface{}{{"size": "38", "stock": 10, "image": placeholder}, {"size": "39", "stock": 15, "image": placeholder}, {"size": "40", "stock": 20, "image": placeholder}})
// 	createProduct("Factory Price Customize Pu Leather Sport Open Toe Men Anti slip Outdoor Beach Sandals", 212000, []string{"https://s.alicdn.com/@sc04/kf/H9c03348c33b44de2abdecdd71f3579afy.jpg_720x720q50.jpg", "https://s.alicdn.com/@sc04/kf/Hc807554f6d4a4817b95e97c0e4a1d8260.jpg_720x720q50.jpg", "https://s.alicdn.com/@sc04/kf/H3636075fde22433cbad21143f45234c8r.jpg_720x720q50.jpg"}, subcategories["Sandals"], []map[string]interface{}{{"size": "40", "stock": 10, "image": placeholder}, {"size": "41", "stock": 15, "image": placeholder}, {"size": "42", "stock": 20, "image": placeholder}})
// 	createProduct("Shoes for Men New Walking Style Styles Casual Shoe Sport Running Sneakers Chaussures Pour Hommes", 325000, []string{"https://s.alicdn.com/@sc04/kf/H83a7cb7c4fa24850a2087db98e7651870.jpg_720x720q50.jpg", "https://s.alicdn.com/@sc04/kf/H1d90140e382e413da9aad9dd7b5b054ct.jpg_720x720q50.jpg"}, subcategories["Walking Style Shoes"], []map[string]interface{}{{"size": "40", "color": "red", "stock": 10, "image": "https://s.alicdn.com/@sc04/kf/Hd4b21d37705a46f392bb17a8c7176af0J.jpg_720x720q50.jpg"}, {"size": "41", "color": "red", "stock": 15, "image": "https://s.alicdn.com/@sc04/kf/Hd4b21d37705a46f392bb17a8c7176af0J.jpg_720x720q50.jpg"}, {"size": "42", "color": "red", "stock": 20, "image": "https://s.alicdn.com/@sc04/kf/Hd4b21d37705a46f392bb17a8c7176af0J.jpg_720x720q50.jpg"}, {"size": "40", "color": "blue", "stock": 10, "image": "https://sc04.alicdn.com/kf/H6eee21aaf1f84c17bd4ba77141ad301df.jpg"}, {"size": "41", "color": "blue", "stock": 15, "image": "https://sc04.alicdn.com/kf/H6eee21aaf1f84c17bd4ba77141ad301df.jpg"}, {"size": "42", "color": "blue", "stock": 20, "image": "https://sc04.alicdn.com/kf/H6eee21aaf1f84c17bd4ba77141ad301df.jpg"}})
// 	createProduct("Top Quality Fashion Zapatillas De Hombre Women Chunky Sneakers Black New Style Walking Style Casual Sneaker Shoes Men", 275000, []string{"https://s.alicdn.com/@sc04/kf/H609ac6c3895645f397958193851a9399x.png_720x720q50.jpg", "https://s.alicdn.com/@sc04/kf/H9643425848ab4efb8855e5af580ca20aA.png_720x720q50.jpg"}, subcategories["Walking Style Shoes"], nil)
// 	createProduct("Sports and Leisure Shoes, Running, Walking, Breathable, Casual, Office Shoes, Wholesale, Customized, High-quality Sports Shoes", 425000, []string{"https://s.alicdn.com/@sc04/kf/H83a7cb7c4fa24850a2087db98e7651870.jpg_720x720q50.jpg", "https://s.alicdn.com/@sc04/kf/H1d90140e382e413da9aad9dd7b5b054ct.jpg_720x720q50.jpg"}, subcategories["Walking Style Shoes"], []map[string]interface{}{{"size": "40", "color": "white", "stock": 10, "image": "https://s.alicdn.com/@sc04/kf/H45bf3bcfe88f4bf7837d88a04b0888a7l.jpg_720x720q50.jpg"}, {"size": "41", "color": "white", "stock": 15, "image": "https://s.alicdn.com/@sc04/kf/H45bf3bcfe88f4bf7837d88a04b0888a7l.jpg_720x720q50.jpg"}, {"size": "42", "color": "white", "stock": 20, "image": "https://s.alicdn.com/@sc04/kf/H45bf3bcfe88f4bf7837d88a04b0888a7l.jpg_720x720q50.jpg"}, {"size": "40", "color": "black", "stock": 10, "image": "https://s.alicdn.com/@sc04/kf/H6cb1b2718899416e8d96f631a1b4a2e1W.jpg_720x720q50.jpg"}, {"size": "41", "color": "black", "stock": 15, "image": "https://s.alicdn.com/@sc04/kf/H6cb1b2718899416e8d96f631a1b4a2e1W.jpg_720x720q50.jpg"}, {"size": "42", "color": "black", "stock": 20, "image": "https://s.alicdn.com/@sc04/kf/H6cb1b2718899416e8d96f631a1b4a2e1W.jpg_720x720q50.jpg"}})
// }


// func SeedFashionApparelProducts(db *gorm.DB) {
// 	placeholder := "https://placehold.co/400x400"

// 	var fashionCategory models.Category
// 	db.Where("name = ?", "Fashion & Apparel").First(&fashionCategory)

// 	subcategories := map[string]uuid.UUID{}
// 	for _, name := range []string{"Men's Clothing"} {
// 		var sub models.Subcategory
// 		db.Where("name = ? AND category_id = ?", name, fashionCategory.ID).First(&sub)
// 		subcategories[name] = sub.ID
// 	}

// 	sizes := map[string]uint{}
// 	var sizeType models.VariantOptionType
// 	db.Where("name = ?", "clothing size").First(&sizeType)
// 	var sizeValues []models.VariantOptionValue
// 	db.Where("type_id = ?", sizeType.ID).Find(&sizeValues)
// 	for _, v := range sizeValues {
// 		sizes[v.Value] = v.ID
// 	}

// 	colors := map[string]uint{}
// 	var colorType models.VariantOptionType
// 	db.Where("name = ?", "colors").First(&colorType)
// 	var colorValues []models.VariantOptionValue
// 	db.Where("type_id = ?", colorType.ID).Find(&colorValues)
// 	for _, v := range colorValues {
// 		colors[v.Value] = v.ID
// 	}

// 	createProduct := func(name string, price float64, images []string, subID uuid.UUID, variants []map[string]interface{}) {
// 		product := models.Product{
// 			ID:            uuid.New(),
// 			Name:          name,
// 			Slug:          strings.ToLower(strings.ReplaceAll(name[:30], " ", "-")),
// 			CategoryID:    fashionCategory.ID,
// 			SubcategoryID: &subID,
// 			Description:   name,
// 			IsFeatured:    false,
// 			IsActive:      true,
// 		}
// 		db.Create(&product)
// 		for i, img := range images {
// 			db.Create(&models.ProductImage{
// 				ID:        uuid.New(),
// 				ProductID: product.ID,
// 				URL:       img,
// 				IsPrimary: i == 0,
// 			})
// 		}

// 		for _, v := range variants {
// 			varID := uuid.New()
// 			db.Create(&models.ProductVariant{
// 				ID:        varID,
// 				ProductID: product.ID,
// 				SKU:       uuid.NewString(),
// 				Price:     price,
// 				Stock:     v["stock"].(int),
// 				IsActive:  true,
// 				ImageURL:  v["image"].(string),
// 			})
// 			if size, ok := v["size"].(string); ok {
// 				db.Create(&models.ProductVariantOption{
// 					ProductVariantID: varID,
// 					OptionValueID:    sizes[size],
// 				})
// 			}
// 			if color, ok := v["color"].(string); ok {
// 				db.Create(&models.ProductVariantOption{
// 					ProductVariantID: varID,
// 					OptionValueID:    colors[color],
// 				})
// 			}
// 		}
// 	}

// 	// Product 1
// 	createProduct("Top Quality Casual Mens Shirt", 120000, []string{
// 		"https://s.alicdn.com/@sc04/kf/H7a6558c2616241aaafb4e6e119042238C.png_720x720q50.jpg",
// 		"https://s.alicdn.com/@sc04/kf/H8ed62347008341f0a6927543e7cc5504c.jpg_720x720q50.jpg",
// 	}, subcategories["Men's Clothing"], []map[string]interface{}{
// 		{"size": "S", "stock": 10, "image": placeholder},
// 		{"size": "M", "stock": 20, "image": placeholder},
// 		{"size": "L", "stock": 40, "image": placeholder},
// 	})

// 	// Product 2
// 	createProduct("Premium Men Jacket Fashion Style x", 235000, []string{
// 		"https://s.alicdn.com/@sc04/kf/H45c2134872bc4b6e89ca5b55db292019q.jpg_720x720q50.jpg",
// 		"https://s.alicdn.com/@sc04/kf/Hf54245977fbb469cbc3282fd31d2d65dl.jpg?avif=close",
// 	}, subcategories["Men's Clothing"], []map[string]interface{}{
// 		{"color": "red", "stock": 13, "image": placeholder},
// 		{"color": "black", "stock": 20, "image": placeholder},
// 		{"color": "blue", "stock": 5, "image": placeholder},
// 		{"color": "white", "stock": 10, "image": placeholder},
// 	})

// 	// Product 3
// 	createProduct("Color Variant Outfit T-Shirt", 177000, []string{
// 		"https://s.alicdn.com/@sc04/kf/H37fdf58f0dd54cdca5a086055fe1e101q.jpg_720x720q50.jpg",
// 	}, subcategories["Men's Clothing"], []map[string]interface{}{
// 		{"color": "black", "size": "M", "stock": 5, "image": "https://s.alicdn.com/@sc04/kf/H37fdf58f0dd54cdca5a086055fe1e101q.jpg_720x720q50.jpg"},
// 		{"color": "black", "size": "L", "stock": 10, "image": "https://s.alicdn.com/@sc04/kf/H37fdf58f0dd54cdca5a086055fe1e101q.jpg_720x720q50.jpg"},
// 		{"color": "blue", "size": "L", "stock": 12, "image": "https://s.alicdn.com/@sc04/kf/Hcf176a6f30b840a89e5f0a3c60bb1ecec.jpg_720x720q50.jpg"},
// 		{"color": "blue", "size": "XL", "stock": 20, "image": "https://s.alicdn.com/@sc04/kf/Hcf176a6f30b840a89e5f0a3c60bb1ecec.jpg_720x720q50.jpg"},
// 		{"color": "pink", "size": "S", "stock": 50, "image": "https://s.alicdn.com/@sc04/kf/H32a280fda84440be8c22f209b589a36bJ.jpg_720x720q50.jpg"},
// 		{"color": "pink", "size": "M", "stock": 25, "image": "https://s.alicdn.com/@sc04/kf/H32a280fda84440be8c22f209b589a36bJ.jpg_720x720q50.jpg"},
// 	})

// 	// Product 4
// 	createProduct("Dualx Tone Hoodie Classic Style", 357000, []string{
// 		"https://s.alicdn.com/@sc04/kf/H1db0a8f8f25041dab98442eaa144ff3ay.jpg_720x720q50.jpg",
// 	}, subcategories["Men's Clothing"], []map[string]interface{}{
// 		{"color": "black", "size": "M", "stock": 5, "image": "https://s.alicdn.com/@sc04/kf/H5722ff2a7e144e5aa83103a52770bed0G.jpg_720x720q50.jpg"},
// 		{"color": "black", "size": "L", "stock": 10, "image": "https://s.alicdn.com/@sc04/kf/H5722ff2a7e144e5aa83103a52770bed0G.jpg_720x720q50.jpg"},
// 		{"color": "red", "size": "M", "stock": 12, "image": "https://s.alicdn.com/@sc04/kf/H07b065e7028640488feeb1278c3cff09M.jpg_720x720q50.jpg"},
// 		{"color": "red", "size": "L", "stock": 12, "image": "https://s.alicdn.com/@sc04/kf/H07b065e7028640488feeb1278c3cff09M.jpg_720x720q50.jpg"},
// 		{"color": "red", "size": "XL", "stock": 20, "image": "https://s.alicdn.com/@sc04/kf/H07b065e7028640488feeb1278c3cff09M.jpg_720x720q50.jpg"},
// 	})
// }


// func SeedHatsAndCaps(db *gorm.DB) {

// 	var category models.Category
// 	db.Where("name = ?", "Fashion & Apparel").First(&category)

// 	var sub models.Subcategory
// 	db.Where("name = ? AND category_id = ?", "Hats and Caps", category.ID).First(&sub)

// 	colors := map[string]uint{}
// 	var colorType models.VariantOptionType
// 	db.Where("name = ?", "colors").First(&colorType)
// 	var colorValues []models.VariantOptionValue
// 	db.Where("type_id = ?", colorType.ID).Find(&colorValues)
// 	for _, v := range colorValues {
// 		colors[v.Value] = v.ID
// 	}

// 	createProduct := func(name string, price float64, images []string, variants []map[string]interface{}) {
// 		product := models.Product{
// 			ID:            uuid.New(),
// 			Name:          name,
// 			Slug:          strings.ToLower(strings.ReplaceAll(name[:30], " ", "-")),
// 			CategoryID:    category.ID,
// 			SubcategoryID: &sub.ID,
// 			Description:   name,
// 			IsFeatured:    false,
// 			IsActive:      true,
// 		}
// 		db.Create(&product)
// 		for i, img := range images {
// 			db.Create(&models.ProductImage{
// 				ID:        uuid.New(),
// 				ProductID: product.ID,
// 				URL:       img,
// 				IsPrimary: i == 0,
// 			})
// 		}
// 		for _, v := range variants {
// 			varID := uuid.New()
// 			db.Create(&models.ProductVariant{
// 				ID:        varID,
// 				ProductID: product.ID,
// 				SKU:       uuid.NewString(),
// 				Price:     price,
// 				Stock:     v["stock"].(int),
// 				IsActive:  true,
// 				ImageURL:  v["image"].(string),
// 			})
// 			if color, ok := v["color"].(string); ok {
// 				db.Create(&models.ProductVariantOption{
// 					ProductVariantID: varID,
// 					OptionValueID:    colors[color],
// 				})
// 			}
// 		}
// 	}

// 	// Product 1: tanpa varian
// 	createProduct("Unisex Classic Hat", 65000, []string{
// 		"https://s.alicdn.com/@sc04/kf/Hfc996a70f57747738dd5625a897b7195K.jpg_720x720q50.jpg",
// 		"https://s.alicdn.com/@sc04/kf/H16978a217d7e41029ff43eacbfd5fd6c0.jpg_720x720q50.jpg",
// 	}, nil)

// 	// Product 2: blue + gray
// 	createProduct("Stylish Colorful Cap", 75000, []string{
// 		"https://s.alicdn.com/@sc04/kf/A3906e6f105314c8290c9e9bdc83ac51aI.jpg_720x720q50.jpg",
// 		"https://s.alicdn.com/@sc04/kf/A7b4aa8232070439ba0fbd6536019e761m.jpg_720x720q50.jpg",
// 		"https://s.alicdn.com/@sc04/kf/A6f2dc7440ff5485abbaa1ec07b94a929q.jpg_720x720q50.jpg",
// 		"https://s.alicdn.com/@sc04/kf/A9e67963f988744db9c7de7dc444b0007U.jpg_720x720q50.jpg",
// 	}, []map[string]interface{}{
// 		{"color": "blue", "stock": 10, "image": "https://sc04.alicdn.com/kf/Hd72c0ef6e89a43869e3306391e27180c5.jpg"},
// 		{"color": "gray", "stock": 20, "image": "https://sc04.alicdn.com/kf/H3463b32013bb4d1693edea4337ff6a4aP.jpg"},
// 	})

// 	// Product 3: black, red, white
// 	createProduct("Multicolor Embroidered Cap", 89000, []string{
// 		"https://s.alicdn.com/@sc04/kf/Hcdd83d479d6942dcb72e509264dac134J.png_720x720q50.jpg",
// 	}, []map[string]interface{}{
// 		{"color": "black", "stock": 30, "image": "https://s.alicdn.com/@sc04/kf/H521ebe131a5d49ec897756bf8c059871O.jpg"},
// 		{"color": "red", "stock": 15, "image": "https://s.alicdn.com/@sc04/kf/H3f5c17c9b79147d997d9fde2fda989c3p.jpg"},
// 		{"color": "white", "stock": 0, "image": "https://sc04.alicdn.com/kf/Hfae099e8a4024fac872081a35047f00bj.jpg"},
// 	})

// 	// Product 4: tanpa variant
// 	createProduct("Modern Flat Brim Cap", 77500, []string{
// 		"https://s.alicdn.com/@sc04/kf/H9f964f8dbdcd447690136e6ff5fb80f19.png_720x720q50.jpg",
// 		"https://s.alicdn.com/@sc04/kf/H42a9dd56e4d440ab96b40a8c73891a68z.png_720x720q50.jpg",
// 	}, nil)
// }


// func SeedWomensClothing(db *gorm.DB) {
// 	var category models.Category
// 	db.Where("name = ?", "Fashion & Apparel").First(&category)

// 	var sub models.Subcategory
// 	db.Where("name = ? AND category_id = ?", "Women's Clothing", category.ID).First(&sub)

// 	colors := map[string]uint{}
// 	var colorType models.VariantOptionType
// 	db.Where("name = ?", "colors").First(&colorType)
// 	var colorValues []models.VariantOptionValue
// 	db.Where("type_id = ?", colorType.ID).Find(&colorValues)
// 	for _, v := range colorValues {
// 		colors[v.Value] = v.ID
// 	}

// 	createProduct := func(name string, price float64, images []string, variants []map[string]interface{}, stock int) {
// 		product := models.Product{
// 			ID:            uuid.New(),
// 			Name:          name,
// 			Slug:          strings.ToLower(strings.ReplaceAll(name[:30], " ", "-")),
// 			CategoryID:    category.ID,
// 			SubcategoryID: &sub.ID,
// 			Description:   name,
// 			IsFeatured:    false,
// 			IsActive:      true,
// 		}
// 		db.Create(&product)
// 		for i, img := range images {
// 			db.Create(&models.ProductImage{
// 				ID:        uuid.New(),
// 				ProductID: product.ID,
// 				URL:       img,
// 				IsPrimary: i == 0,
// 			})
// 		}
// 		for _, v := range variants {
// 			varID := uuid.New()
// 			db.Create(&models.ProductVariant{
// 				ID:        varID,
// 				ProductID: product.ID,
// 				SKU:       uuid.NewString(),
// 				Price:     price,
// 				Stock:     v["stock"].(int),
// 				IsActive:  true,
// 				ImageURL:  v["image"].(string),
// 			})
// 			if color, ok := v["color"].(string); ok {
// 				db.Create(&models.ProductVariantOption{
// 					ProductVariantID: varID,
// 					OptionValueID:    colors[color],
// 				})
// 			}
// 		}
// 	}

// 	// Product 1
// 	createProduct("Elegant Red and Black Dress", 152000, []string{
// 		"https://s.alicdn.com/@sc04/kf/H6b1532ff613d4fe992a79a8f220218dcy.jpg_720x720q50.jpg",
// 	}, []map[string]interface{}{
// 		{"color": "red", "stock": 20, "image": "https://s.alicdn.com/@sc04/kf/Ha4b04e36c7524776a4c00042d442fd03z.jpg_720x720q50.jpg"},
// 		{"color": "black", "stock": 25, "image": "https://s.alicdn.com/@sc04/kf/H97eef9a0f6f446e499d440fcdea82dd7Q.jpg_720x720q50.jpg"},
// 	}, 0)

// 	// Product 2
// 	createProduct("Stylish Orange and Red Casual Dress", 215000, []string{
// 		"https://s.alicdn.com/@sc04/kf/H08730c53ec1f4c25a6ac260f72ca87fbN.jpg_720x720q50.jpg",
// 	}, []map[string]interface{}{
// 		{"color": "orange", "stock": 20, "image": "https://s.alicdn.com/@sc04/kf/Hb53d9a848df448b3bd9f5cae730bbe33p.jpg_720x720q50.jpg"},
// 		{"color": "red", "stock": 25, "image": "https://s.alicdn.com/@sc04/kf/H525c5ec4f6974984a2731f0cb3c5c51av.jpg_720x720q50.jpg"},
// 	}, 0)

// 	// Product 3
// 	createProduct("Breathable Casual Blouse", 172000, []string{
// 		"https://s.alicdn.com/@sc04/kf/H3988b57157ca44e092a53ed1cb53f187L.jpg_720x720q50.jpg",
// 		"https://s.alicdn.com/@sc04/kf/H17aa21955dd243b2b696e1b29f75dfe2h.jpg_720x720q50.jpg",
// 		"https://s.alicdn.com/@sc04/kf/H79066d0dc2d64e5ba4844bce49c5c53d3.jpg_720x720q50.jpg",
// 	}, nil, 32)

// 	// Product 4
// 	createProduct("Simple Summer Tank Top", 54000, []string{
// 		"https://s.alicdn.com/@sc04/kf/H13950fd4ea254e2d9a81c0ebcf653c157.jpg_720x720q50.jpg",
// 	}, nil, 50)
// }




