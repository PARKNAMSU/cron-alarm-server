package user_repository

import "nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/common_repository"

type UserRepository struct {
	common_repository.Repository
}

var (
	Repository *UserRepository
)

func (r *UserRepository) GetUser(input GetUserInput) *GetUserOutput {
	// todo: email 통해 유저데이터 가져오기 구현
	return &GetUserOutput{}
}

func (r *UserRepository) CreateUser(input CreateUserInput) error {
	// todo: 유저데이터 생성 구현
	return nil
}

func (r *UserRepository) SetUserLoginData(input SetUserLoginDataInput) error {
	// todo: 유저 로그인 데이터 세팅 구현
	return nil
}

func (r *UserRepository) SetUserOauth(input SetUserOauthInput) error {
	// todo:  oauth 유저 데이터 세팅 구현
	return nil
}

func (r *UserRepository) SetUserInformation(input SetUserInformationInput) error {
	// todo: 유저 정보 세팅 구현
	return nil
}

func (r *UserRepository) Authorization(input AuthorizationInput) error {
	// todo: 유저 인증 세팅
	return nil
}

func (r *UserRepository) SetUserRefreshToken(input SetUserRefreshTokenInput) error {
	// todo: 유저 refresh token 세팅
	return nil
}

func (r *UserRepository) DeleteUser(input DeleteUserInput) error {
	// todo: 유저 delete 세팅
	return nil
}
