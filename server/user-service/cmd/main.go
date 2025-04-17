package main

import (
	"log"
	"os"

	global "github.com/fiqrioemry/microservice-ecommerce/server/pkg/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"

	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/config"

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
	router := gin.Default()
	router.Use(middleware.Logger(), middleware.Recovery(), middleware.CORS(), middleware.RateLimiter(5, 10), middleware.LimitFileSize(5<<20))

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
	routes.AuthRoutes(router, authHandler)
	routes.AdminRoutes(router, authHandler)
	routes.UserRoutes(router, profileHandler, addressHandler)

	// Jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Fatal(router.Run(":" + port))
}
