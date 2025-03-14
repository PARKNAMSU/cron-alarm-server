package database

import (
	"fmt"
	"os"
)

type TableKey = string

var (
	MYSQL_TABLE = map[TableKey]string{
		"user":             "user",
		"userInformation":  "user_information",
		"userOauth":        "user_oauth",
		"userLoginData":    "user_login_data",
		"userRefreshToken": "user_refresh_token",
		"userApiKey":       "user_api_key",
		"userAuthCode":     "user_auth_code",
		"admin":            "admin",

		"alarmMethod": "alarm_method",
		"alarmLog":    "alarm_log",

		"taskLog": "task_log",

		"logUserAuthCode": "log_user_auth_code",

		"statPlatformAlarm": "stat_platform_alarm",
	}
)

var (
	mysqlMasterDB *CustomDB
	mysqlSlaveDB  *CustomDB
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
	if mysqlMasterDB == nil {
		mysqlMasterDB = &CustomDB{
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
		mysqlMasterDB.connect()
		if isTransaction {
			mysqlMasterDB.transaction()
		}
	}
	return mysqlMasterDB
}

func GetMysqlSlave() *CustomDB {
	if mysqlSlaveDB == nil {
		mysqlSlaveDB = &CustomDB{
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
		mysqlSlaveDB.connect()
	}
	return mysqlSlaveDB
}
