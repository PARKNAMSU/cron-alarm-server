package root_repository

import "nspark-cron-alarm.com/cron-alarm-server/app/pkg/database"

type RootRepositoryImpl interface {
	Rollback()
	Commit()
	GetMasterDB() *database.CustomDB
	GetSlaveDB() *database.CustomDB
	SetMasterDB(db *database.CustomDB)
	SetSlaveDB(db *database.CustomDB)
}

type RootRepository struct {
	masterDB *database.CustomDB
	slaveDB  *database.CustomDB
}

func (r *RootRepository) Rollback() {
	r.masterDB.Rollback()
}

func (r *RootRepository) Commit() {
	r.masterDB.Commit()
}

func (r *RootRepository) GetMasterDB() *database.CustomDB {
	return r.masterDB
}

func (r *RootRepository) GetSlaveDB() *database.CustomDB {
	return r.slaveDB
}

func (r *RootRepository) SetMasterDB(db *database.CustomDB) {
	r.masterDB = db
}

func (r *RootRepository) SetSlaveDB(db *database.CustomDB) {
	r.slaveDB = db
}
