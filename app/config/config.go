package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBPort     string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func LoadConfig() Config {
	// Load .env file if exists
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Return the application configuration
	return Config{
		Port:       getEnv("PORT", "8080"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "ecommerce_db"),
		JWTSecret:  getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
