package main

import (
	"log"
	"os"

	global "github.com/fiqrioemry/microservice-ecommerce/server/pkg/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/gin-gonic/gin"

	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/handlers"
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/repositories"
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/routes"
	"github.com/fiqrioemry/microservice-ecommerce/server/order-service/internal/services"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/grpc"
)

func main() {
	utils.LoadEnv()
	config.InitDatabase()
	global.InitRedis()
	config.InitMidtrans()

	db := config.DB

	orderRepo := repositories.NewOrderRepository(db)

	userGrcpAdd := os.Getenv("USER_GRPC_ADDR")
	cartGrcpAdd := os.Getenv("CART_GRPC_ADDR")
	productGrcpAdd := os.Getenv("PRODUCT_GRPC_ADDR")

	productGRPC, err := grpc.NewProductGRPCClient(productGrcpAdd)
	if err != nil {
		log.Fatal("failed to connect to product-service:", err)
	}

	cartGRPC, err := grpc.NewCartGRPCClient(cartGrcpAdd)
	if err != nil {
		log.Fatal("failed to connect to cart-service:", err)
	}

	userGRPC, err := grpc.NewUserGRPCClient(userGrcpAdd)
	if err != nil {
		log.Fatal("failed to connect to user-service:", err)
	}

	orderService := services.NewOrderService(orderRepo, cartGRPC, userGRPC, productGRPC)

	orderHandler := handlers.NewOrderHandler(orderService)

	r := gin.Default()
	r.Use(
		middleware.Logger(),
		middleware.Recovery(),
		middleware.CORS(),
		middleware.RateLimiter(5, 10),
		middleware.LimitFileSize(5<<20),
	)

	routes.OrderRoutes(r, orderHandler)

	port := os.Getenv("PORT")
	log.Println("order service running on port:", port)
	log.Fatal(r.Run(":" + port))
}
