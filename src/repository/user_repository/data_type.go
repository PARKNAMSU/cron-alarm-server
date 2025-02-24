package user_repository

import "time"

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

type GetUserOutput struct {
	UserId    int
	Method    string
	Status    string
	IpAddr    string
	Email     *string
	Password  *string
	Name      *string
	OauthId   *string
	OauthHost *string
	Auth      int
	AuthType  *string
	createdAt time.Time
	updatedAt *time.Time
}
