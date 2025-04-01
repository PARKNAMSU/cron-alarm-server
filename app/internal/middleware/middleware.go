package middleware

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"nspark-cron-alarm.com/cron-alarm-server/app/config"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/platform_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/user_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/types"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/encrypt_tool"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/jwt_tool"
)

var (
	middleware *Middleware
)

type Middleware struct {
	userRepository     user_repository.UserRepositoryImpl
	platformRepository platform_repository.PlatformRepositoryImpl
}

func NewMiddleware(
	userRepository user_repository.UserRepositoryImpl,
	platformRepository platform_repository.PlatformRepositoryImpl,
) *Middleware {
	if middleware == nil {
		middleware =
			&Middleware{
				userRepository:     userRepository,
				platformRepository: platformRepository,
			}
	}
	return middleware
}

// 유저 검증 미들웨어
func (m *Middleware) UserValidation(c *fiber.Ctx) error {
	accessToken := c.Get("access-token")
	refreshToken := c.Get("refresh-token")

	userData, err := jwt_tool.GetData[types.UserTokenData](accessToken, config.JWT_ACCESS_TOKEN_KEY)

	if err != nil {
		tokenData := m.userRepository.GetRefreshToken(refreshToken)

		if tokenData == nil || tokenData.Status == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid user token",
				"code":    "INVALID-USER-TOKEN",
			})
		}

		userData, err = jwt_tool.GetData[types.UserTokenData](refreshToken, config.JWT_REFRESH_TOKEN_KEY)
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

	hostname := c.Hostname()

	list := m.platformRepository.GetPlatform(platform_repository.GetPlatformInput{
		SearchType:  platform_repository.GET_PLATFORM_HOST,
		ApiKey:      &hostname,
		IsGetUsable: true,
	})

	if len(list) == 0 {
		return errors.New("user api key not found")
	}

	info := list[0]

	decryptKey, err := encrypt_tool.Decrypt(headerKey, config.USER_API_ENCRYPT_KEY)

	if err != nil {
		return err
	}

	headerKey = string(decryptKey)

	decryptKey, err = encrypt_tool.Decrypt(info.ApiKey, config.USER_API_ENCRYPT_KEY)

	if err != nil {
		return err
	}

	apiKey := string(decryptKey)

	if info.Status != 1 || headerKey != apiKey {
		return errors.New("invalid api key")
	}

	c.Context().SetUserValue("platformInformation", info)

	return c.Next()
}

func (m *Middleware) BodyParsor(c *fiber.Ctx) error {
	if len(c.Body()) == 0 {
		return c.Next()
	}

	body := make(map[string]any)
	json.Unmarshal(c.Body(), &body)

	// c.Context().SetUserValue("body", body)
	return c.Next()
}
