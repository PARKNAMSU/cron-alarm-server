package user_usecase

import (
	"time"

	"nspark-cron-alarm.com/cron-alarm-server/app/internal/global_type"
)

type SignUpInput struct {
	Email    string
	Password string
	IpAddr   string
}

type SignUpOutput struct {
	global_type.UserTokenData
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignInInput struct {
	Email    string
	Password string
	IpAddr   string
}

type SignInOutput struct {
	global_type.UserTokenData
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthorizationInput struct {
	UserData global_type.UserTokenData
	Code     string
	IpAddr   string
}

type AuthorizationOutput struct {
	global_type.UserTokenData
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthCodeSendInput struct {
	UserId         int
	ReceiveAccount string
	AuthType       string
	IpAddr         string
}

type AuthCodeSendOutput struct {
	Message string `json:"message"`
}

type ApiKeyIssueInput struct {
	UserData global_type.UserTokenData
	Hostname string
	IpAddr   string
}

type ApiKeyIssueOutput struct {
	ApiKey    string    `json:"apiKey"`
	ExpiredAt time.Time `json:"expiredAt"`
	Hostname  string    `json:"hostname"`
}
