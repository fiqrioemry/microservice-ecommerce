package main

import (
	"log"
	"net"
	"os"

	global "github.com/fiqrioemry/microservice-ecommerce/server/pkg/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/seeders"

	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/handlers"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/repositories"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/routes"
	"github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/services"

	userpb "github.com/fiqrioemry/microservice-ecommerce/server/proto/user"
	usergrpc "github.com/fiqrioemry/microservice-ecommerce/server/user-service/internal/grpc"
)

// TODO : Ganti session-based authentication ke JWT dan tambahkan oAuth2.0

func main() {
	utils.LoadEnv()
	config.InitDatabase()
	global.InitRedis()
	global.InitMailer()
	global.InitCloudinary()

	db := config.DB

	// auth
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// profile
	profileService := services.NewProfileService(userRepo)
	profileHandler := handlers.NewProfileHandler(profileService)

	// location
	locationRepo := repositories.NewLocationRepository(db)
	locationService := services.NewLocationService(locationRepo)
	locationHandler := handlers.NewLocationHandler(locationService)

	// address
	addressRepo := repositories.NewAddressRepository(db)
	addressService := services.NewAddressService(addressRepo, locationRepo)
	addressHandler := handlers.NewAddressHandler(addressService)

	go func() {
		lis, err := net.Listen("tcp", ":50052")
		if err != nil {
			log.Fatalf("failed to listen on port 50052: %v", err)
		}

		grpcServer := grpc.NewServer()
		userpb.RegisterUserServiceServer(grpcServer, usergrpc.NewUserGRPCServer(addressService))

		log.Println("User gRPC server running on port 50052...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC server: %v", err)
		}
	}()

	r := gin.Default()
	// TODO : Tambahkan pengaturan untuk connection pooling di database
	r.Use(middleware.Logger(), middleware.Recovery(), middleware.CORS(), middleware.RateLimiter(5, 10), middleware.LimitFileSize(5<<20))

	// Routing dependency injection
	routes.AuthRoutes(r, authHandler)
	routes.LocationRoutes(r, locationHandler)
	routes.UserRoutes(r, profileHandler, addressHandler)

	seeders.SeedUsers(db)
	seeders.SeedLocations(db)

	port := os.Getenv("PORT")
	log.Println("user service running on port:", port)
	log.Fatal(r.Run(":" + port))
}
