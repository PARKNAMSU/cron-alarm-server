package platform_usecase

import (
	"time"

	"nspark-cron-alarm.com/cron-alarm-server/app/internal/types"
)

type ApiKeyIssueInput struct {
	UserData types.UserTokenData
	Hostname string
	IpAddr   string
}

type ApiKeyIssueOutput struct {
	ApiKey    string    `json:"apiKey"`
	ExpiredAt time.Time `json:"expiredAt"`
	Hostname  string    `json:"hostname"`
}
