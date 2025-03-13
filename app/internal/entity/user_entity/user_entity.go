package user_entity

import (
	"database/sql"
	"time"

	"nspark-cron-alarm.com/cron-alarm-server/app/internal/entity/common_entity"
)

type UserEntity struct { // table: user
	Id     int    `db:"id"`
	Method string `db:"method"` // normal: 일반 유저, oauth: 소셜 로그인 유저
	Status int    `db:"status"` // 0: 탈퇴, 1: 정상, 2: 정지
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

type UserInformationEntity struct { // table: user_information (유저 데이터 전달용 정보)
	UserId   int            `db:"user_id"`
	Email    sql.NullString `db:"email"`
	Name     sql.NullString `db:"name"`
	Auth     int            `db:"auth"`
	AuthType sql.NullString `db:"auth_type"`
	common_entity.Timestamp
}

type UserDataEntity struct {
	UserId    int     `db:"user_id"`
	Email     *string `db:"email"`
	Password  *string `db:"password"`
	Method    string  `db:"method"` // normal: 일반 유저, oauth: 소셜 로그인 유저
	Status    int     `db:"status"`
	IpAddr    string  `db:"ip_addr"`
	Name      *string `db:"name"`
	Auth      int     `db:"auth"`
	AuthType  *string `db:"auth_type"`
	OauthId   *string `db:"oauth_id"`
	OauthHost *string `db:"oauth_host"`
	common_entity.Timestamp
}

// refresh token 검증을 위한 테이블 - 상태 및 ip 주소 같이 저장하여 유효 토큰 여부 관리
type UserRefreshTokenEntity struct { // table: user_refresh_token
	Token     string `db:"token"`
	UserId    int    `db:"user_id"`
	Status    int    `db:"status"`     // 1: 사용 가능, 0: 탈취됨
	IpAddr    string `db:"ip_addr"`    // 토큰 발급 당시의 IP 주소
	ExpiredAt string `db:"expired_at"` // 토큰 만료 시간
	common_entity.Timestamp
}

type UserApiKeyEntity struct { // table: user_api_key
	ApiKey    string    `db:"api_key"`
	Status    int       `db:"status"`   // 1: 사용 가능, 0: 사용 중지
	Hostname  string    `db:"hostname"` // 해당 키를 사용하는 호스트 이름
	UserId    int       `db:"user_id"`
	ExpiredAt time.Time `db:"expired_at"`
	common_entity.Timestamp
}

type UserAuthCodeEntity struct { // table: user_auth_code
	UserId         int       `db:"user_id"`         // pk
	ReceiveAccount string    `db:"receive_account"` // pk
	Action         string    `db:"action"`          // pk 인증 코드 사용 목적 auth : 계정 인증, password : 비밀번호 변경
	AuthType       string    `db:"auth_type"`       // email, phone
	Code           string    `db:"code"`
	Status         *int      `db:"status"` // 0: 인증 전, 1: 인증 완료
	ExpiredAt      time.Time `db:"expired_at"`
	common_entity.Timestamp
}
