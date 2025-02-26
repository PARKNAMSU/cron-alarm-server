package user_repository

import (
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/entity/user_entity"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/database"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/query_tool"
)

type UserRepositoryImpl interface {
	GetUser(input GetUserInput) *GetUserOutput
	CreateUser(input CreateUserInput) (int, error)
	SetUserLoginData(input SetUserLoginDataInput) error
	SetUserOauth(input SetUserOauthInput) error
	SetUserInformation(input SetUserInformationInput) error
	Authorization(input AuthorizationInput) error
	SetUserRefreshToken(input SetUserRefreshTokenInput) error
	DeleteUser(input DeleteUserInput) error
	GetUserApiKey(userId int) *string
	GetRefreshToken(token string) *GetRefreshTokenInput
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
	var data []user_entity.UserDataEntity

	where := map[string]any{}

	if input.SelectKeyType == GET_USER_KEY_EMAIL {
		where["ui.email"] = input.Email
	} else if input.SelectKeyType == GET_USER_KEY_ID {
		where["u.id"] = input.UserId
	} else {
		return nil
	}

	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["user"],
		Action: query_tool.SELECT,
		Select: []string{"ui.user_id"},
		As:     "u",
		Join: []query_tool.JoinParams{
			{
				Table: database.MYSQL_TABLE["user_information"],
				As:    "ui",
				On:    "u.id = ui.user_id",
				Type:  "INNER",
			},
			{
				Table: database.MYSQL_TABLE["user_oauth"],
				As:    "uo",
				On:    "u.id = ui.user_id",
				Type:  "LEFT",
			},
			{
				Table: database.MYSQL_TABLE["user_login_data"],
				As:    "ul",
				On:    "u.id = ui.user_id",
				Type:  "LEFT",
			},
		},
		Where: where,
	})

	r.slaveDB.QuerySelect(
		&data,
		query,
		params...,
	)

	if len(data) == 0 {
		return nil
	}

	user := data[0]

	return &GetUserOutput{
		UserId:    user.UserId,
		Method:    user.Method,
		Status:    user.Status,
		IpAddr:    user.IpAddr,
		Email:     user.Email,
		Password:  user.Password,
		Name:      user.Name,
		OauthId:   user.OauthId,
		OauthHost: user.OauthHost,
		Auth:      user.Auth,
		AuthType:  user.AuthType,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (r *userRepository) CreateUser(input CreateUserInput) (int, error) {
	// todo: 유저데이터 생성 구현
	return 0, nil
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

func (r *userRepository) GetRefreshToken(token string) *GetRefreshTokenInput {
	// todo: refresh token 가져오기
	return nil
}

func (r *userRepository) DeleteUser(input DeleteUserInput) error {
	// todo: 유저 delete 세팅
	return nil
}

func (r *userRepository) GetUserApiKey(userId int) *string {
	// todo: 유저 api key 가져오기
	var key *string
	return key
}
