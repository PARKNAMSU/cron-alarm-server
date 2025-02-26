package middleware

import (
	"errors"
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
	headerApiKey := c.Get("x-api-key")
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

	key := m.userRepository.GetUserApiKey(userData.UserId)

	if key == nil {
		return errors.New("user api key not found")
	}

	userKeyBytes, _ := encrypt_tool.Decrypt(headerApiKey, config.USER_API_ENCRYPT_KEY)
	requestBytes, _ := encrypt_tool.Decrypt(headerApiKey, config.USER_API_ENCRYPT_KEY)

	if string(userKeyBytes) != string(requestBytes) {
		return errors.New("invalid api key")
	}

	c.Response().Header.Set("access-token", accessToken)
	c.Response().Header.Set("refresh-token", refreshToken)

	c.Context().SetUserValue("userData", userData)

	return c.Next()
}
