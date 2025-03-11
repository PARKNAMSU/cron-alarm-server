package user_controller

import (
	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/usecase/user_usecase"
)

type UserController struct {
	usecase user_usecase.UserUsecaseImpl
}

var (
	controller *UserController
)

func NewController(usecase user_usecase.UserUsecaseImpl) *UserController {
	if controller == nil {
		controller = &UserController{
			usecase: usecase,
		}
	}
	return controller
}

func (c *UserController) SignUp(ctx *fiber.Ctx) error {
	body := ctx.Context().Value("body").(map[string]any)
	input := user_usecase.SignUpInput{
		Email:    body["email"].(string),
		Password: body["password"].(string),
		IpAddr:   ctx.IP(),
	}
	output, err := c.usecase.SignUp(input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	return ctx.JSON(fiber.Map{
		"data": output,
	})
}

func (c *UserController) SignIn(ctx *fiber.Ctx) error {
	body := ctx.Context().Value("body").(map[string]any)
	input := user_usecase.SignInInput{
		Email:    body["email"].(string),
		Password: body["password"].(string),
		IpAddr:   ctx.IP(),
	}

	output, err := c.usecase.SignIn(input)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.JSON(fiber.Map{"data": output})
}
