package log_repository

import (
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/root_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/database"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/query_tool"
)

type LogRepositoryImpl interface {
	root_repository.RootRepositoryImpl
	InsertLogUserAuthCode(input InsertLogUserAuthCodeInput) error
}

type logRepository struct {
	root_repository.RootRepository
}

var (
	repo *logRepository
)

func NewRepository(masterDB *database.CustomDB, slaveDB *database.CustomDB) LogRepositoryImpl {
	if repo == nil {
		repo = &logRepository{}
		repo.SetMasterDB(masterDB)
		repo.SetSlaveDB(slaveDB)
	}
	return repo
}

func (r *logRepository) InsertLogUserAuthCode(input InsertLogUserAuthCodeInput) error {
	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["logUserAuthCode"],
		Action: query_tool.INSERT,
		Set: map[string]any{
			"user_id":         input.UserId,
			"receive_account": input.ReceiveAccount,
			"auth_type":       input.AuthType,
			"code":            input.Code,
			"action":          input.Action,
			"status":          input.Status,
			"ip_addr":         input.IpAddr,
		},
	})
	_, err := r.GetMasterDB().QueryExecute(query, params)
	return err
}
