package cart

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterCartRoutes(r fiber.Router, db *gorm.DB) {
	h := NewCartHandler(db)

	r.Post("/", h.Add)
	r.Get("/", h.List)
	r.Delete("/", h.Clear)
}
