package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

var SecretKey = "secret" // sebaiknya dari env var

func JWTProtected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(SecretKey),
		ContextKey:   "user", // disimpan di c.Locals("user")
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

// Untuk ambil user_id dari token
func GetUserIDFromToken(c *fiber.Ctx) (uint, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))
	return userID, nil
}
