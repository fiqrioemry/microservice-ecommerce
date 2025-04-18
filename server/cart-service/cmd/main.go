package main

import (
	"log"
	"os"

	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/handlers"
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/repositories"
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/routes"
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/services"
	global "github.com/fiqrioemry/microservice-ecommerce/server/pkg/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	config.InitDatabase()
	global.InitRedis()

	db := config.DB

	repo := repositories.NewCartRepository(db)
	service := services.NewCartService(repo)
	handler := handlers.NewCartHandler(service)

	r := gin.Default()
	r.Use(middleware.Logger(), middleware.Recovery(), middleware.CORS(), middleware.RateLimiter(5, 10))

	routes.CartRoutes(r, handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5003"
	}
	log.Fatal(r.Run(":" + port))
}
