package v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/di"
)

var validate = validator.New()

var (
	middlewareInject = di.InitMiddleware()
	userController   = di.InitUserController()
)

func Router() func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Route("/sample", SampleV1Router())
		router.Route("/user", UserRouter())
		router.Route("/open", OpenRouter())
	}
}
