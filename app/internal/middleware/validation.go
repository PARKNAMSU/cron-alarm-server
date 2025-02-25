package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

var (
	apiKey = os.Getenv("X_API_KEY")
)

func APIValidation(c *fiber.Ctx) error {
	apiKey := c.Get("x-api-key")
	if apiKey != apiKey {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}
	return c.Next()
}

func UserValidation(c *fiber.Ctx) error {
	// todo: 유저 검증 로직 추가
	return c.Next()
}
