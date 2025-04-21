package seeders

import (
	"encoding/json"
	"log"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/models"
)

func SeedUsers(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), 10)

	adminUser := models.User{
		ID:       uuid.New(),
		Email:    "admin@example.com",
		Password: string(password),
		Role:     "admin",
		Profile: models.Profile{
			Fullname: "Admin User",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=Admin",
		},
	}

	customerUser := models.User{
		ID:       uuid.New(),
		Email:    "customer@example.com",
		Password: string(password),
		Role:     "customer",
		Profile: models.Profile{
			Fullname: "Customer User",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=Customer",
		},
	}

	if err := db.Create(&adminUser).Error; err != nil {
		log.Println("Failed to seed admin:", err)
	}
	if err := db.Create(&customerUser).Error; err != nil {
		log.Println("Failed to seed customer:", err)
	}

	log.Println("✅ User seeding completed")
}

type ProvinceSeeder struct {
	ID   string `json:"province_id"`
	Name string `json:"province"`
}

type CitySeeder struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

func SeedProvincesAndCities(db *gorm.DB) {
	provinceFile, err := os.ReadFile("provinces.json")
	if err != nil {
		log.Println("Error reading province.json:", err)
		return
	}
	cityFile, err := os.ReadFile("cities.json")
	if err != nil {
		log.Println("Error reading cities.json:", err)
		return
	}

	var provinceWrapper struct {
		RajaOngkir struct {
			Results []ProvinceSeeder `json:"results"`
		} `json:"rajaongkir"`
	}

	var cityWrapper struct {
		RajaOngkir struct {
			Results []CitySeeder `json:"results"`
		} `json:"rajaongkir"`
	}

	if err := json.Unmarshal(provinceFile, &provinceWrapper); err != nil {
		log.Println("Failed to unmarshal province.json:", err)
		return
	}
	if err := json.Unmarshal(cityFile, &cityWrapper); err != nil {
		log.Println("Failed to unmarshal cities.json:", err)
		return
	}

	provinceIDMap := make(map[string]uint)
	for _, p := range provinceWrapper.RajaOngkir.Results {
		province := models.Province{
			Name: p.Name,
		}
		db.Create(&province)
		provinceIDMap[p.ID] = province.ID
	}

	for _, c := range cityWrapper.RajaOngkir.Results {
		city := models.City{
			ProvinceID: provinceIDMap[c.ProvinceID],
			Name:       c.CityName,
			Type:       c.Type,
			PostalCode: c.PostalCode,
		}
		db.Create(&city)
	}

	log.Println("✅ Province and City seeding completed.")
}
