package main

import (
	"log"
	"os"

	global "github.com/fiqrioemry/microservice-ecommerce/server/pkg/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/seeders"

	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/handlers"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/repositories"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/routes"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	config.InitDatabase()
	global.InitRedis()
	global.InitMailer()

	db := config.DB

	// router initiate
	r := gin.Default()
	r.Use(middleware.Logger(), middleware.Recovery(), middleware.CORS(), middleware.RateLimiter(5, 10), middleware.LimitFileSize(5<<20))

	// service & handler dependency
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	profileService := services.NewProfileService(userRepo)
	profileHandler := handlers.NewProfileHandler(profileService)

	addressRepo := repositories.NewAddressRepository(db)
	addressService := services.NewAddressService(addressRepo)
	addressHandler := handlers.NewAddressHandler(addressService)

	// Routing dependency injection
	routes.AuthRoutes(r, authHandler)
	routes.AdminRoutes(r, authHandler)
	routes.UserRoutes(r, profileHandler, addressHandler)

	seeders.SeedUsers(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Fatal(r.Run(":" + port))
}
