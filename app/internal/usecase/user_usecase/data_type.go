package user_usecase

import (
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/types"
)

type SignUpInput struct {
	Email    string
	Password string
	IpAddr   string
}

type SignUpOutput struct {
	types.UserTokenData
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignInInput struct {
	Email    string
	Password string
	IpAddr   string
}

type SignInOutput struct {
	types.UserTokenData
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthorizationInput struct {
	UserData types.UserTokenData
	Code     string
	IpAddr   string
}

type AuthorizationOutput struct {
	types.UserTokenData
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
