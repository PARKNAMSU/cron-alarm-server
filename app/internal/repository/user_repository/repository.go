package user_repository

import (
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/database"
)

type UserRepositoryImpl interface {
	GetUser(input GetUserInput) *GetUserOutput
	CreateUser(input CreateUserInput) error
	SetUserLoginData(input SetUserLoginDataInput) error
	SetUserOauth(input SetUserOauthInput) error
	SetUserInformation(input SetUserInformationInput) error
	Authorization(input AuthorizationInput) error
	SetUserRefreshToken(input SetUserRefreshTokenInput) error
	DeleteUser(input DeleteUserInput) error
}

type userRepository struct {
	masterDB *database.CustomDB
	slaveDB  *database.CustomDB
}

var (
	repo *userRepository
)

func NewRepository(masterDB *database.CustomDB, slaveDB *database.CustomDB) UserRepositoryImpl {
	if repo == nil {
		repo = &userRepository{
			masterDB: masterDB,
			slaveDB:  slaveDB,
		}
	}
	return repo
}

func (r *userRepository) GetUser(input GetUserInput) *GetUserOutput {
	// todo: email 통해 유저데이터 가져오기 구현
	return &GetUserOutput{}
}

func (r *userRepository) CreateUser(input CreateUserInput) error {
	// todo: 유저데이터 생성 구현
	return nil
}

func (r *userRepository) SetUserLoginData(input SetUserLoginDataInput) error {
	// todo: 유저 로그인 데이터 세팅 구현
	return nil
}

func (r *userRepository) SetUserOauth(input SetUserOauthInput) error {
	// todo:  oauth 유저 데이터 세팅 구현
	return nil
}

func (r *userRepository) SetUserInformation(input SetUserInformationInput) error {
	// todo: 유저 정보 세팅 구현
	return nil
}

func (r *userRepository) Authorization(input AuthorizationInput) error {
	// todo: 유저 인증 세팅
	return nil
}

func (r *userRepository) SetUserRefreshToken(input SetUserRefreshTokenInput) error {
	// todo: 유저 refresh token 세팅
	return nil
}

func (r *userRepository) DeleteUser(input DeleteUserInput) error {
	// todo: 유저 delete 세팅
	return nil
}
