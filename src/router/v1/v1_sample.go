package v1

import (
	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/src/controller/sample_controller"
)

func SampleV1Router() func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Get("/", sample_controller.SampleController)
	}
}
