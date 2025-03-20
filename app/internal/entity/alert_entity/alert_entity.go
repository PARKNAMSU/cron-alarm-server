package alarm_entity

import "nspark-cron-alarm.com/cron-alarm-server/app/internal/entity/common_entity"

type PlatformAlertMethod struct {
	Id       int    `db:"id"`
	Hostname string `db:"hostname"`
	UserId   int    `db:"user_id"`
	MethodId int    `db:"method"` // fk alert_method.id
	common_entity.Timestamp
}

type AlertMethod struct {
	Id     int    `db:"id"`
	Method string `db:"method"`
	Status int    `db:"status"` // 0: 비활성화, 1: 활성화
	common_entity.Timestamp
}

type EmailAlertInformation struct {
	PlatformMethodId string  `db:"platform_method_id"` // pk
	ToEmail          string  `db:"to_email"`
	Title            string  `db:"title"`
	CustomForm       *string `db:"custom_form"`
	common_entity.Timestamp
}

type DiscordAlertInformation struct {
	PlatformMethodId string `db:"platform_method_id"` // pk
	WebhookUrl       string `db:"webhook_url"`
	Title            string `db:"title"`
	common_entity.Timestamp
}
