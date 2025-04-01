package types

import "time"

type SelectKeyType uint8

type UserTokenData struct {
	UserId         int        `json:"userId"`
	Method         string     `json:"method"`
	Status         int        `json:"status"`
	IpAddr         string     `json:"ipAddr"`
	Email          *string    `json:"email,omitempty"`
	Password       *string    `json:"password,omitempty"`
	Name           *string    `json:"name"`
	OauthId        *string    `json:"oauthId,omitempty"`
	OauthHost      *string    `json:"oauthHost,omitempty"`
	Auth           int        `json:"auth"`
	AuthType       *string    `json:"authType"`
	MaxPlatformCnt int        `json:"maxPlatformCnt"`
	Grade          string     `json:"grade"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}

type CustomError struct {
	Code   string `json:"code"`
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}
