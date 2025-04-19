package main

import (
	"log"
	"net"
	"os"

	global "github.com/fiqrioemry/microservice-ecommerce/server/pkg/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/config"
	grpcServer "github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/grpc"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/repositories"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/routes"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/seeders"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/services"
	productpb "github.com/fiqrioemry/microservice-ecommerce/server/proto/product"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	config.InitDatabase()
	global.InitRedis()
	global.InitCloudinary()

	db := config.DB

	// variant
	variantRepo := repositories.NewVariantRepository(db)
	variantService := services.NewVariantService(variantRepo)
	variantHandler := handlers.NewVariantHandler(variantService)

	// product
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Category

	categoryRepo := repositories.NewUnifiedCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	attributeRepo := repositories.NewAttributeRepository(db)
	attributeService := services.NewAttributeService(attributeRepo)
	attributeHandler := handlers.NewAttributeHandler(attributeService)

	r := gin.Default()
	r.Use(middleware.Logger(), middleware.Recovery(), middleware.CORS(), middleware.RateLimiter(5, 10), middleware.LimitFileSize(5<<20))

	routes.VariantRoutes(r, variantHandler)
	routes.ProductRoutes(r, productHandler)
	routes.CategoryRoutes(r, categoryHandler)
	routes.AttributeRoutes(r, attributeHandler)

	seeders.SeedInitialData(db)

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		productpb.RegisterProductServiceServer(s, grpcServer.NewProductGRPCServer(productRepo))
		log.Println("gRPC server running on port 50051")

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// Jalankan server
	port := os.Getenv("PORT")
	log.Println("product service running on port:", port)
	log.Fatal(r.Run(":" + port))
}
