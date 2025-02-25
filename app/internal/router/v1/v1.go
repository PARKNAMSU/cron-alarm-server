package v1

import (
	"github.com/gofiber/fiber/v2"
)

func Router() func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Route("/sample", SampleV1Router())
		router.Route("/user", UserRouter())
	}
}
