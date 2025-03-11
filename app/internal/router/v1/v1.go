package v1

import (
	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/di"
)

var (
	middleware     = di.InitMiddleware()
	userController = di.InitUserController()
)

func Router() func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Route("/sample", SampleV1Router())
		router.Route("/user", UserRouter())
		router.Route("/open", OpenRouter())
	}
}
