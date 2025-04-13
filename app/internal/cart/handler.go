package cart

import (
	"github.com/CodinginID/api-service/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CartHandler struct {
	service CartService
}

func NewCartHandler(db *gorm.DB) *CartHandler {
	repo := NewCartRepository(db)
	service := NewCartService(repo)
	return &CartHandler{service}
}

type AddCartRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func (h *CartHandler) Add(c *fiber.Ctx) error {
	userID, err := middleware.GetUserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var req AddCartRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err = h.service.Add(userID, req.ProductID, req.Quantity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Item added to cart"})
}

func (h *CartHandler) List(c *fiber.Ctx) error {
	userID, err := middleware.GetUserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	items, err := h.service.List(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(items)
}

func (h *CartHandler) Clear(c *fiber.Ctx) error {
	userID, err := middleware.GetUserIDFromToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	err = h.service.Clear(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Cart cleared"})
}
