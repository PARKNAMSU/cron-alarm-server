package v1

import (
	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/di"
)

// api key 를 보유한 경우 외부에서도 해당 라우터로 접근 가능
func OpenRouter() func(router fiber.Router) {
	middleware := di.InitMiddleware()
	return func(router fiber.Router) {
		router.Use(middleware.APIKeyValidation)
	}
}
