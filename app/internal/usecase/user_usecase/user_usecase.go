package user_usecase

import (
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/user_repository"
)

type UserUsecaseImpl interface {
}

type userUsecase struct {
	userRepository user_repository.UserRepositoryImpl
}

var (
	usecase *userUsecase
)

func NewUsecase(userRepo user_repository.UserRepositoryImpl) UserUsecaseImpl {
	if usecase == nil {
		usecase = &userUsecase{
			userRepository: userRepo,
		}
	}
	return usecase
}
