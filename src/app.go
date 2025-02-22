package src

import (
	"github.com/gofiber/fiber/v2"
	v1 "nspark-cron-alarm.com/cron-alarm-server/src/router/v1"
	v2 "nspark-cron-alarm.com/cron-alarm-server/src/router/v2"
)

var (
	app *fiber.App
)

var (
	appCfg = fiber.Config{
		ServerHeader: "Fiber",
		AppName:      "cron alarm service",
		Prefork:      true,
	}
)

func GetApp() *fiber.App {
	app := fiber.New(appCfg)

	app.Route("/api/v1", v1.APIV1Router())
	app.Route("/api/v2", v2.APIV2Router())

	app.Listen(":8080")

	return app
}
