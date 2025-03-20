package permission_entity

import "nspark-cron-alarm.com/cron-alarm-server/app/internal/entity/common_entity"

type UserPermissionEntity struct {
	Id             int    `db:"id"`
	MaxPlatformCnt int    `db:"max_platform_cnt"` // 최대 플랫폼 개수
	Grade          string `db:"grade"`            // 권한 이름
	common_entity.Timestamp
}

type PlatformPermissionEntity struct {
	Id           int `db:"id"`
	MaxAlertCnt  int `db:"max_alert_cnt"`  // 플랫폼 별 최대 알람 개수
	MaxMethodCnt int `db:"max_method_cnt"` // 플랫폼 별 최대 알람 전송 방법 개수
	common_entity.Timestamp
}
