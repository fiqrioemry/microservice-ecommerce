package main

import (
	"log"
	"net"
	"os"

	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/config"
	cartServer "github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/grpc"
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/handlers"
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/repositories"
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/routes"
	"github.com/fiqrioemry/microservice-ecommerce/server/cart-service/internal/services"
	cartClient "github.com/fiqrioemry/microservice-ecommerce/server/pkg/grpc"
	cartpb "github.com/fiqrioemry/microservice-ecommerce/server/proto/cart"
	"google.golang.org/grpc"

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

	productGrcpAdd := os.Getenv("PRODUCT_GRPC_ADDR")

	//initiate cart client to connect to product service
	productGRPCClient, err := cartClient.NewProductGRPCClient(productGrcpAdd)
	if err != nil {
		log.Fatalf("Failed to connect to product-service gRPC: %v", err)
	}

	repo := repositories.NewCartRepository(db)
	service := services.NewCartService(repo)
	handler := handlers.NewCartHandler(service, *productGRPCClient)

	go func() {
		lis, err := net.Listen("tcp", ":50053")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		// expose grpc service to other service
		cartpb.RegisterCartServiceServer(s, cartServer.NewCartGRPCServer(service))
		log.Println("gRPC server running on port 50051")

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	r := gin.Default()
	r.Use(middleware.Logger(), middleware.Recovery(), middleware.CORS(), middleware.RateLimiter(5, 10))

	routes.CartRoutes(r, handler)

	port := os.Getenv("PORT")
	log.Println("Cart service running on port:", port)
	log.Fatal(r.Run(":" + port))
}
