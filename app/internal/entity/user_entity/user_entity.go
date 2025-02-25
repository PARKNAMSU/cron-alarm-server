package user_entity

import (
	"database/sql"

	"nspark-cron-alarm.com/cron-alarm-server/src/entity/common_entity"
)

type UserEntity struct { // table: user
	Id     int    `db:"id"`
	Method string `db:"method"` // normal: 일반 유저, oauth: 소셜 로그인 유저
	Status int    `db:"status"`
	IpAddr string `db:"ip_addr"`
	common_entity.Timestamp
}

type UserLoginDataEntity struct { // table: user_login_data (유저 로그인 인증 정보)
	UserId   int    `db:"user_id"`
	Email    string `db:"email"`
	Password string `db:"password"`
	common_entity.Timestamp
}

type UserOauthEntity struct { // table: user_oauth (유저 소셜 로그인 정보)
	UserId    int    `db:"user_id"`
	OauthId   string `db:"oauth_id"`
	OauthHost string `db:"oauth_host"`
	common_entity.Timestamp
}

type UserInformation struct { // table: user_information (유저 데이터 전달용 정보)
	UserId   int            `db:"user_id"`
	Email    sql.NullString `db:"email"`
	Name     sql.NullString `db:"name"`
	Auth     int            `db:"auth"`
	AuthType sql.NullString `db:"auth_type"`
	common_entity.Timestamp
}

// refresh token 검증을 위한 테이블
type UserRefreshTokenEntity struct { // table: user_refresh_token
	UserId int    `db:"user_id"`
	Token  string `db:"token"`
	IpAddr string `db:"ip_addr"`
	common_entity.Timestamp
}
