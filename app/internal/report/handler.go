package report

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ReportHandler struct {
	service ReportService
}

func NewReportHandler(db *gorm.DB) *ReportHandler {
	repo := NewReportRepository(db)
	service := NewReportService(repo)
	return &ReportHandler{service}
}

type TopCustomer struct {
	CustomerID int     `json:"customer_id"`
	TotalSpent float64 `json:"total_spent"`
}

func (h *ReportHandler) GetTopCustomers(c *fiber.Ctx) error {
	topCustomers, err := h.service.GetTopCustomers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if len(topCustomers) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "No top customers found"})
	}
	// Return the top customers as JSON
	return c.JSON(topCustomers)
}
