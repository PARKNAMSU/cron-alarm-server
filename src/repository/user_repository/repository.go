package user_repository

import "nspark-cron-alarm.com/cron-alarm-server/src/repository/common_repository"

type UserRepository struct {
	common_repository.Repository
}

var (
	Repository *UserRepository
)

func (r *UserRepository) GetUserByEmail(email string) {
	// todo: email 통해 유저데이터 가져오기 구현
}

func (r *UserRepository) CreateUser(input CreateUserInput) {
	// todo: 유저데이터 생성 구현
}

func (r *UserRepository) SetUserLoginDataInput(input SetUserLoginDataInput) {
	// todo: 유저 로그인 데이터 세팅 구현
}
