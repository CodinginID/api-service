package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/CodinginID/api-service/app/internal/auth"
	"github.com/CodinginID/api-service/app/pkg/db"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, continuing...")
	}

	// Init DB
	database, err := db.InitPostgres()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Init Fiber App
	app := fiber.New()

	// Init Route Groups
	api := app.Group("/api/v1")
	auth.RegisterAuthRoutes(api.Group("/auth"), database)

	// Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
