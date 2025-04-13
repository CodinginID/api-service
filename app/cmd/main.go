package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger" // Correct import for Fiber's logger

	"github.com/CodinginID/api-service/config"
	"github.com/CodinginID/api-service/internal/auth"
	"github.com/CodinginID/api-service/internal/cart"
	"github.com/CodinginID/api-service/internal/middleware"
	"github.com/CodinginID/api-service/internal/order"
	"github.com/CodinginID/api-service/internal/product"
	"github.com/CodinginID/api-service/pkg/db"
)

func main() {

	cfg := config.LoadConfig()
	// Init DB
	database, err := db.InitPostgres(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	// Check if the database connection was successful
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	database.AutoMigrate(&auth.User{})

	// Init Fiber App
	app := fiber.New()

	// Enable request logging with Fiber's logger middleware
	app.Use(logger.New())

	// Set up CORS if needed
	app.Use(cors.New())

	// Init Route Groups
	api := app.Group("/api/v1")
	auth.RegisterAuthRoutes(api.Group("/auth"), database)
	
	productGroup := api.Group("/products")
	productGroup.Use(middleware.JWTProtected())
	product.RegisterProductRoutes(productGroup, database)

	cartGroup := api.Group("/cart")
	cartGroup.Use(middleware.JWTProtected())
	cart.RegisterCartRoutes(cartGroup, database)

	orderGroup := api.Group("/order")
	orderGroup.Use(middleware.JWTProtected())
	order.RegisterOrderRoutes(orderGroup, database)

	// Start Server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
