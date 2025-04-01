package user_repository

import (
	"time"

	"nspark-cron-alarm.com/cron-alarm-server/app/internal/types"
)

var (
	GET_USER_KEY_EMAIL types.SelectKeyType = 0
	GET_USER_KEY_ID    types.SelectKeyType = 1
)

type CreateUserInput struct {
	IpAddr string
	Method string
}

type SetUserLoginDataInput struct {
	UserId   int
	Email    string
	Password string
}

type SetUserOauthInput struct {
	UserId    int
	OauthId   string
	OauthHost string
}

type SetUserInformationInput struct {
	UserId int
	Email  *string
	Name   *string
}

type AuthorizationInput struct {
	UserId   int
	AuthType string
}

type SetUserRefreshTokenInput struct {
	UserId    int
	Token     string
	IpAddr    string
	ExpiredAt time.Time
}

type GetRefreshTokenInput struct {
	Token  string
	UserId int
	IpAddr string
	Status int
}

type GetUserInput struct {
	UserId        uint
	Email         string
	SelectKeyType types.SelectKeyType
}

type GetUserOutput struct {
	UserId         int
	Method         string
	Status         int
	IpAddr         string
	Email          *string
	Password       *string
	Name           *string
	OauthId        *string
	OauthHost      *string
	Auth           int
	AuthType       *string
	Grade          *string
	MaxPlatformCnt int
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

type DeleteUserInput struct {
	UserId int
}

type SetAuthCodeInput struct {
	UserId         int
	ReceiveAccount string
	AuthType       string
	Code           string
	Action         string
	ExpiredAt      time.Time
	Status         int
}

type GetAvailableAuthCodeOutput struct {
	UserId         int
	ReceiveAccount string
	Code           string
	Action         string
	AuthType       string
	Status         *int
}
