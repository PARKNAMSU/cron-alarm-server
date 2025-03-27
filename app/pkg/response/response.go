package response

import "github.com/gofiber/fiber/v2"

func ServerError(c *fiber.Ctx, errData any) error {
	return c.Status(fiber.StatusInternalServerError).JSON(errData)
}

func CustomError(c *fiber.Ctx, errData any, status int) error {
	return c.Status(status).JSON(errData)
}

func SendJson(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(data)
}

func SendText(c *fiber.Ctx, str string) error {
	return c.SendString(str)
}
