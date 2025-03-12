package alarm_entity

import "nspark-cron-alarm.com/cron-alarm-server/app/internal/entity/common_entity"

type CustomerAlarmMethod struct {
	ApiKey   string `db:"api_key"` // pk
	UserId   int    `db:"user_id"`
	MethodId int    `db:"method"` // fk alarm_method.id
	common_entity.Timestamp
}

type AlarmMethod struct {
	Id     int    `db:"id"`
	Method string `db:"method"`
	Status int    `db:"status"` // 0: 비활성화, 1: 활성화
	common_entity.Timestamp
}

type EmailAlarmInformation struct {
	ApiKey     string `db:"api_key"` // pk
	ToEmail    string `db:"to_email"`
	Title      string `db:"title"`
	CustomForm string `db:"custom_form"`
	common_entity.Timestamp
}
