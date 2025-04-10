package platform_entity

import (
	"time"

	"nspark-cron-alarm.com/cron-alarm-server/app/internal/entity/common_entity"
)

type PlatformEntity struct { // table: user_api_key
	Hostname     string    `db:"hostname"`      // 해당 키를 사용하는 호스트 이름
	ApiKey       string    `db:"api_key"`       // 플랫폼 인증 api key
	Status       int       `db:"status"`        // 1: 사용 가능, 0: 사용 중지
	PlatformName string    `db:"platform_name"` // 플랫폼 이름
	UserId       int       `db:"user_id"`
	ExpiredAt    time.Time `db:"expired_at"`
	common_entity.Timestamp
}
