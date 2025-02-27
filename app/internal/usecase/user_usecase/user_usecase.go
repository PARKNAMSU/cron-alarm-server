package user_usecase

import (
	"errors"
	"time"

	"nspark-cron-alarm.com/cron-alarm-server/app/config"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/global_type"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/user_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/common_tool"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/encrypt_tool"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/jwt_tool"
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
	userData := global_type.UserTokenData{
		Email: &input.Email,
	}

	if id, err := u.userRepository.CreateUser(user_repository.CreateUserInput{
		IpAddr: input.IpAddr,
	}); err != nil {
		return nil, err
	} else {
		output.UserId = id
		userData.UserId = id
	}

	encryptPassword, err := encrypt_tool.Encrypt([]byte(input.Password), config.USER_PASSWORD_ENCRYPT_KEY)

	if err != nil {
		return nil, err
	}

	errors := common_tool.ParallelExec(
		func() error {
			return u.userRepository.SetUserInformation(user_repository.SetUserInformationInput{
				UserId: output.UserId,
				Email:  &input.Email,
			})
		},
		func() error {
			return u.userRepository.SetUserLoginData(user_repository.SetUserLoginDataInput{
				UserId:   output.UserId,
				Email:    input.Email,
				Password: encryptPassword,
			})
		},
	)

	if len(errors) > 0 {
		return nil, errors[0]
	}

	userData.Method = "normal"
	userData.Status = 1
	userData.IpAddr = input.IpAddr
	userData.Password = &encryptPassword
	userData.Auth = 0
	userData.CreatedAt = time.Now()

	accessToken := jwt_tool.GenerateToken(userData, config.JWT_ACCESS_TOKEN_KEY, config.JWT_ACCESS_TOKEN_PERIOD)
	refreshToken := jwt_tool.GenerateToken(userData, config.JWT_REFRESH_TOKEN_KEY, config.JWT_REFRESH_TOKEN_PERIOD)

	if err := u.userRepository.SetUserRefreshToken(user_repository.SetUserRefreshTokenInput{
		UserId:    output.UserId,
		Token:     refreshToken,
		IpAddr:    input.IpAddr,
		ExpiredAt: time.Now().Add(config.JWT_REFRESH_TOKEN_PERIOD),
	}); err != nil {
		return nil, err
	}

	return &SignInOutput{
		UserId:       output.UserId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
