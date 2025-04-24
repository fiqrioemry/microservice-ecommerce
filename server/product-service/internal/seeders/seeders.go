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
		"Fashion & Apparel": {"Men's Clothing", "Women's Skirt", "Men's Pants", "Women's Dress"},
		// "Shoes & Accessories":  {"Sandals", "Walking Style Shoes", "Dress Shoes & Oxford"},
		// "Health & Care":        {"Collagen", "Vitamin", "Sport Nutritions"},
		"Gadget & Electronics": {"Phones & Tablet", "Electronic Devices", "Weareable Devices"}, // done - 9 product
		"Food & Beverage":      {"Health Drink", "Noodle & Pasta", "Snack food"},               // done - 9 product
		// "Beauty & Skin Care":   {"Lip Gloss", "Hair Extention", "Make Up"},
		// "Sport & Entertainment": {"Cruise Bike", "Baseball", "Roller Wheels"},
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

func SeedFashionAndApparel(db *gorm.DB) {
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
			Name:        "Jacket Denim Warna Biru Bahan Ekslusif",
			Description: "Jaket denim warna biru dongker adalah jaket yang terbuat dari bahan denim yang memiliki warna biru tua...",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745429277/erem_shirt_01_shijri.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745429275/erem_shirt_02_dusksh.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745429265/erem_shirt_03_ykizqa.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "M", 315000, 10, ""},
				{"", "L", 315000, 20, ""},
				{"", "XL", 315000, 30, ""},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Men's Clothing",
			Name:        "Kaos Distro Pria Lengan Pendek NY Kaos Oblong Cowok",
			Description: "Kaos Distro Pria Lengan Pendek NY Kaos Oblong Cowok adalah jenis kaos yang diproduksi dengan jumlah terbatas...",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509051/cloth_mens_01_l4sqob.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509051/cloth_mens_02_rzapkt.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"white", "M", 98500, 10, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509050/cloth_mens_04_xttcat.webp"},
				{"white", "L", 98500, 20, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509050/cloth_mens_04_xttcat.webp"},
				{"gray", "M", 98500, 10, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509053/cloth_mens_03_nwcb4c.webp"},
				{"gray", "L", 98500, 20, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509053/cloth_mens_03_nwcb4c.webp"},
				{"gray", "XL", 98500, 30, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509053/cloth_mens_03_nwcb4c.webp"},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Men's Clothing",
			Name:        "Hoodie Addict - Zipper Hoodie Dewasa Polos Hitam Pria",
			Description: "Hoodie Addict Zipper adalah jaket hoodie dengan ritsleting (zipper) yang populer...",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509457/jaket01_tld8i0.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509457/jaket02_ru71to.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509458/jaket03_ygtnw2.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "M", 138500, 10, ""},
				{"", "L", 138500, 20, ""},
				{"", "XL", 138500, 30, ""},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Men's Clothing",
			Name:        "Hoodie Boxy Oversize Men Decorder Gray",
			Description: "Hoodie boxy oversize adalah hoodie dengan siluet yang lebih lebar dan berbentuk kotak (boxy)...",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509054/jaket_mens_02_tyjlul.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509053/jaket_mens_01_a21ye5.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "M", 112500, 10, ""},
				{"", "L", 112500, 20, ""},
				{"", "XL", 112500, 30, ""},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Women's Dress",
			Name:        "Elegant Floral Summer Dress Blossom",
			Description: "Dress ini dirancang untuk memberikan kesan anggun dan modern bagi setiap wanita. Menggunakan bahan berkualitas tinggi yang ringan dan nyaman dipakai sepanjang hari. Potongannya mengikuti lekuk tubuh dengan elegan namun tetap memberikan kenyamanan.",
			IsFeatured:  false,
			Discount:    0.0,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510300/dress01_w1clnu.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510301/dress02_xnlphu.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510304/dress03_d3y08s.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "M", 199000, 10, ""},
				{"", "L", 199000, 20, ""},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Women's Dress",
			Name:        "Chic Long Sleeve Bodycon Dress",
			Description: "Didesain dengan gaya timeless yang tak lekang oleh tren. Panjang rok yang midi membuatnya tetap sopan namun tetap stylish. Dress ini dirancang untuk memberikan kesan anggun dan modern bagi setiap wanita. Bagian pinggang dibuat elastis untuk fleksibilitas ukuran dan kenyamanan ekstra.",
			IsFeatured:  false,
			Discount:    0.0,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510300/wom_dress03_bqsuif.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510299/wom_dress02_susije.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510299/wom_dress01_zgzscq.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "M", 245000, 15, ""},
				{"", "L", 245000, 20, ""},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Men's Pants",
			Name:        "Malvose Celana Pria Formal Bahan Premium Black Slimfit",
			Description: "Celana Pria Formal Bahan Premium Black Slimfit adalah celana formal dengan potongan slimfit yang terbuat dari bahan premium. Celana ini cocok untuk berbagai acara formal, semi formal, dan bahkan kasual, seperti ke kantor atau kondangan. ",
			IsFeatured:  false,
			Discount:    0.15,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510924/pants01_x4memd.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510925/pants02_cloota.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510925/pants03_rx1ixk.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "M", 345000, 10, ""},
				{"", "L", 345000, 20, ""},
				{"", "XL", 345000, 30, ""},
			},
		},
		{
			Category:    "Fashion & Apparel",
			Subcategory: "Men's Pants",
			Name:        "celana cargo panjang pria celana outdoor pria longgar kasual korduroi kulot",
			Description: "Celana cargo panjang pria ini adalah pilihan ideal untuk kegiatan outdoor, dikarenakan desainnya yang longgar dan kasual, serta dilengkapi dengan saku-saku besar di samping (cargo pockets). Bahan korduroi memberikan kesan unik dan nyaman, cocok untuk berbagai aktivitas, termasuk kulot.",
			IsFeatured:  false,
			Discount:    0.15,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510904/men_pants01_tgqmbn.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510916/men_pants02_yjdzug.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "M", 215000, 10, ""},
				{"", "L", 215000, 20, ""},
				{"", "XL", 215000, 30, ""},
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
			Description: "Air minum Nestle Pure Life 600mL adalah air mineral yang diproduksi dengan Standar Internasional oleh Nestle Global Waters. Tersebar diberbagai negara di dunia, air minum Nestle Pure Life tersedia di lebih dari 40 negara di dunia dan menjadi Top 3 di 13 negara. Selain cocok untuk memenuhi kebutuhan hidrasi kamu dan keluarga setiap hari. Air minum Nestle Pure Life juga cocok untuk kamu yang membutuhkan air minum dengan kesegaran nyata dalam kondisi apapun.",
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
			Name:        "ESSENLI Pure Matcha Powder Japan Bubuk Matcha Murni Drink",
			Description: "ESSENLI Pure Matcha Powder Japan adalah bubuk matcha murni (bubuk teh hijau Jepang) yang dikeringkan dengan metode khusus dan digiling menjadi bubuk halus. Matcha ini kaya akan antioksidan, seperti polifenol dan EGCG, serta berbagai nutrisi seperti protein, gula, vitamin, dan mineral. ESSENLI Pure Matcha Powder Japan bisa digunakan untuk berbagai macam minuman, makanan, dan bahkan untuk membuat masker wajah. Contohnya adalah untuk membuat matcha latte, matcha ice cream, matcha cake, matcha pasta, dan sebagainya.",
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
		{
			Category:    "Gadget & Electronics",
			Subcategory: "Phone & Tablet",
			Name:        "Motorola G45 Snapdragon 6s Gen 3",
			Description: "Moto G45 5G pakai prosesor Qualcomm Snapdragon 6s Gen 3. Prosesor ini andal untuk menjalankan aplikasi-aplikasi secara bersamaan, membuat multi-tasking dapat dilakukan tanpa lag, sekaligus hemat daya. Performanya didukung oleh konfigurasi RAM 8 GB fisik + 8 GB RAM virtual (Extended RAM) dan penyimpanan internal 256 GB.",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421821/motorola_phone_01_hpmjaf.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421821/motorola_phone_03_pbvpd1.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421820/motorola_phone_02_wqlrdz.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "4gb", 1250000, 10, ""},
				{"", "6gb", 1350000, 20, ""},
				{"", "8gb", 1450000, 30, ""},
			},
		},
		{
			Category:    "Gadget & Electronics",
			Subcategory: "Phone & Tablet",
			Name:        "Samsung Galaxy A16 - Garansi Resmi Sein Tam",
			Description: "Samsung Galaxy A16 adalah smartphone Android yang menawarkan kombinasi layar Super AMOLED 6,7 inci, baterai 5000mAh, dan kamera 50MP. Perangkat ini memiliki desain tipis dengan ketebalan 7,9mm. Samsung Galaxy A16 tersedia dalam beberapa pilihan memori internal dan RAM, serta dilengkapi dengan fitur Super Fast Charging.",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421821/motorola_phone_01_hpmjaf.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421821/motorola_phone_03_pbvpd1.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421820/motorola_phone_02_wqlrdz.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "4gb", 2599999, 10, ""},
				{"", "6gb", 2699999, 20, ""},
				{"", "8gb", 2799999, 30, ""},
			},
		},
		{
			Category:    "Gadget & Electronics",
			Subcategory: "Phone & Tablet",
			Name:        "Asus Zenfone 11 Ultra 12 5G",
			Description: "Asus Zenfone 11 Ultra 12/256GB adalah smartphone flagship dengan layar 6.78 inci AMOLED, chipset Snapdragon 8 Gen 3, RAM 12GB, storage 256GB, dan baterai 5500 mAh. Perangkat ini memiliki kamera belakang 50MP utama dan 32MP telephoto, serta kamera depan 32MP. Zenfone 11 Ultra juga dilengkapi dengan fitur 6-axis hybrid gimbal untuk video yang stabil.",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423036/asus_phone_05_bgoxso.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423035/asus_phone_04_qe1lqw.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"black", "256gb", 8499000, 10, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745422573/asus_phone_01_wyvgsx.webp"},
				{"black", "512gb", 8899000, 20, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745422573/asus_phone_01_wyvgsx.webp"},
				{"grey", "256gb", 8499000, 10, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745422573/asus_phone_03_ptjmet.webp"},
				{"grey", "512gb", 8899000, 20, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745422573/asus_phone_03_ptjmet.webp"},
				{"blue", "256gb", 8499000, 10, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745422573/asus_phone_02_mbvwyi.webp"},
				{"blue", "512gb", 8899000, 20, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745422573/asus_phone_02_mbvwyi.webp"},
			},
		},
		{
			Category:    "Gadget & Electronics",
			Subcategory: "Phone & Tablet",
			Name:        "Infinix XPad 11 Tablet 5G Premium",
			Description: "Infinix XPad 11 adalah tablet Android dengan layar 11 inci dan refresh rate 90Hz, ditenagai oleh chipset MediaTek Helio G99. 7000mAh, RAM hingga 8GB, dan Android 14. Ia juga dilengkapi dengan fitur-fitur seperti Folax Voice Assistant, Multi-Device Collaboration, dan pengisian cepat.",
			IsFeatured:  true,
			Discount:    0.15,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423645/infinix_tablet_01_mh0wgd.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423643/infinix_tablet_02_fptycg.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "4gb", 2250000, 10, ""},
				{"", "8gb", 2350000, 20, ""},
			},
		},
		{
			Category:    "Gadget & Electronics",
			Subcategory: "Phone & Tablet",
			Name:        "Huawei MatePad 11 Snapdragon 865",
			Description: "Huawei MatePad 11 adalah tablet dengan layar 11 inci, ditenagai oleh chipset Snapdragon 865, RAM 6GB, dan memori internal 128GB yang dapat diperluas. Tablet ini juga dilengkapi dengan sistem operasi Harmony OS 3.1. Secara keseluruhan, Huawei MatePad 11 adalah tablet yang menawarkan performa baik, layar yang bagus, dan berbagai fitur tambahan, menjadikannya pilihan yang menarik untuk berbagai kebutuhan, mulai dari produktivitas hingga hiburan.",
			IsFeatured:  true,
			Discount:    0.1,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423869/huawei_tablet_01_qz7bbi.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423859/huawei_tablet_03_qbokzz.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423858/huawei_tablet_02_twk4ey.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "6gb", 5500000, 10, ""},
				{"", "8gb", 5900000, 20, ""},
			},
		},
		{
			Category:    "Gadget & Electronics",
			Subcategory: "Phone & Tablet",
			Name:        "Xiaomi Pad SE NEW Garansi",
			Description: "Xiaomi Redmi Pad SE adalah tablet Android yang memiliki layar FHD+ 11 inci dengan refresh rate 90 Hz, ditenagai oleh prosesor Snapdragon 680, RAM 4GB, dan penyimpanan internal 128GB, serta baterai 8000mAh. Tablet ini dilengkapi dengan empat speaker dengan Dolby Atmos, dan kamera depan 5MP dan kamera belakang 8MP. Redmi Pad SE hadir dengan layar IPS LCD berukuran 10,1 inci, memberikan tampilan yang luas dan jelas. Resolusi layar sebesar 1200 x 2000 piksel, dengan tingkat kecerahan hingga 340 nits dan rasio kontras 1500:1, cocok untuk berbagai kebutuhan mulai dari streaming video, browsing, hingga bermain game.",
			IsFeatured:  true,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424296/xiaomi_tablet_02_oxh1ad.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424295/xiaomi_tablet_01_wkjuec.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "6gb", 2450000, 10, ""},
				{"", "8gb", 275000, 20, ""}, // ← Perlu dicek, kemungkinan typo (mungkin maksudnya 2.750.000?)
			},
		},
		{
			Category:    "Gadget & Electronics",
			Subcategory: "Wearable Devices",
			Name:        "Xiaomi Mi band 4 Smartwatch",
			Description: "Miliki smartband pintar xiaomi Mi Band 4 Generasi terbaru, hadir dengan beragam fitur canggih dengan peningkatan yang lebih baik dari generasi sebelumnya. Kapasitas baterai Xiaomi Mi Band 4 50 % lebih besar dari xiaomi mi band 2 yang mampu bertahan hingga lebih dari 20 hari penggunaan. XIaomi Mi Band 4 dilengkapi dengan bluetooth 4.2 untuk konektivitasnya dan untuk ketahanan airnya pun turut ditingkatkan yang kini mampu bertahan hingga kedalaman 50 meter.",
			IsFeatured:  true,
			Discount:    0.025,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420230/smart_watch_mi_band_4_2_mjutcx.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420230/smart_watch_mi_band_4_n3vcip.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"", "", 750000, 50, ""},
			},
		},
		{
			Category:    "Gadget & Electronics",
			Subcategory: "Wearable Devices",
			Name:        "Samsung Galaxy Watch 4 Classic 42mm",
			Description: "Samsung Watch 4 hadir dengan display Sapphire Crystal, GPS, sleep tracker dan body composition. Smartwatch yang menawarkan berbagai fitur kesehatan dan kebugaran, serta integrasi yang mulus dengan perangkat Galaxy lainnya. Smartwatch ini dilengkapi dengan sensor BioActive yang mampu memantau detak jantung, tekanan darah, kadar oksigen dalam darah, dan kualitas tidur. Selain itu, Galaxy Watch juga mendukung fitur-fitur lain seperti menerima panggilan dan pesan, mengontrol musik, dan memberikan notifikasi.",
			IsFeatured:  false,
			Discount:    0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420675/samsung_watch_03_bmlayk.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420675/samsung_watch_03_bmlayk.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"white", "", 875000, 10, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420675/samsung_watch_04_uh1fjs.webp"},
				{"black", "", 875000, 20, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420675/samsung_watch_02_szbzqg.webp"},
			},
		},
		{
			Category:    "Gadget & Electronics",
			Subcategory: "Wearable Devices",
			Name:        "HUAWEI WATCH FIT Special Edition Smartwatch",
			Description: "HUAWEI WATCH FIT Special Edition Smartwatch | 1.64 HD AMOLED | 24/7 Active Health Management | Built-in GPS | Fast Charging. Notifikasi panggilan Bluetooth & balas pesan cepat Kompatibel dengan luas, bisa digunakan bersama semua OS Tersedia dalam 3 varian warna: Nebula Pink, Forest Green, Starry Black.",
			IsFeatured:  false,
			Discount:    0.02,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421186/huawei_smartwatch_04_r8ftp5.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421185/huawei_smartwatch_02_ihjja7.webp",
			},
			Variants: []struct {
				Color string
				Size  string
				Price float64
				Stock int
				Image string
			}{
				{"blue", "", 545000, 10, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421187/huawei_smartwatch_05_qbvhc7.webp"},
				{"black", "", 545000, 20, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421185/huawei_smartwatch_03_wswy7h.webp"},
				{"pink", "", 545000, 30, "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421185/huawei_smartwatch_01_iwdoic.webp"},
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
