package platform_usecase

import (
	"time"

	"nspark-cron-alarm.com/cron-alarm-server/app/internal/common"
)

type ApiKeyIssueInput struct {
	UserData common.UserTokenData
	Hostname string
	IpAddr   string
}

type ApiKeyIssueOutput struct {
	ApiKey    string    `json:"apiKey"`
	ExpiredAt time.Time `json:"expiredAt"`
	Hostname  string    `json:"hostname"`
}

type SetPlatformInput struct {
	Hostname     string
	UserId       int
	PlatformName string
	Status       int
}
