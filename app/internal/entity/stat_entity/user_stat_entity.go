package stat_entity

type StatPlatformAlarmByDateEntity struct {
	Date     string `db:"date"`      // 일별 날짜 (YYYY-MM-DD)
	UserId   int    `db:"user_id"`   // 사용자 아이디
	Hostname string `db:"host_name"` // 사이트 이름
	AlarmCnt int    `db:"alarm_cnt"` // 알람 발생 횟수
}
