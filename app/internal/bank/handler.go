package bank

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Request struct {
	Amount int `json:"amount"`
}

type BankAccountHandler struct {
	service BankAccountService
}

func NewBankAccountHandler(db *gorm.DB) *BankAccountHandler {
	repo := NewBankAccountRepository(db)
	service := NewBankAccountService(repo)
	return &BankAccountHandler{service}
}

func (h *BankAccountHandler) GetBalance(c *fiber.Ctx) error {
	var req Request
	amount := h.service.GetBalance()
	req.Amount = amount
	return c.JSON(fiber.Map{"balance": req.Amount})
}
func (h *BankAccountHandler) Deposit(c *fiber.Ctx) error {
	var req Request
	if err := c.BodyParser(&req); err != nil || req.Amount <= 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid request"})
	}
	amount := req.Amount
	h.service.Deposit(amount)
	return c.JSON(fiber.Map{"message": "Deposit successful"})
}
func (h *BankAccountHandler) Withdraw(c *fiber.Ctx) error {
	var req Request
	if err := c.BodyParser(&req); err != nil || req.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	amount := req.Amount
	if err := h.service.Withdraw(amount); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	c.JSON(fiber.Map{"message": "Withdraw successful"})
	return nil
}
