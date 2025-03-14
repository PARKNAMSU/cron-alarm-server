package stat_repository

import (
	"time"

	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/root_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/database"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/query_tool"
)

type StatRepositoryImpl interface {
	root_repository.RootRepositoryImpl
	SetStatPlatformAlarmByDate(hostname string, userId int) error
}

type statRepository struct {
	root_repository.RootRepository
}

var (
	repo *statRepository
)

func NewRepository(masterDB *database.CustomDB, slaveDB *database.CustomDB) StatRepositoryImpl {
	if repo == nil {
		repo = &statRepository{}
		repo.SetMasterDB(masterDB)
		repo.SetSlaveDB(slaveDB)
	}
	return repo
}

func (r *statRepository) SetStatPlatformAlarmByDate(hostname string, userId int) error {
	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["statPlatformAlarm"],
		Action: query_tool.DUPLICATE,
		Set: map[string]any{
			"date":      time.Now().Format("2006-01-02"),
			"hostname":  hostname,
			"user_id":   userId,
			"alarm_cnt": 1,
		},
		Duplicate: map[string]any{
			"alarm_cnt": "`alarm_cnt` + 1",
		},
	})
	_, err := r.GetMasterDB().QueryExecute(query, params)
	return err
}

func (r *statRepository) GetTodayAlarmCnt(input GetTodayAlarmCntInput) int {
	cnts := []int{}
	queryParams := query_tool.QueryParams{
		Table: database.MYSQL_TABLE["statPlatformAlarm"],
		Where: map[string]any{
			"date":    time.Now().Format("2006-01-02"),
			"user_id": input.UserId,
		},
	}

	if input.Hostname != nil {
		queryParams.Where["hostname"] = *input.Hostname
		queryParams.Select = []string{"alarm_cnt AS cnt"}
	} else {
		queryParams.Select = []string{"SUM(alarm_cnt) AS cnt"}
		queryParams.Conditions = "GROUP BY `date`, `user_id`"
	}

	query, params := query_tool.QueryBuilder(queryParams)

	r.GetSlaveDB().QuerySelect(&cnts, query, params)

	if len(cnts) == 0 {
		return 0
	}
	return cnts[0]
}
