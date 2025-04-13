package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(router fiber.Router, db *gorm.DB) {
	handler := NewAuthHandler(db)

	router.Post("/register", handler.Register)
	router.Post("/login", handler.Login)
}
