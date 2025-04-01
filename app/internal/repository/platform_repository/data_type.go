package platform_repository

import (
	"time"

	"nspark-cron-alarm.com/cron-alarm-server/app/internal/types"
)

var (
	GET_PLATFORM_USER_ID types.SelectKeyType = 0
	GET_PLATFORM_API_KEY types.SelectKeyType = 1
	GET_PLATFORM_HOST    types.SelectKeyType = 2
)

type InserPlatformInput struct {
	Hostname  string
	ApiKey    string
	ExpiredAt time.Time
	UserId    int
}

type UpdatePlatformInput struct {
	PlatformName *string
	ExpiredAt    *time.Time
	Hostname     string
	UserId       int
}

type GetPlatformInput struct {
	UserId      *int
	ApiKey      *string
	Hostname    *string
	SearchType  types.SelectKeyType
	IsGetUsable bool
}

type GetPlatformOutput struct {
	UserId       int
	ApiKey       string
	Status       int
	Hostname     string
	PlatformName string
}
