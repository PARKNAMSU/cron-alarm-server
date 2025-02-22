package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

var (
	API_KEY = os.Getenv("X_API_KEY")
)

func APIValidation(c *fiber.Ctx) error {
	apiKey := c.Get("x-api-key")
	if apiKey != API_KEY {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}
	return c.Next()
}
