package user_controller

import (
	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/global_type"
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

func (c *UserController) Authorization(ctx *fiber.Ctx) error {
	userData, isExist := ctx.Context().Value("userData").(global_type.UserTokenData)
	body := ctx.Context().Value("body").(map[string]any)

	code := body["code"].(string)

	if !isExist {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid user token",
			"code":    "INVALID-USER-TOKEN",
		})
	}

	output, err := c.usecase.Authorization(user_usecase.AuthorizationInput{
		UserData: userData,
		Code:     code,
		IpAddr:   ctx.IP(),
	})

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"data": output})
}

func (c *UserController) AuthCodeSend(ctx *fiber.Ctx) error {
	userData, isOk := ctx.Context().Value("userData").(global_type.UserTokenData)
	body := ctx.Context().Value("body").(map[string]any)

	account := body["receiveAccount"].(string)
	authType := body["authType"].(string)

	if !isOk {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid user token",
			"code":    "INVALID-USER-TOKEN",
		})
	}

	_, err := c.usecase.AuthCodeSend(user_usecase.AuthCodeSendInput{
		UserId:         userData.UserId,
		ReceiveAccount: account,
		AuthType:       authType,
		IpAddr:         ctx.IP(),
	})

	if err == nil {
		return err
	}

	return ctx.JSON(fiber.Map{"data": "success"})
}

func (c *UserController) ApiKeyIssue(ctx *fiber.Ctx) error {
	// todo : api key 발급 로직 추가
	return ctx.JSON(fiber.Map{"data": "success"})
}
