package v1

import (
	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/di"
)

func UserRouter() func(router fiber.Router) {
	controller := di.InitUserController()
	return func(router fiber.Router) {
		router.Post("/signUp", controller.SignUp)
	}
}
