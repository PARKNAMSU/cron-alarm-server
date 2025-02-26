package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/database"
)

var (
	middleware *Middleware
	apiKey     = os.Getenv("X_API_KEY")
)

type Middleware struct {
	slaveDB  *database.CustomDB
	masterDB *database.CustomDB
}

func NewMiddleware(slaveDB *database.CustomDB, masterDB *database.CustomDB) *Middleware {
	if middleware == nil {
		middleware =
			&Middleware{
				slaveDB:  slaveDB,
				masterDB: masterDB,
			}
	}
	return middleware
}

func (m *Middleware) APIValidation(c *fiber.Ctx) error {
	headerApiKey := c.Get("x-api-key")
	if apiKey != headerApiKey {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}
	return c.Next()
}

func (m *Middleware) UserValidation(c *fiber.Ctx) error {
	// todo: 유저 검증 로직 추가
	return c.Next()
}
