package user_usecase

import (
	"errors"

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

func (u *userUsecase) SignIn(input SignInInput) (*SignInOutput, error) {
	user := u.userRepository.GetUser(user_repository.GetUserInput{
		Email:         input.Email,
		SelectKeyType: user_repository.GET_USER_KEY_EMAIL,
	})

	if user != nil {
		return nil, errors.New("EXIST_USER")
	}

	output := &SignInOutput{}

	if id, err := u.userRepository.CreateUser(user_repository.CreateUserInput{
		IpAddr: input.IpAddr,
	}); err != nil {
		return nil, err
	} else {
		output.UserId = id
	}

	// todo: 로그인 데이터 세팅 구현

	return &SignInOutput{}, nil
}
