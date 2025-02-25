package internal

import (
	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/middleware"
	v1 "nspark-cron-alarm.com/cron-alarm-server/app/internal/router/v1"
	v2 "nspark-cron-alarm.com/cron-alarm-server/app/internal/router/v2"
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

	app.Use(middleware.APIValidation)

	app.Use(func(c *fiber.Ctx) error {
		c.Accepts("html")                           // "html"
		c.Accepts("text/html")                      // "text/html"
		c.Accepts("json", "text")                   // "json"
		c.Accepts("application/json")               // "application/json"
		c.Accepts("text/plain", "application/json") // "application/json", due to quality
		return c.Next()
	})

	app.Route("/api/v1", v1.APIV1Router())
	app.Route("/api/v2", v2.APIV2Router())

	app.Listen(":8080")

	return app
}
