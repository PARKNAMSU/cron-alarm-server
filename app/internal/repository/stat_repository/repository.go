package stat_repository

import (
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/root_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/database"
)

type StatRepositoryImpl interface {
	root_repository.RootRepositoryImpl
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
