package report

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterReportRoutes(router fiber.Router, db *gorm.DB) {
	handler := NewReportHandler(db)

	router.Get("/top-customers", handler.GetTopCustomers)
}
