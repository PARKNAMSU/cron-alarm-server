package v1

import (
	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/common"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/middlewares"
)

func UserRouter() func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Post(
			"/signUp",
			middlewares.BodyParsor[common.SignUpRequest](),
			userController.SignUp,
		)

		router.Post(
			"/signIn",
			middlewares.BodyParsor[common.SignInRequest](),
			userController.SignIn,
		)

		router.Post(
			"/auth/code",
			middlewares.BodyParsor[common.AuthCodeSendRequest](),
			userController.AuthCodeSend,
		)

	}
}
