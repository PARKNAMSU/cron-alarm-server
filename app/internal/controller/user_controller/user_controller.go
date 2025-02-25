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

func (c *UserController) SignIn(ctx *fiber.Ctx) error {
	// todo : 회원가입 구현
	return ctx.JSON(fiber.Map{"message": "success"})
}
