package seeders

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/models"
	"gorm.io/gorm"
)

func SeedLocations(db *gorm.DB) {
	seedProvinces(db)
	seedCities(db)
	seedDistricts(db)
	seedSubdistricts(db)
	seedPostalCodes(db)
}

// --- Province ---
func seedProvinces(db *gorm.DB) {
	file, err := os.Open("internal/seeders/province.csv")
	if err != nil {
		log.Fatal("Failed to open province.csv:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var provinces []models.Province
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, _ := strconv.ParseUint(record[0], 10, 64) // id
		provinces = append(provinces, models.Province{
			ID:   uint(id),
			Name: record[1], // name
		})
	}

	if err := db.Create(&provinces).Error; err != nil {
		log.Fatal("Failed to seed provinces:", err)
	}
	log.Println("✅ Province seeding completed")
}

// --- City ---
func seedCities(db *gorm.DB) {
	file, err := os.Open("internal/seeders/city.csv")
	if err != nil {
		log.Fatal("Failed to open city.csv:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read()

	var cities []models.City
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, _ := strconv.ParseUint(record[0], 10, 64)         // id
		provinceID, _ := strconv.ParseUint(record[1], 10, 64) // province_id
		name := record[2]                                     // city name

		cities = append(cities, models.City{
			ID:         uint(id),
			ProvinceID: uint(provinceID),
			Name:       name,
		})
	}

	if err := db.Create(&cities).Error; err != nil {
		log.Fatal("Failed to seed cities:", err)
	}
	log.Println("✅ City seeding completed")
}

// --- District ---
func seedDistricts(db *gorm.DB) {
	file, err := os.Open("internal/seeders/district.csv")
	if err != nil {
		log.Fatal("Failed to open district.csv:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read()

	var districts []models.District
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, _ := strconv.ParseUint(record[0], 10, 64)     // id
		cityID, _ := strconv.ParseUint(record[1], 10, 64) // city_id
		name := record[2]                                 // district name

		districts = append(districts, models.District{
			ID:     uint(id),
			CityID: uint(cityID),
			Name:   name,
		})
	}

	if err := db.Create(&districts).Error; err != nil {
		log.Fatal("Failed to seed districts:", err)
	}
	log.Println("✅ District seeding completed")
}

// --- Subdistrict ---
func seedSubdistricts(db *gorm.DB) {
	file, err := os.Open("internal/seeders/subdistrict.csv")
	if err != nil {
		log.Fatal("Failed to open subdistrict.csv:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read()

	var subdistricts []models.Subdistrict
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, _ := strconv.ParseUint(record[0], 10, 64)         // id
		districtID, _ := strconv.ParseUint(record[1], 10, 64) // district_id
		name := record[2]                                     // subdistrict name

		subdistricts = append(subdistricts, models.Subdistrict{
			ID:         uint(id),
			DistrictID: uint(districtID),
			Name:       name,
		})
	}

	if err := db.Create(&subdistricts).Error; err != nil {
		log.Fatal("Failed to seed subdistricts:", err)
	}
	log.Println("✅ Subdistrict seeding completed")
}

// --- Postal Code ---
func seedPostalCodes(db *gorm.DB) {
	file, err := os.Open("internal/seeders/postal_code.csv")
	if err != nil {
		log.Fatal("Failed to open postal_code.csv:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read()

	var postalCodes []models.PostalCode
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, _ := strconv.ParseUint(record[0], 10, 64)            // id
		subdistrictID, _ := strconv.ParseUint(record[1], 10, 64) // subdistrict_id
		code := record[5]                                        // postal_code

		postalCodes = append(postalCodes, models.PostalCode{
			ID:            uint(id),
			SubdistrictID: uint(subdistrictID),
			PostalCode:    code,
		})
	}

	if err := db.Create(&postalCodes).Error; err != nil {
		log.Fatal("Failed to seed postal codes:", err)
	}
	log.Println("✅ Postal Code seeding completed")
}
