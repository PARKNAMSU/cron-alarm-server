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
	SignUp(input SignUpInput) (*SignUpOutput, error)
	SignIn(input SignInInput) (*SignInOutput, error)
	Authorization(input AuthorizationInput) (AuthorizationOutput, error)
	AuthCodeSend(input AuthCodeSendInput) AuthCodeSendOutput
	ApiKeyIssue(input ApiKeyIssueInput) ApiKeyIssueOutput
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

func (u *userUsecase) SignUp(input SignUpInput) (*SignUpOutput, error) {
	user := u.userRepository.GetUser(user_repository.GetUserInput{
		Email:         input.Email,
		SelectKeyType: user_repository.GET_USER_KEY_EMAIL,
	})

	if user != nil {
		return nil, errors.New("EXIST_USER")
	}

	output := &SignUpOutput{}
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

	output.AccessToken = accessToken
	output.RefreshToken = refreshToken
	output.Auth = userData.Auth
	output.AuthType = userData.AuthType
	output.Email = userData.Email
	output.IpAddr = userData.IpAddr
	output.Method = userData.Method
	output.Name = userData.Name
	output.CreatedAt = userData.CreatedAt
	output.UpdatedAt = userData.UpdatedAt

	return output, nil
}

func (u *userUsecase) SignIn(input SignInInput) (*SignInOutput, error) {
	// todo: 로그인 구현
	user := u.userRepository.GetUser(user_repository.GetUserInput{
		Email:         input.Email,
		SelectKeyType: user_repository.GET_USER_KEY_EMAIL,
	})

	if user == nil {
		return nil, errors.New("NOT_EXIST_USER")
	}

	decryptPassword, _ := encrypt_tool.Decrypt(*user.Password, config.USER_PASSWORD_ENCRYPT_KEY)

	if input.Password != string(decryptPassword) {
		return nil, errors.New("INVALID_PASSWORD")
	}

	output := &SignInOutput{}

	output.Email = &input.Email
	output.UserId = user.UserId
	output.Method = user.Method
	output.Status = user.Status
	output.IpAddr = user.IpAddr
	output.Name = user.Name
	output.Auth = user.Auth
	output.OauthId = user.OauthId
	output.OauthHost = user.OauthHost
	output.AuthType = user.AuthType
	output.CreatedAt = user.CreatedAt
	output.UpdatedAt = user.UpdatedAt

	userData := global_type.UserTokenData{
		Email:     &input.Email,
		UserId:    user.UserId,
		Method:    user.Method,
		Status:    user.Status,
		IpAddr:    user.IpAddr,
		Name:      user.Name,
		Auth:      user.Auth,
		OauthId:   user.OauthId,
		OauthHost: user.OauthHost,
		AuthType:  user.AuthType,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	accessToken := jwt_tool.GenerateToken(userData, config.JWT_ACCESS_TOKEN_KEY, config.JWT_ACCESS_TOKEN_PERIOD)
	refreshToken := jwt_tool.GenerateToken(userData, config.JWT_REFRESH_TOKEN_KEY, config.JWT_REFRESH_TOKEN_PERIOD)

	if err := u.userRepository.SetUserRefreshToken(user_repository.SetUserRefreshTokenInput{
		UserId:    user.UserId,
		Token:     refreshToken,
		IpAddr:    input.IpAddr,
		ExpiredAt: time.Now().Add(config.JWT_REFRESH_TOKEN_PERIOD),
	}); err != nil {
		return nil, err
	}

	output.AccessToken = accessToken
	output.RefreshToken = refreshToken
	output.UserId = user.UserId

	return output, nil
}

func (u *userUsecase) Authorization(input AuthorizationInput) (AuthorizationOutput, error) {
	// todo : 계정 인증 로직 추가
	return AuthorizationOutput{}, nil
}

func (u *userUsecase) AuthCodeSend(input AuthCodeSendInput) AuthCodeSendOutput {
	// todo : 인증 코드 전송 로직 추가
	return AuthCodeSendOutput{}
}

func (u *userUsecase) ApiKeyIssue(input ApiKeyIssueInput) ApiKeyIssueOutput {
	// todo : API 키 발급 로직 추가
	return ApiKeyIssueOutput{}
}
