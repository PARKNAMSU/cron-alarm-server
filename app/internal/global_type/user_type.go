package global_type

import "time"

type UserTokenData struct {
	UserId    int        `json:"userId"`
	Method    string     `json:"method"`
	Status    int        `json:"status"`
	IpAddr    string     `json:"ipAddr"`
	Email     *string    `json:"email,omitempty"`
	Password  *string    `json:"password,omitempty"`
	Name      *string    `json:"name"`
	OauthId   *string    `json:"oauthId,omitempty"`
	OauthHost *string    `json:"oauthHost,omitempty"`
	Auth      int        `json:"auth"`
	AuthType  *string    `json:"authType"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
