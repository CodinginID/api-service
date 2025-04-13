package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthHandler struct {
	service AuthService
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	repo := NewAuthRepository(db)
	service := NewAuthService(repo)
	return &AuthHandler{service}
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.service.Register(req.Username, req.Email, req.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Registration successful"})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}

func (h *AuthHandler) GetUserDetailByUsername(c *fiber.Ctx) error {
	username := c.Query("username")
	user, err := h.service.GetUserDetailByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func (h *AuthHandler) GetUsersAfterDate(c *fiber.Ctx) error {
	date := c.Query("date")
	if date == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Date query parameter is required"})
	}

	users, err := h.service.GetUsersAfterDate(date)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}
