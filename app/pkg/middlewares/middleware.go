package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/response"
)

// database 를 사용하지 않는 간단한 미들웨어는 여기에 생성한다.

var validate = validator.New()

func typeCheck[T any](fl validator.FieldLevel) bool {
	value := fl.Field().Interface()
	switch value.(type) {
	case T:
		return true
	}
	return false
}

func ValidateInit() {
	validate.RegisterValidation("string", typeCheck[string])
	validate.RegisterValidation("int", typeCheck[int])
	validate.RegisterValidation("bool", typeCheck[bool])
}

func BodyParsor[T any]() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var req T

		if err := c.BodyParser(&req); err != nil {
			return response.CustomError(
				c,
				fiber.Map{
					"error": "Invalid request body",
				},
				fiber.StatusBadRequest,
			)
		}

		// 유효성 검사 수행
		if err := validate.Struct(&req); err != nil {
			return response.CustomError(
				c,
				fiber.Map{
					"error": "Invalid request body",
				},
				fiber.StatusBadRequest,
			)
		}

		c.Context().SetUserValue("body", req)

		return c.Next()
	}
}
