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
	file, err := os.Open("province.csv")
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
		id, _ := strconv.ParseUint(record[0], 10, 64)
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
	file, err := os.Open("city.csv")
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
		id, _ := strconv.ParseUint(record[0], 10, 64)
		provinceID, _ := strconv.ParseUint(record[1], 10, 64)
		city := models.City{
			ID:         uint(id),
			ProvinceID: uint(provinceID),
			Name:       record[2],
		}
		cities = append(cities, city)
	}

	if err := db.Create(&cities).Error; err != nil {
		log.Fatal("Failed to seed cities:", err)
	}
	log.Println("✅ City seeding completed")
}

func seedDistricts(db *gorm.DB) {
	file, err := os.Open("district.csv")
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
		id, _ := strconv.ParseUint(record[0], 10, 64)
		cityID, _ := strconv.ParseUint(record[1], 10, 64)
		district := models.District{
			ID:     uint(id),
			CityID: uint(cityID),
			Name:   record[2],
		}
		districts = append(districts, district)
	}

	if err := db.Create(&districts).Error; err != nil {
		log.Fatal("Failed to seed districts:", err)
	}
	log.Println("✅ District seeding completed")
}

func seedSubdistricts(db *gorm.DB) {
	file, err := os.Open("subdistrict.csv")
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
		id, _ := strconv.ParseUint(record[0], 10, 64)
		districtID, _ := strconv.ParseUint(record[1], 10, 64)
		subdistrict := models.Subdistrict{
			ID:         uint(id),
			DistrictID: uint(districtID),
			Name:       record[2],
		}
		subdistricts = append(subdistricts, subdistrict)
	}

	if err := db.Create(&subdistricts).Error; err != nil {
		log.Fatal("Failed to seed subdistricts:", err)
	}
	log.Println("✅ Subdistrict seeding completed")
}

func seedPostalCodes(db *gorm.DB) {
	file, err := os.Open("postal_code.csv")
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
		id, _ := strconv.ParseUint(record[0], 10, 64)
		subdistrictID, _ := strconv.ParseUint(record[1], 10, 64)
		postalCode := models.PostalCode{
			ID:            uint(id),
			SubdistrictID: uint(subdistrictID),
			PostalCode:    record[5],
		}
		postalCodes = append(postalCodes, postalCode)
	}

	if err := db.Create(&postalCodes).Error; err != nil {
		log.Fatal("Failed to seed postal codes:", err)
	}
	log.Println("✅ Postal code seeding completed")
}
