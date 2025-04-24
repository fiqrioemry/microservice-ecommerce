package seeders

import (
	"log"
	"strings"

	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/models"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

// no Error - pass
func SeedBanner(db *gorm.DB) {
	banners := []models.Banner{
		// Top Banner
		{ID: uuid.New(), Position: "top", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383472/topbanner03_lgpcf5.webp"},
		{ID: uuid.New(), Position: "top", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383471/topbanner02_supj7d.webp"},
		{ID: uuid.New(), Position: "top", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383470/topbanner01_wvpc7l.webp"},

		// Bottom Banner
		{ID: uuid.New(), Position: "bottom", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383469/bottombanner02_kh2krk.webp"},
		{ID: uuid.New(), Position: "bottom", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383469/bottombanner01_k1lylg.webp"},

		// Side Banner 1
		{ID: uuid.New(), Position: "side1", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner01_gyfi00.webp"},
		{ID: uuid.New(), Position: "side1", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner04_bh6d5e.webp"},

		// Side Banner 2
		{ID: uuid.New(), Position: "side2", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner02_rdtezb.webp"},
		{ID: uuid.New(), Position: "side2", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner03_kraq61.webp"},
	}

	for _, b := range banners {
		if err := db.FirstOrCreate(&b, "image_url = ?", b.ImageURL).Error; err != nil {
			log.Printf("failed to seed banner: %v\n", err)
		}
	}
}

// no Error - pass
func SeedCategoriesAndSubcategories(db *gorm.DB) {
	placeholder := "https://placehold.co/400x400"
	categories := map[string][]string{
		"Fashion & Apparel":     {"Men's Clothing", "Hats and Caps", "Women's Clothing"},
		"Shoes & Accessories":   {"Sandals", "Walking Style Shoes", "Dress Shoes & Oxford"},
		"Gadget & Electronics":  {"Phones & Tablet", "Electronic Devices", "Weareable Devices"},
		"Health & Care":         {"Collagen", "Vitamin", "Sport Nutritions"},
		"Food & Beverage":       {"Health Drink", "Noodle & Pasta", "Snack food"}, // done - 9 product
		"Beauty & Skin Care":    {"Lip Gloss", "Hair Extention", "Make Up"},
		"Sport & Entertainment": {"Cruise Bike", "Baseball", "Roller Wheels"},
	}

	for catName, subs := range categories {
		cat := models.Category{
			ID:    uuid.New(),
			Name:  catName,
			Slug:  strings.ToLower(strings.ReplaceAll(catName, " ", "-")),
			Image: placeholder,
		}

		err := db.Where("name = ?", cat.Name).FirstOrCreate(&cat).Error
		if err != nil {
			log.Println("failed to create category:", catName, err)
			continue
		}

		for _, subName := range subs {
			sub := models.Subcategory{
				ID:         uuid.New(),
				Name:       subName,
				Slug:       strings.ToLower(strings.ReplaceAll(subName, " ", "-")),
				CategoryID: cat.ID,
				Image:      placeholder,
			}

			db.Where("name = ? AND category_id = ?", sub.Name, cat.ID).FirstOrCreate(&sub)
		}
	}
}

// no Error - pass
func SeedVariantTypesAndValues(db *gorm.DB) {
	data := []struct {
		Name   string
		Values []string
	}{
		{"shoes size", []string{"38", "39", "40", "41", "42", "43", "44"}},
		{"clothing size", []string{"S", "M", "L", "XL"}},
		{"colors", []string{"red", "green", "gray", "pink", "blue", "black", "white", "brown", "orange"}},
		{"ram capacity", []string{"4gb", "6gb", "8gb", "12gb"}},
		{"memory capacity", []string{"64gb", "128gb", "256gb", "512gb"}},
		{"screen size", []string{"12\"", "14\"", "18\"", "20\"", "24\"", "30\""}},
	}

	for _, item := range data {
		typeModel := models.VariantOptionType{Name: item.Name}
		if err := db.Where("name = ?", item.Name).FirstOrCreate(&typeModel).Error; err != nil {
			log.Printf("failed to create variant type: %s", item.Name)
			continue
		}

		for _, val := range item.Values {
			valModel := models.VariantOptionValue{
				TypeID: typeModel.ID,
				Value:  val,
			}
			db.Where("type_id = ? AND value = ?", typeModel.ID, val).FirstOrCreate(&valModel)
		}
	}
}

func SeedProductDataOne(db *gorm.DB) {
	products := []struct {
		Category    string
		Subcategory string
		Description string
		Name        string
		IsFeatured  bool
		Discount    float64
		Images      []string
		Variants    []struct {
			Color string
			Size  string
			Price float64
			Stock int
			Image string
		}
	}{
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Men's Clothing",
			Description: "Pakaian softshell merupakan pakaian serba guna. Sebagai sentuhan baru pada pakaian luar hardshell klasik, bahan ini menawarkan pengalaman yang lebih fleksibel sehingga cocok untuk olahraga, pakaian olahraga, golf, dan bahkan pakaian sehari-hari.",
			Name:        "Men's Soft Shell Assault Jacket",
			IsFeatured:  false,
			Discount:    0.0,
			Images: []string{
				"https://s.alicdn.com/@sc04/kf/H7a6558c2616241aaafb4e6e119042238C.png_720x720q50.jpg",
				"https://s.alicdn.com/@sc04/kf/H8ed62347008341f0a6927543e7cc5504c.jpg_720x720q50.jpg",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"black", "M", 265000, 5, ""},
				{"black", "L", 265000, 10, ""},
				{"white", "M", 265000, 5, ""},
				{"white", "L", 265000, 10, ""},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Men's Clothing",
			Name:        "High Quality Men's Youth Fashion Jacket",
			IsFeatured:  false,
			Discount:    0.0,
			Images: []string{
				"https://s.alicdn.com/@sc04/kf/H45c2134872bc4b6e89ca5b55db292019q.jpg_720x720q50.jpg",
				"https://s.alicdn.com/@sc04/kf/Hf54245977fbb469cbc3282fd31d2d65dl.jpg?avif=close",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"red", "M", 335000, 5, ""},
				{"red", "L", 335000, 10, ""},
				{"black", "M", 335000, 5, ""},
				{"black", "L", 335000, 10, ""},
				{"blue", "M", 335000, 5, ""},
				{"blue", "L", 335000, 10, ""},
				{"blue", "XL", 335000, 15, ""},
				{"blue", "XXL", 335000, 20, ""},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Men's Clothing",
			Name:        "Men's High Quality 300 Gsm Boxy Graphic Oversized",
			IsFeatured:  false,
			Discount:    0.0,
			Images: []string{
				"https://s.alicdn.com/@sc04/kf/H37fdf58f0dd54cdca5a086055fe1e101q.jpg_720x720q50.jpg",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"black", "M", 177000, 5, "https://s.alicdn.com/@sc04/kf/H37fdf58f0dd54cdca5a086055fe1e101q.jpg_720x720q50.jpg"},
				{"black", "L", 177000, 10, "https://s.alicdn.com/@sc04/kf/H37fdf58f0dd54cdca5a086055fe1e101q.jpg_720x720q50.jpg"},
				{"blue", "M", 177000, 5, " https://s.alicdn.com/@sc04/kf/Hcf176a6f30b840a89e5f0a3c60bb1ecec.jpg_720x720q50.jpg"},
				{"blue", "L", 177000, 10, " https://s.alicdn.com/@sc04/kf/Hcf176a6f30b840a89e5f0a3c60bb1ecec.jpg_720x720q50.jpg"},
				{"blue", "XL", 177000, 15, " https://s.alicdn.com/@sc04/kf/Hcf176a6f30b840a89e5f0a3c60bb1ecec.jpg_720x720q50.jpg"},
				{"pink", "M", 177000, 5, " https://s.alicdn.com/@sc04/kf/H32a280fda84440be8c22f209b589a36bJ.jpg_720x720q50.jpg"},
				{"pink", "L", 177000, 10, " https://s.alicdn.com/@sc04/kf/H32a280fda84440be8c22f209b589a36bJ.jpg_720x720q50.jpg"},
				{"pink", "XL", 177000, 15, " https://s.alicdn.com/@sc04/kf/H32a280fda84440be8c22f209b589a36bJ.jpg_720x720q50.jpg"},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Men's Clothing",
			Name:        "Business Casual Tuxedo Men's ",
			IsFeatured:  false,
			Discount:    0.0,
			Images: []string{
				"https://s.alicdn.com/@sc04/kf/H1db0a8f8f25041dab98442eaa144ff3ay.jpg_720x720q50.jpg",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"black", "M", 275000, 5, "https://s.alicdn.com/@sc04/kf/H5722ff2a7e144e5aa83103a52770bed0G.jpg_720x720q50.jpg"},
				{"black", "L", 275000, 10, "https://s.alicdn.com/@sc04/kf/H5722ff2a7e144e5aa83103a52770bed0G.jpg_720x720q50.jpg"},
				{"red", "M", 275000, 10, "https://s.alicdn.com/@sc04/kf/H07b065e7028640488feeb1278c3cff09M.jpg_720x720q50.jpg"},
				{"red", "L", 275000, 20, " https://s.alicdn.com/@sc04/kf/H07b065e7028640488feeb1278c3cff09M.jpg_720x720q50.jpg"},
				{"red", "XL", 275000, 30, " https://s.alicdn.com/@sc04/kf/H07b065e7028640488feeb1278c3cff09M.jpg_720x720q50.jpg"},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Hats and Caps",
			Name:        "High Quality Fashion Street Versatile Baseball",
			IsFeatured:  false,
			Discount:    0.1,
			Images: []string{
				"https://s.alicdn.com/@sc04/kf/Hfc996a70f57747738dd5625a897b7195K.jpg_720x720q50.jpg",
				"https://s.alicdn.com/@sc04/kf/H16978a217d7e41029ff43eacbfd5fd6c0.jpg_720x720q50.jpg",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "", 65000, 50, ""},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Hats and Caps",
			Name:        "TCAP Custom 5 Panel PVC Rubber Patch Waterproof",
			IsFeatured:  false,
			Discount:    0.1,
			Images: []string{
				"https://s.alicdn.com/@sc04/kf/A3906e6f105314c8290c9e9bdc83ac51aI.jpg_720x720q50.jpg",
				"https://s.alicdn.com/@sc04/kf/A7b4aa8232070439ba0fbd6536019e761m.jpg_720x720q50.jpg",
				"https://s.alicdn.com/@sc04/kf/A6f2dc7440ff5485abbaa1ec07b94a929q.jpg_720x720q50.jpg",
				"https://s.alicdn.com/@sc04/kf/A9e67963f988744db9c7de7dc444b0007U.jpg_720x720q50.jpg",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"blue", "", 87000, 20, "https://sc04.alicdn.com/kf/Hd72c0ef6e89a43869e3306391e27180c5.jpg"},
				{"gray", "", 87000, 20, "https://sc04.alicdn.com/kf/H3463b32013bb4d1693edea4337ff6a4aP.jpg"},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Hats and Caps",
			Name:        "New Vintage Street Baseball Caps Custom",
			IsFeatured:  false,
			Discount:    0.1,
			Images: []string{
				"https://s.alicdn.com/@sc04/kf/Hcdd83d479d6942dcb72e509264dac134J.png_720x720q50.jpg",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"black", "", 102000, 10, "https://s.alicdn.com/@sc04/kf/H521ebe131a5d49ec897756bf8c059871O.jpg_720x720q50.jpg"},
				{"red", "", 102000, 20, "https://s.alicdn.com/@sc04/kf/H3f5c17c9b79147d997d9fde2fda989c3p.jpg_720x720q50.jpg"},
				{"white", "", 102000, 30, "https://sc04.alicdn.com/kf/Hfae099e8a4024fac872081a35047f00bj.jpg"},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Hats and Caps",
			Name:        "Factory Price Classic Adjustable Women and Men",
			IsFeatured:  false,
			Discount:    0.1,
			Images: []string{
				"https://s.alicdn.com/@sc04/kf/Hcdd83d479d6942dcb72e509264dac134J.png_720x720q50.jpg",
				"https://s.alicdn.com/@sc04/kf/H9f964f8dbdcd447690136e6ff5fb80f19.png_720x720q50.jpg",
				"https://s.alicdn.com/@sc04/kf/H42a9dd56e4d440ab96b40a8c73891a68z.png_720x720q50.jpg",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "", 77500, 50, ""},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Women's Clothing",
			Name:        "Autumn Allover Print Twist High Waist Bodycon Skirts",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://s.alicdn.com/@sc04/kf/H6b1532ff613d4fe992a79a8f220218dcy.jpg_720x720q50.jpg",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"red", "", 152000, 10, "https://s.alicdn.com/@sc04/kf/Ha4b04e36c7524776a4c00042d442fd03z.jpg_720x720q50.jpg"},
				{"black", "", 152000, 20, "https://s.alicdn.com/@sc04/kf/H97eef9a0f6f446e499d440fcdea82dd7Q.jpg_720x720q50.jpg"},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Women's Clothing",
			Name:        "Women Clothing Dress Lady Elegant Temperament",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://s.alicdn.com/@sc04/kf/H08730c53ec1f4c25a6ac260f72ca87fbN.jpg_720x720q50.jpg",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"orange", "", 215000, 10, "https://s.alicdn.com/@sc04/kf/Hb53d9a848df448b3bd9f5cae730bbe33p.jpg_720x720q50.jpg"},
				{"red", "", 215000, 20, "https://s.alicdn.com/@sc04/kf/H525c5ec4f6974984a2731f0cb3c5c51av.jpg_720x720q50.jpg"},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Women's Clothing",
			Name:        "Spring Fashion Casual Women's Clothes Ladies",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://s.alicdn.com/@sc04/kf/H3988b57157ca44e092a53ed1cb53f187L.jpg_720x720q50.jpg",
				"https://s.alicdn.com/@sc04/kf/H17aa21955dd243b2b696e1b29f75dfe2h.jpg_720x720q50.jpg",
				"https://s.alicdn.com/@sc04/kf/H79066d0dc2d64e5ba4844bce49c5c53d3.jpg_720x720q50.jpg",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "", 172000, 10, ""},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Women's Clothing",
			Name:        "Spring Fashion Casual Women's Clothes Ladies",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://s.alicdn.com/@sc04/kf/H9bc860bc18094c18892a3973003cbde8W.png_720x720q50.jpg",
				"https://s.alicdn.com/@sc04/kf/H06b3dc8823ad4a568f015e193ce3d131j.png_720x720q50.jpg",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "", 97000, 20, ""},
			},
		},
	}

	for _, p := range products {
		var cat models.Category
		db.Where("name = ?", p.Category).First(&cat)

		var sub models.Subcategory
		db.Where("name = ? AND category_id = ?", p.Subcategory, cat.ID).First(&sub)

		product := models.Product{
			ID:            uuid.New(),
			CategoryID:    cat.ID,
			SubcategoryID: &sub.ID,
			Name:          p.Name,
			Slug:          strings.ToLower(strings.ReplaceAll(p.Name, " ", "-")),
			IsFeatured:    p.IsFeatured,
			Discount:      &p.Discount,
			IsActive:      true,
		}
		db.Create(&product)

		for i, img := range p.Images {
			db.Create(&models.ProductImage{
				ID:        uuid.New(),
				ProductID: product.ID,
				URL:       img,
				IsPrimary: i == 0,
			})
		}

		for _, v := range p.Variants {
			var colorVal, sizeVal models.VariantOptionValue
			db.Where("value = ?", v.Color).First(&colorVal)
			db.Where("value = ?", v.Size).First(&sizeVal)

			variant := models.ProductVariant{
				ID:        uuid.New(),
				ProductID: product.ID,
				Price:     v.Price,
				Stock:     v.Stock,
				ImageURL:  v.Image,
			}
			db.Create(&variant)

			db.Create(&models.ProductVariantOption{
				ProductVariantID: variant.ID,
				OptionValueID:    colorVal.ID,
			})
			db.Create(&models.ProductVariantOption{
				ProductVariantID: variant.ID,
				OptionValueID:    sizeVal.ID,
			})
		}
	}
}

func SeedFoodBeverage(db *gorm.DB) {
	products := []struct {
		Category    string
		Subcategory string
		Description string
		Name        string
		IsFeatured  bool
		Discount    float64
		Weight      float64
		Length      float64
		Width       float64
		Height      float64
		Images      []string
		Variants    []struct {
			Color string
			Size  string
			Price float64
			Stock int
			Image string
		}
	}{
		{
			Category:    "Food & Beverage",
			Subcategory: "Snack Food",
			Name:        "HOTTO PURTO 1 POUCH 16 SACHET | Superfood Multigrain Purple Potato Oat",
			Description: "Hotto Purto, merupakan minuman kesehatan tinggi serat yang kaya akan nutrisi dan rendah kalori. Diformulasikan secara khusus dengan bahan-bahan premium seperti ubi ungu, oat dari Swedia, serta 15 biji-bijian (multigrain). Merupakan pilihan yang tepat untuk dijadikan sarapan praktis untuk keluarga tercinta. 15 MULTIGRAIN Menurut penelitian, pola makan tidak sehat membunuh 11 juta orang di dunia pertahunnya. Kurangnya konsumsi biji-bijian dan kacang-kacangan menjadi salah satu penyebab terbesar kematiannya. Hotto mengandung 15 jenis biji-bijian yang menjadikannya sebagai sumber nutrisi, mineral, protein dan kaya akan serat. ",
			IsFeatured:  false,
			Discount:    0.0,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424592/hoto_snack_01_lf8uml.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424593/hoto_snack_02_sek5gt.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424599/hoto_snack_03_six5wh.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "", 135000, 50, ""},
			},
		},
		{
			Category:    "Food & Beverage",
			Subcategory: "Snack Food",
			Name:        "Covita - Healthy Protein Bar 40 gr Gluten Free - Peanut Choco",
			Description: "Cemilan sehat berprotein (Plant-Based) atau cemilan untuk kegiatan olahraga. Bersumber dari bahan protein alami untuk sebelum dan sesudah berolahraga. 15 MULTIGRAIN Menurut penelitian, pola makan tidak sehat membunuh 11 juta orang di dunia pertahunnya. Kurangnya konsumsi biji-bijian dan kacang-kacangan menjadi salah satu penyebab terbesar kematiannya. Hotto mengandung 15 jenis biji-bijian yang menjadikannya sebagai sumber nutrisi, mineral, protein dan kaya akan serat",
			IsFeatured:  false,
			Discount:    0.0,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424765/bars_snack_01_ghf8uj.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424766/bars_snack_02_nsbgth.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424767/bars_snack_03_vcsloc.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "", 67000, 50, ""},
			},
		},
		{
			Category:    "Food & Beverage",
			Subcategory: "Snack Food",
			Name:        "Covita - Peach Gum Collagen Dessert with Tangerine",
			Description: "Peach Gum Collagen Dessert with Tangerine adalah hidangan penutup yang populer, terutama di Cina, yang kaya akan kolagen dan manfaat kesehatan lainnya. Peach gum, yang terbuat dari getah pohon persik liar, mengandung kolagen dan asam amino yang tinggi, serta manfaat lainnya seperti melancarkan pencernaan, meningkatkan stamina, dan menjaga kesehatan kulit. Penambahan buah tangerine memberikan rasa segar dan aroma yang menyenangkan.",
			IsFeatured:  false,
			Discount:    0.0,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425054/grain_snack_01_hurkzb.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425055/grain_snack_02_cnqxkk.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425057/grain_snack_03_sm9sze.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "", 87000, 100, ""},
			},
		},
		{
			Category:    "Food & Beverage",
			Subcategory: "Health Drink",
			Name:        "Madu Asli Hutan Honey Life Gold 650ml",
			Description: "Madu Honey Life merupakan Spesialis focus madu Segar hutan liar ( bukan Ternak ) mentah, murni dari Alam dan Organik. ...",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425496/honey_drink_01_qjl69j.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425499/honey_drink_02_dyufai.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "", 168000, 30, ""},
			},
		},
		{
			Category:    "Food & Beverage",
			Subcategory: "Health Drink",
			Name:        "Nestle Pure Life Air Minum Ukuran 600mL - 1 Pack",
			Description: "Air minum Nestle Pure Life 600mL adalah air mineral ...",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425497/nestle_drink_02_bd5mye.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425501/nestle_drink_01_vgnua8.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "", 115000, 30, ""},
			},
		},
		{
			Category:    "Food & Beverage",
			Subcategory: "Health Drink",
			Name:        "ESSENLI Pure Matcha Powder Japan / Bubuk Matcha Murni Drink",
			Description: "ESSENLI Pure Matcha Powder Japan adalah bubuk matcha murni ...",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425829/matcha_drink_01_nq1pzd.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425832/matcha_drink_02_nviqwj.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425827/matcha_drink_03_y1mbxw.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "", 75500, 30, ""},
			},
		},
		{
			Category:    "Food & Beverage",
			Subcategory: "Noodle & Pasta",
			Name:        "Mie Porang dietmeal GORENG rendah kalori",
			Description: "Mie Porang dietmeal GORENG rendah kalori adalah varian mie yang terbuat dari umbi porang, cocok untuk mereka yang sedang diet atau ingin menjaga berat badan karena rendah kalori dan bebas gluten. Mie ini juga tinggi serat, membantu menjaga kesehatan pencernaan dan memberikan efek kenyang lebih lama. Terbuat dari tepung porang, umbi-umbian yang rendah kalori dan tinggi serat dan tidak mengandung gluten, sehingga aman dikonsumsi oleh orang dengan intoleransi gluten.",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426605/indomie_noodle_02_leaptj.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426601/indomie_noodle_01_wztuyg.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "", 8900, 1500, ""},
			},
		},
		{
			Category:    "Food & Beverage",
			Subcategory: "Noodle & Pasta",
			Name:        "Bihunku All Rasa Soto Nyus",
			Description: "Bihunku All Rasa adalah bihun instan yang lezat, mudah dimasak, dan cocok untuk disantap kapan saja dan di mana saja. Bihunku terbuat dari perpaduan beras dan tepung jagung pilihan, dengan bumbu alami yang khas. Produk ini rendah lemak dan kolesterol, serta mengandung serat yang membuat kenyang tahan lama. Tersedia dalam berbagai varian rasa, seperti Ayam Bawang, Soto Spesial, Goreng Spesial, dan lainnya. ",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426599/bihun_noodle_02_ibzcpd.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426611/bihun_noodle_01_t0egqo.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "", 11600, 1500, ""},
			},
		},
		{
			Category:    "Food & Beverage",
			Subcategory: "Noodle & Pasta",
			Name:        "ORIMIE Goreng dari Orimen Kids",
			Description: "ORIMIE Goreng dari Orimen Kids adalah pilihan mie goreng yang lebih sehat untuk anak-anak, tanpa pewarna, pengawet, atau MSG. Mie ini juga bebas dari babi dan minyak babi. Bumbunya dibuat dengan bahan-bahan alami, seperti bubuk bawang putih, bawang merah, garam, gula, kaldu, dan lada putih, serta minyak yang dibumbui dengan daun bawang, bawang putih, kulit ayam, dan bawang merah",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426605/orime_noodle_01_bpuprf.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426606/orime_noodle_02_yjx3u0.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426610/orime_noodle_03_k8ljlt.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "", 23500, 1500, ""},
			},
		},
	}

	for _, p := range products {
		var cat models.Category
		db.Where("name = ?", p.Category).First(&cat)

		var sub models.Subcategory
		db.Where("name = ? AND category_id = ?", p.Subcategory, cat.ID).First(&sub)

		product := models.Product{
			ID:            uuid.New(),
			CategoryID:    cat.ID,
			SubcategoryID: &sub.ID,
			Name:          p.Name,
			Weight:        1000.0,
			Width:         15.0,
			Height:        15.0,
			Length:        15.0,
			Slug:          strings.ToLower(strings.ReplaceAll(p.Name, " ", "-")),
			IsFeatured:    p.IsFeatured,
			Discount:      &p.Discount,
			IsActive:      true,
		}
		db.Create(&product)

		for i, img := range p.Images {
			db.Create(&models.ProductImage{
				ID:        uuid.New(),
				ProductID: product.ID,
				URL:       img,
				IsPrimary: i == 0,
			})
		}

		for _, v := range p.Variants {
			var colorVal, sizeVal models.VariantOptionValue
			db.Where("value = ?", v.Color).First(&colorVal)
			db.Where("value = ?", v.Size).First(&sizeVal)

			variant := models.ProductVariant{
				ID:        uuid.New(),
				ProductID: product.ID,
				Price:     v.Price,
				Stock:     v.Stock,
				ImageURL:  v.Image,
			}
			db.Create(&variant)

			db.Create(&models.ProductVariantOption{
				ProductVariantID: variant.ID,
				OptionValueID:    colorVal.ID,
			})
			db.Create(&models.ProductVariantOption{
				ProductVariantID: variant.ID,
				OptionValueID:    sizeVal.ID,
			})
		}
	}
}

func SeedGadgetElectronic(db *gorm.DB) {
	products := []struct {
		Category    string
		Subcategory string
		Description string
		Name        string
		IsFeatured  bool
		Discount    float64
		Weight      float64
		Length      float64
		Width       float64
		Height      float64
		Images      []string
		Variants    []struct {
			Color string
			Size  string
			Price float64
			Stock int
			Image string
		}
	}{
		// write code here ......
	}

	for _, p := range products {
		var cat models.Category
		db.Where("name = ?", p.Category).First(&cat)

		var sub models.Subcategory
		db.Where("name = ? AND category_id = ?", p.Subcategory, cat.ID).First(&sub)

		product := models.Product{
			ID:            uuid.New(),
			CategoryID:    cat.ID,
			SubcategoryID: &sub.ID,
			Name:          p.Name,
			Weight:        1000.0,
			Width:         15.0,
			Height:        15.0,
			Length:        15.0,
			Slug:          strings.ToLower(strings.ReplaceAll(p.Name, " ", "-")),
			IsFeatured:    p.IsFeatured,
			Discount:      &p.Discount,
			IsActive:      true,
		}
		db.Create(&product)

		for i, img := range p.Images {
			db.Create(&models.ProductImage{
				ID:        uuid.New(),
				ProductID: product.ID,
				URL:       img,
				IsPrimary: i == 0,
			})
		}

		for _, v := range p.Variants {
			var colorVal, sizeVal models.VariantOptionValue
			db.Where("value = ?", v.Color).First(&colorVal)
			db.Where("value = ?", v.Size).First(&sizeVal)

			variant := models.ProductVariant{
				ID:        uuid.New(),
				ProductID: product.ID,
				Price:     v.Price,
				Stock:     v.Stock,
				ImageURL:  v.Image,
			}
			db.Create(&variant)

			db.Create(&models.ProductVariantOption{
				ProductVariantID: variant.ID,
				OptionValueID:    colorVal.ID,
			})
			db.Create(&models.ProductVariantOption{
				ProductVariantID: variant.ID,
				OptionValueID:    sizeVal.ID,
			})
		}
	}
}
