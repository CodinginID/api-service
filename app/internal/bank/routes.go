package bank

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterBankRoutes(router fiber.Router, db *gorm.DB) {
	handler := NewBankAccountHandler(db)

	router.Get("/balance", handler.GetBalance)
	router.Post("/deposit", handler.Deposit)
	router.Post("/withdraw", handler.Withdraw)
}
