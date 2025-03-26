package v1

import (
	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/config"
)

func UserRouter() func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Post(
			"/signUp",
			middleware.BodyValidator("email", config.REQUEST_DATA_TYPE_STRING),
			middleware.BodyValidator("password", config.REQUEST_DATA_TYPE_STRING),
			userController.SignUp,
		)

		router.Post(
			"/signIn",
			middleware.BodyValidator("email", config.REQUEST_DATA_TYPE_STRING),
			middleware.BodyValidator("password", config.REQUEST_DATA_TYPE_STRING),
			userController.SignIn,
		)

		router.Post(
			"/auth/code",
			middleware.BodyValidator("receiveAccount", config.REQUEST_DATA_TYPE_STRING),
			middleware.BodyValidator("authType", config.REQUEST_DATA_TYPE_STRING),
			userController.AuthCodeSend,
		)

	}
}
