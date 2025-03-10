package log_repository

import (
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/root_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/database"
)

type LogRepositoryImpl interface {
	root_repository.RootRepositoryImpl
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
