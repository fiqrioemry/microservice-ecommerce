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

		id, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid province ID: %v", err)
		}

		province := models.Province{
			ID:   uint(id),
			Name: record[1],
		}
		provinces = append(provinces, province)
	}

	if err := db.Create(&provinces).Error; err != nil {
		log.Fatal("Failed to seed provinces:", err)
	}
	log.Println("✅ Province seeding completed")
}

func seedCities(db *gorm.DB) {
	file, err := os.Open("internal/seeders/city.csv")
	if err != nil {
		log.Fatal("Failed to open city.csv:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var cities []models.City

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid city ID: %v", err)
		}

		provinceID, err := strconv.ParseUint(record[2], 10, 64)
		if err != nil {
			log.Fatalf("Invalid province ID for city: %v", err)
		}

		city := models.City{
			ID:         uint(id),
			ProvinceID: uint(provinceID),
			Name:       record[1],
		}
		cities = append(cities, city)
	}

	if err := db.Create(&cities).Error; err != nil {
		log.Fatal("Failed to seed cities:", err)
	}
	log.Println("✅ City seeding completed")
}

func seedDistricts(db *gorm.DB) {
	file, err := os.Open("internal/seeders/district.csv")
	if err != nil {
		log.Fatal("Failed to open district.csv:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var districts []models.District

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid district ID: %v", err)
		}

		cityID, err := strconv.ParseUint(record[2], 10, 64)
		if err != nil {
			log.Fatalf("Invalid city ID for district: %v", err)
		}

		district := models.District{
			ID:     uint(id),
			CityID: uint(cityID),
			Name:   record[1],
		}
		districts = append(districts, district)
	}

	if err := db.Create(&districts).Error; err != nil {
		log.Fatal("Failed to seed districts:", err)
	}
	log.Println("✅ District seeding completed")
}

func seedSubdistricts(db *gorm.DB) {
	file, err := os.Open("internal/seeders/subdistrict.csv")
	if err != nil {
		log.Fatal("Failed to open subdistrict.csv:", err)
	}
	defer file.Close()

	var count int64
	db.Model(&models.Subdistrict{}).Count(&count)
	if count > 0 {
		log.Println("✅ Subdistricts already seeded, skipping...")
		return
	}

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var subdistricts []models.Subdistrict

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid subdistrict ID: %v", err)
		}

		districtID, err := strconv.ParseUint(record[2], 10, 64)
		if err != nil {
			log.Fatalf("Invalid district ID for subdistrict: %v", err)
		}

		subdistrict := models.Subdistrict{
			ID:         uint(id),
			DistrictID: uint(districtID),
			Name:       record[1],
		}
		subdistricts = append(subdistricts, subdistrict)
	}

	if err := db.CreateInBatches(&subdistricts, 500).Error; err != nil {
		log.Fatal("Failed to seed subdistricts:", err)
	}
	log.Println("✅ Subdistrict seeding completed")
}

func seedPostalCodes(db *gorm.DB) {
	file, err := os.Open("internal/seeders/postal_code.csv")
	if err != nil {
		log.Fatal("Failed to open postal_code.csv:", err)
	}
	defer file.Close()

	var count int64
	db.Model(&models.PostalCode{}).Count(&count)
	if count > 0 {
		log.Println("✅ Postal codes already seeded, skipping...")
		return
	}

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var postalCodes []models.PostalCode

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid postal ID: %v", err)
		}
		provinceID, err := strconv.ParseUint(record[4], 10, 64)
		if err != nil {
			log.Fatalf("Invalid province ID for postal code: %v", err)
		}
		cityID, err := strconv.ParseUint(record[3], 10, 64)
		if err != nil {
			log.Fatalf("Invalid city ID for postal code: %v", err)
		}
		districtID, err := strconv.ParseUint(record[2], 10, 64)
		if err != nil {
			log.Fatalf("Invalid district ID for postal code: %v", err)
		}
		subdistrictID, err := strconv.ParseUint(record[1], 10, 64)
		if err != nil {
			log.Fatalf("Invalid subdistrict ID for postal code: %v", err)
		}

		postalCode := models.PostalCode{
			ID:            uint(id),
			SubdistrictID: uint(subdistrictID),
			DistrictID:    uint(districtID),
			CityID:        uint(cityID),
			ProvinceID:    uint(provinceID),
			PostalCode:    record[5],
		}
		postalCodes = append(postalCodes, postalCode)
	}

	if err := db.CreateInBatches(&postalCodes, 500).Error; err != nil {
		log.Fatal("Failed to seed postal codes:", err)
	}
	log.Println("✅ Postal code seeding completed")
}
