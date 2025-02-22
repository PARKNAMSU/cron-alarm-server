package sample_controller

import "github.com/gofiber/fiber/v2"

func SampleController(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "Hello, World!"})
}
