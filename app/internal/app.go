package internal

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/di"
	v1 "nspark-cron-alarm.com/cron-alarm-server/app/internal/router/v1"
	v2 "nspark-cron-alarm.com/cron-alarm-server/app/internal/router/v2"
)

var (
	appCfg = fiber.Config{
		ServerHeader: "Fiber",
		AppName:      "cron alarm service",
		Prefork:      true,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			if err != nil {
				fmt.Println("[Unexpected Error]:[", err.Error(), "]")
				return ctx.Status(code).SendString("Internal Server Error")
			}
			return nil
		},
	}
	middleware = di.InitMiddleware()
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

	app.Route("/api/v1", v1.Router())
	app.Route("/api/v2", v2.Router())

	app.Hooks().OnListen(func(listenData fiber.ListenData) error {
		fmt.Printf("[Server is running on]:[%s:%s]\n", listenData.Host, listenData.Port)
		return nil
	})

	app.Hooks().OnShutdown(func() error {
		fmt.Printf("Server is shutting down [time ]:[%s]\n", time.Now().Format("2006-01-02 15:04:05"))
		return nil
	})

	app.Listen(":8080")

	return app
}
