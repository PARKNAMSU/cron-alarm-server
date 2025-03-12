package user_repository

import (
	"time"

	"nspark-cron-alarm.com/cron-alarm-server/app/internal/entity/user_entity"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/root_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/database"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/query_tool"
)

type UserRepositoryImpl interface {
	root_repository.RootRepositoryImpl
	GetUser(input GetUserInput) *GetUserOutput
	CreateUser(input CreateUserInput) (int, error)
	SetUserLoginData(input SetUserLoginDataInput) error
	SetUserOauth(input SetUserOauthInput) error
	SetUserInformation(input SetUserInformationInput) error
	Authorization(input AuthorizationInput) error
	SetUserRefreshToken(input SetUserRefreshTokenInput) error
	DeleteUser(input DeleteUserInput) error
	GetUserApiKey(input GetUserApiKeyInput) *GetUserApiKeyOutput
	GetRefreshToken(token string) *GetRefreshTokenInput
}

type userRepository struct {
	root_repository.RootRepository
}

var (
	repo *userRepository
)

func NewRepository(masterDB *database.CustomDB, slaveDB *database.CustomDB) UserRepositoryImpl {
	if repo == nil {
		repo = &userRepository{}
		repo.SetMasterDB(masterDB)
		repo.SetSlaveDB(slaveDB)
	}
	return repo
}

func (r *userRepository) GetUser(input GetUserInput) *GetUserOutput {
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

	r.GetSlaveDB().QuerySelect(
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
	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["user"],
		Action: query_tool.INSERT,
		Set: map[string]any{
			"ip_addr": input.IpAddr,
			"method":  "normal",
			"status":  1,
		},
	})

	result, err := r.GetMasterDB().QueryExecute(query, params...)

	if err != nil {
		return 0, nil
	}

	id, _ := result.LastInsertId()

	return int(id), err
}

func (r *userRepository) SetUserLoginData(input SetUserLoginDataInput) error {
	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["user"],
		Action: query_tool.DUPLICATE,
		Set: map[string]any{
			"email":    input.Email,
			"password": input.Password,
			"user_id":  input.UserId,
		},
		Duplicate: map[string]any{
			"password": input.Password,
		},
	})
	_, err := r.GetMasterDB().QueryExecute(query, params...)
	return err
}

func (r *userRepository) SetUserOauth(input SetUserOauthInput) error {
	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["userOauth"],
		Action: query_tool.DUPLICATE,
		Set: map[string]any{
			"user_id":    input.UserId,
			"oauth_id":   input.OauthId,
			"oauth_host": input.OauthHost,
		},
		Duplicate: map[string]any{
			"oauth_id":   input.OauthId,
			"oauth_host": input.OauthHost,
		},
	})
	_, err := r.GetMasterDB().QueryExecute(query, params...)
	return err
}

func (r *userRepository) SetUserInformation(input SetUserInformationInput) error {
	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["userInformation"],
		Action: query_tool.DUPLICATE,
		Set: map[string]any{
			"user_id": input.UserId,
			"email":   input.Email,
			"name":    input.Name,
		},
		Duplicate: map[string]any{
			"name": input.Name,
		},
	})
	_, err := r.GetMasterDB().QueryExecute(query, params...)
	return err
}

func (r *userRepository) Authorization(input AuthorizationInput) error {
	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["userInformation"],
		Action: query_tool.UPDATE,
		Set: map[string]any{
			"auth_type": input.AuthType,
			"auth":      1,
		},
		Where: map[string]any{
			"user_id": input.UserId,
		},
	})
	_, err := r.GetMasterDB().QueryExecute(query, params...)
	return err
}

func (r *userRepository) SetUserRefreshToken(input SetUserRefreshTokenInput) error {
	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["userRefreshToken"],
		Action: query_tool.DUPLICATE,
		Set: map[string]any{
			"user_id":    input.UserId,
			"token":      input.Token,
			"ip_addr":    input.IpAddr,
			"expired_at": input.ExpiredAt,
		},
		Duplicate: map[string]any{
			"token":      input.Token,
			"ip_addr":    input.IpAddr,
			"expired_at": input.ExpiredAt,
		},
	})
	_, err := r.GetMasterDB().QueryExecute(query, params...)
	return err
}

func (r *userRepository) GetRefreshToken(token string) *GetRefreshTokenInput {
	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["userRefreshToken"],
		Action: query_tool.SELECT,
		Where: map[string]any{
			"token": token,
			"expired_at": query_tool.CompareColumn{
				CompareType: query_tool.GREATER,
				Value:       time.Now().Format("2006-01-02 15:04:05"),
			},
		},
	})

	var data *user_entity.UserRefreshTokenEntity
	r.GetSlaveDB().QuerySelect(data, query, params...)

	if data == nil {
		return nil
	}
	return &GetRefreshTokenInput{
		Token:  data.Token,
		UserId: data.UserId,
		Status: data.Status,
		IpAddr: data.IpAddr,
	}
}

func (r *userRepository) DeleteUser(input DeleteUserInput) error {
	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table: database.MYSQL_TABLE["user"],
		Set:   map[string]any{"status": 0},
		Where: map[string]any{"id": input.UserId},
	})
	_, err := r.GetMasterDB().QueryExecute(query, params...)
	return err
}

func (r *userRepository) GetUserApiKey(input GetUserApiKeyInput) *GetUserApiKeyOutput {
	var key *user_entity.UserApiKeyEntity

	where := map[string]any{}

	if input.SearchType == GET_USER_API_KEY_USER_ID {
		where["user_id"] = *input.UserId
	} else if input.SearchType == GET_USER_API_KEY_API_KEY {
		where["api_key"] = *input.ApiKey
	} else {
		return nil
	}

	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["userApiKey"],
		Action: query_tool.SELECT,
		Where:  where,
	})

	r.GetSlaveDB().QuerySelect(key, query, params...)

	if key == nil {
		return nil
	}

	return &GetUserApiKeyOutput{
		UserId: key.UserId,
		ApiKey: key.ApiKey,
	}
}

func (r *userRepository) SetUserApiKey(userId int, key string, expiredAt time.Time) error {
	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["userApiKey"],
		Action: query_tool.DUPLICATE,
		Set: map[string]any{
			"user_id":    userId,
			"api_key":    key,
			"expired_at": expiredAt,
		},
		Duplicate: map[string]any{
			"api_key":    key,
			"expired_at": expiredAt,
		},
	})

	_, err := r.GetMasterDB().QueryExecute(query, params...)

	return err
}
