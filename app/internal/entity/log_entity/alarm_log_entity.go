package log_entity

import "nspark-cron-alarm.com/cron-alarm-server/app/internal/entity/common_entity"

type AlarmLog struct {
	Id       int    `db:"id"`
	ApiKey   string `db:"api_key"`
	UserId   int    `db:"user_id"`
	MethodId int    `db:"method_id"`
	Message  string `db:"message"`
	IpAddr   string `db:"ip_addr"`
	common_entity.Timestamp
}
