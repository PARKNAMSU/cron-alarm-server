package middleware

import (
	"encoding/json"
	"errors"
	"reflect"
	"time"

	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/config"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/global_type"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/user_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/encrypt_tool"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/jwt_tool"
)

var (
	middleware *Middleware
)

type Middleware struct {
	userRepository user_repository.UserRepositoryImpl
}

func NewMiddleware(userRepository user_repository.UserRepositoryImpl) *Middleware {
	if middleware == nil {
		middleware =
			&Middleware{
				userRepository: userRepository,
			}
	}
	return middleware
}

// 유저 검증 미들웨어
func (m *Middleware) UserValidation(c *fiber.Ctx) error {
	accessToken := c.Get("access-token")
	refreshToken := c.Get("refresh-token")

	userData, err := jwt_tool.GetData[global_type.UserTokenData](accessToken, config.JWT_ACCESS_TOKEN_KEY)

	if err != nil {
		tokenData := m.userRepository.GetRefreshToken(refreshToken)

		if tokenData == nil || tokenData.Status == 0 {
			return errors.New("invalid refresh token")
		}

		userData, err = jwt_tool.GetData[global_type.UserTokenData](refreshToken, config.JWT_REFRESH_TOKEN_KEY)
		accessToken = jwt_tool.GenerateToken(userData, config.JWT_ACCESS_TOKEN_KEY, time.Minute*30)

		if err != nil {
			return err
		}
	}

	c.Response().Header.Set("access-token", accessToken)
	c.Response().Header.Set("refresh-token", refreshToken)

	c.Context().SetUserValue("userData", userData)

	return c.Next()
}

func (m *Middleware) APIKeyValidation(c *fiber.Ctx) error {
	headerKey := c.Get("x-api-key", "")

	if headerKey == "" {
		return errors.New("api key not exist")
	}

	decryptKey, err := encrypt_tool.Decrypt(headerKey, config.USER_API_ENCRYPT_KEY)
	headerKey = string(decryptKey)

	if err != nil {
		return err
	}

	key := m.userRepository.GetUserApiKey(user_repository.GetUserApiKeyInput{
		SearchType: user_repository.GET_USER_API_KEY_API_KEY,
		ApiKey:     &headerKey,
	})

	if key == nil {
		return errors.New("user api key not found")
	}

	return c.Next()
}

func (m *Middleware) BodyParsor(c *fiber.Ctx) error {
	if len(c.Body()) == 0 {
		return c.Next()
	}

	body := make(map[string]any)
	json.Unmarshal(c.Body(), &body)

	c.Context().SetUserValue("body", body)
	return c.Next()
}

func (m *Middleware) BodyValidator(key string, dataType config.REQUEST_DATA_TYPE) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid body",
			"code":    "INVALID-BODY",
		})
		if len(c.Body()) == 0 {
			return err
		}
		var body map[string]any
		bodyErr := json.Unmarshal(c.Body(), &body)

		if bodyErr != nil {
			return err
		}

		validation := false

		bodyData, ok := body[key]

		if !ok {
			return err
		}

		switch dataType {
		case config.REQUEST_DATA_TYPE_BOOL:
			validation =
				reflect.TypeOf(bodyData).Kind() == reflect.Bool
		case config.REQUEST_DATA_TYPE_INT:
			validation =
				reflect.TypeOf(bodyData).Kind() == reflect.Int ||
					reflect.TypeOf(bodyData).Kind() == reflect.Int8 ||
					reflect.TypeOf(bodyData).Kind() == reflect.Int16 ||
					reflect.TypeOf(bodyData).Kind() == reflect.Int32 ||
					reflect.TypeOf(bodyData).Kind() == reflect.Int64
		case config.REQUEST_DATA_TYPE_FLOAT:
			validation =
				reflect.TypeOf(body[key]).Kind() == reflect.Float32 ||
					reflect.TypeOf(body[key]).Kind() == reflect.Float64
		case config.REQUEST_DATA_TYPE_SLICE:
			validation =
				reflect.TypeOf(body[key]).Kind() == reflect.Slice
		case config.REQUEST_DATA_TYPE_STRING:
			validation =
				reflect.TypeOf(body[key]).Kind() == reflect.String
		case "":
			validation = true
		}

		if !validation {
			return err
		}
		return c.Next()
	}
}
