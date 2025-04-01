package platform_controller

import (
	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/usecase/platform_usecase"
)

type PlatformController struct {
	usecase platform_usecase.PlatformUsecaseImpl
}

var (
	controller *PlatformController
)

func NewController(usecase platform_usecase.PlatformUsecaseImpl) *PlatformController {
	if controller == nil {
		controller = &PlatformController{
			usecase: usecase,
		}
	}
	return controller
}

func (c *PlatformController) ApiKeyIssue(ctx *fiber.Ctx) error {
	// todo : api key 발급 로직 추가
	output, err := c.usecase.ApiKeyIssue(platform_usecase.ApiKeyIssueInput{})

	if err != nil {
		return ctx.Status(err.Status).JSON(fiber.Map{
			"message": err.Msg,
			"code":    err.Code,
		})
	}

	return ctx.JSON(output)
}
