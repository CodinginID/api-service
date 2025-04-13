package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterProductRoutes(r fiber.Router, db *gorm.DB) {
	h := NewProductHandler(db)

	r.Post("/", h.Create)
	r.Get("/", h.List)
	r.Get("/:id", h.Detail)
}
