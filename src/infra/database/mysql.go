package database

import (
	"fmt"
	"os"
)

func getMysqlConnect(config dbConfig) string {
	return fmt.Sprintf(
		`%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&maxAllowedPacket=%d`,
		config.user,
		config.password,
		config.host,
		config.database,
		config.charset,
		config.maxAllowedPacket,
	)
}

func GetMysqlMaster(isTransaction bool) *CustomDB {
	mysqlMasterDB := &CustomDB{
		engine: mysqlEngine,
		config: dbConfig{
			user:             os.Getenv("MYSQL_USER_MASTER"),
			password:         os.Getenv("MYSQL_PASSWORD_MASTER"),
			host:             os.Getenv("MYSQL_HOST_MASTER"),
			database:         os.Getenv("MYSQL_DATABASE_MASTER"),
			charset:          "utf8mb4",
			maxAllowedPacket: 0,
		},
		isTransaction: isTransaction,
	}
	mysqlMasterDB.Connect()
	return mysqlMasterDB
}

func GetMysqlSlave() *CustomDB {
	mysqlSlaveDB := &CustomDB{
		engine: mysqlEngine,
		config: dbConfig{
			user:             os.Getenv("MYSQL_USER_SLAVE"),
			password:         os.Getenv("MYSQL_PASSWORD_SLAVE"),
			host:             os.Getenv("MYSQL_HOST_SLAVE"),
			database:         os.Getenv("MYSQL_DATABASE_SLAVE"),
			charset:          "utf8mb4",
			maxAllowedPacket: 0,
		},
		isTransaction: false,
	}
	mysqlSlaveDB.Connect()
	return mysqlSlaveDB
}
