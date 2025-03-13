package log_entity

import "nspark-cron-alarm.com/cron-alarm-server/app/internal/entity/common_entity"

type LogUserAuthCodeEntity struct {
	Id             int    `db:"id"`
	UserId         int    `db:"user_id"`
	ReceiveAccount string `db:"receive_account"`
	AuthType       string `db:"auth_type"` // email, phone
	Code           string `db:"code"`
	Action         string `db:"action"` // auth: 인증, password: 비밀번호 찾기
	Status         *int   `db:"status"` // 인증 액션에서만 사용. 0: 인증 전, 1: 인증 완료
	IpAddr         string `db:"ip_addr"`
	common_entity.Timestamp
}
