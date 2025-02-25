package v1

import (
	"github.com/gofiber/fiber/v2"
)

func APIV1Router() func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Route("/sample", SampleV1Router())
	}
}
