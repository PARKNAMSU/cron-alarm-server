package di

import (
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/controller/user_controller"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/middleware"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/log_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/user_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/usecase/user_usecase"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/database"
)

/* 의존성 객체 정의 및 초기화 */

var ( // db 객체 초기화
	masterDB = database.GetMysqlMaster(true)
	slaveDB  = database.GetMysqlSlave()
)

var ( // repo 객체 초기화
	userRepository = user_repository.NewRepository(masterDB, slaveDB)
	logRepositroy  = log_repository.NewRepository(masterDB, slaveDB)
)

var ( // usecase 객체 초기화
	userUsecase = user_usecase.NewUsecase(
		userRepository,
		logRepositroy,
	)
)

func InitUserController() *user_controller.UserController {
	return user_controller.NewController(userUsecase)
}

func InitMiddleware() *middleware.Middleware {
	return middleware.NewMiddleware(userRepository)
}
