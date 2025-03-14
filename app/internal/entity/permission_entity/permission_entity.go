package permission_entity

import "nspark-cron-alarm.com/cron-alarm-server/app/internal/entity/common_entity"

type UserPermissionEntity struct {
	Id             int    `db:"id"`
	MaxPlatformCnt int    `db:"max_platform_cnt"` // 최대 플랫폼 개수
	MaxAlarmCnt    int    `db:"max_alarm_cnt"`    // 최대 알람 개수
	Grade          string `db:"grade"`            // 권한 이름
	common_entity.Timestamp
}

type PlatformPermissionEntity struct {
	Id          int `db:"id"`
	MaxAlarmCnt int `db:"max_alarm_cnt"` // 플랫폼 별 최대 알람 개수
	common_entity.Timestamp
}
