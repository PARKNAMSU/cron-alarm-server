package common_repository

import "nspark-cron-alarm.com/cron-alarm-server/src/infra/database"

type Repository struct {
	MasterDB *database.CustomDB
	SlaveDB  *database.CustomDB
}

func (r *Repository) InitRepository() {
	r = &Repository{
		MasterDB: database.GetMysqlMaster(true),
		SlaveDB:  database.GetMysqlSlave(),
	}
}
