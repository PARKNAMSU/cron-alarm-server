package alarm_entity

import "nspark-cron-alarm.com/cron-alarm-server/app/internal/entity/common_entity"

type AlarmMethod struct {
	ApiKey string `db:"api_key"` // pk
	UserId int    `db:"user_id"`
	Method string `db:"method"` // kafka , discord , email , sms, sqs
	common_entity.Timestamp
}
