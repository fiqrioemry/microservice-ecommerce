package main

import (
	"log"
	"os"

	global "github.com/fiqrioemry/microservice-ecommerce/server/pkg/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/middleware"
	"github.com/fiqrioemry/microservice-ecommerce/server/pkg/utils"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/config"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/handlers"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/repositories"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/routes"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/seeders"
	"github.com/fiqrioemry/microservice-ecommerce/server/product-service/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	config.InitDatabase()
	global.InitRedis()
	global.InitCloudinary()

	db := config.DB

	//  size
	sizeRepo := repositories.NewSizeRepository(db)
	sizeService := services.NewSizeService(sizeRepo)
	sizeHandler := handlers.NewSizeHandler(sizeService)

	// color
	colorRepo := repositories.NewColorRepository(db)
	colorService := services.NewColorService(colorRepo)
	colorHandler := handlers.NewColorHandler(colorService)

	// product
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Category
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// subcategory
	subcategoryRepo := repositories.NewSubcategoryRepository(db)
	subcategoryService := services.NewSubcategoryService(subcategoryRepo)
	subcategoryHandler := handlers.NewSubcategoryHandler(subcategoryService)

	// variant
	variantRepo := repositories.NewVariantRepository(db)
	variantService := services.NewVariantService(variantRepo)
	variantHandler := handlers.NewVariantHandler(variantService)

	// attribute
	attrRepo := repositories.NewAttributeRepository(db)
	valRepo := repositories.NewAttributeValueRepository(db)
	attrService := services.NewAttributeService(attrRepo, valRepo)
	attrHandler := handlers.NewAttributeHandler(attrService)

	r := gin.Default()
	r.Use(middleware.Logger(), middleware.Recovery(), middleware.CORS(), middleware.RateLimiter(5, 10), middleware.LimitFileSize(5<<20))

	routes.SizeRoutes(r, sizeHandler)
	routes.ColorRoutes(r, colorHandler)
	routes.AttributeRoutes(r, attrHandler)
	routes.ProductRoutes(r, productHandler)
	routes.VariantRoutes(r, variantHandler)
	routes.CategoryRoutes(r, categoryHandler)
	routes.SubcategoryRoutes(r, subcategoryHandler)

	seeders.SeedProductData(db)
	seeders.SeedProductOptions(db)

	seeders.SeedVariantsAndAttributes(db)

	// Jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Fatal(r.Run(":" + port))
}
