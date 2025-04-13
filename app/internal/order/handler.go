package order

import (
	"github.com/CodinginID/api-service/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	service OrderService
}

func NewOrderHandler(service OrderService) *OrderHandler {
	return &OrderHandler{service}
}

func (h *OrderHandler) Checkout(c *fiber.Ctx) error {
	userID, err := middleware.GetUserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	order, err := h.service.Checkout(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

func (h *OrderHandler) GetOrderHistory(c *fiber.Ctx) error {
	userID, err := middleware.GetUserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	orders, err := h.service.GetOrderHistory(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(orders)
}
