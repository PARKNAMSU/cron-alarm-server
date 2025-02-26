package user_repository

import "time"

type GetUserSelectKeyType uint8

var (
	GET_USER_KEY_EMAIL GetUserSelectKeyType = 0
	GET_USER_KEY_ID    GetUserSelectKeyType = 1
)

type CreateUserInput struct {
	IpAddr string
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
	UserId int
	Token  string
	IpAddr string
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
	SelectKeyType GetUserSelectKeyType
}

type GetUserOutput struct {
	UserId    int
	Method    string
	Status    int
	IpAddr    string
	Email     *string
	Password  *string
	Name      *string
	OauthId   *string
	OauthHost *string
	Auth      int
	AuthType  *string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type DeleteUserInput struct {
	UserId int
}
