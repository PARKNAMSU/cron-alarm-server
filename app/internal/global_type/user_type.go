package global_type

import "time"

type UserTokenData struct {
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
	CreatedAt time.Time
	UpdatedAt *time.Time
}
