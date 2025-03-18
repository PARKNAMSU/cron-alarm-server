package user_repository

import (
	"errors"
	"time"

	"nspark-cron-alarm.com/cron-alarm-server/app/internal/entity/user_entity"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/root_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/database"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/common_tool"
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
	GetUserPlatform(input GetUserPlatformInput) []GetUserPlatformOutput
	GetRefreshToken(token string) *GetRefreshTokenInput
	InsertUserPlatform(input InsertUserPlatformInput) error
	UpdateUserPlatform(input UpdateUserPlatformInput) error
	SetUserAuthCode(input SetAuthCodeInput) error
	UserAuthorization(userId int, authType string) error
	GetAvailableAuthCode(userId int, action string) *GetAvailableAuthCodeOutput
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
		Select: []string{
			"ui.user_id",
			"ui.email",
			"ul.password",
			"u.method",
			"u.status",
			"u.ip_addr",
			"ui.name",
			"uo.oauth_id",
			"uo.oauth_host",
			"ui.auth",
			"ui.auth_type",
			"u.created_at",
			"u.updated_at",
			"up.grade",
			"up.max_platform_cnt",
		},
		As: "u",
		Join: []query_tool.JoinParams{
			{
				Table: database.MYSQL_TABLE["userInformation"],
				As:    "ui",
				On:    "u.id = ui.user_id",
				Type:  "INNER",
			},
			{
				Table: database.MYSQL_TABLE["userPermission"],
				As:    "up",
				On:    "ui.permission_id = up.id",
				Type:  "LEFT",
			},
			{
				Table: database.MYSQL_TABLE["userOauth"],
				As:    "uo",
				On:    "u.id = ui.user_id",
				Type:  "LEFT",
			},
			{
				Table: database.MYSQL_TABLE["userLoginData"],
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

	output := &GetUserOutput{
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

	if user.Grade != nil {
		output.Grade = user.Grade
		output.MaxPlatformCnt = *user.MaxPlatformCnt
	}

	return output
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

func (r *userRepository) GetUserPlatform(input GetUserPlatformInput) []GetUserPlatformOutput {
	list := make([]user_entity.UserPlatformEntity, 0)

	where := map[string]any{}

	if input.IsGetUsable {
		where["expired_at"] = query_tool.CompareColumn{
			CompareType: query_tool.GREATER,
			Value:       time.Now().Format("2006-01-02 15:04:05"),
		}
		where["status"] = 1
	}
	if input.SearchType == GET_USER_API_KEY_USER_ID {
		where["user_id"] = *input.UserId
	} else if input.SearchType == GET_USER_API_KEY_API_KEY {
		where["api_key"] = *input.ApiKey
	} else if input.SearchType == GET_USER_API_KEY_HOST {
		where["hostname"] = *input.Hostname
	} else {
		return []GetUserPlatformOutput{}
	}

	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["userApiKey"],
		Action: query_tool.SELECT,
		Where:  where,
	})

	r.GetSlaveDB().QuerySelect(list, query, params...)

	if len(list) == 0 {
		return []GetUserPlatformOutput{}
	}
	return common_tool.ArrayMap(list, func(data user_entity.UserPlatformEntity) GetUserPlatformOutput {
		return GetUserPlatformOutput{
			UserId:       data.UserId,
			ApiKey:       data.ApiKey,
			Hostname:     data.Hostname,
			Status:       data.Status,
			PlatformName: data.PlatformName,
		}
	})
}

func (r *userRepository) InsertUserPlatform(input InsertUserPlatformInput) error {
	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["userPlatform"],
		Action: query_tool.INSERT,
		Set: map[string]any{
			"user_id":    input.UserId,
			"api_key":    input.ApiKey,
			"expired_at": input.ExpiredAt,
			"status":     1,
			"hostname":   input.Hostname,
		},
	})

	_, err := r.GetMasterDB().QueryExecute(query, params...)

	return err
}

func (r *userRepository) UpdateUserPlatform(input UpdateUserPlatformInput) error {
	set := map[string]any{}

	if input.PlatformName != nil {
		set["platform_name"] = *input.PlatformName
	}
	if input.ExpiredAt != nil {
		set["expired_at"] = *input.ExpiredAt
	}

	if len(set) == 0 {
		return errors.New("no update data")
	}

	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["userPlatform"],
		Action: query_tool.INSERT,
		Set:    set,
		Where: map[string]any{
			"hostname": input.Hostname,
			"user_id":  input.UserId,
		},
	})

	_, err := r.GetMasterDB().QueryExecute(query, params...)

	return err
}

func (r *userRepository) SetUserAuthCode(input SetAuthCodeInput) error {
	var set map[string]any
	var duplicate map[string]any

	if input.Action == "auth" {
		set = map[string]any{
			"user_id":         input.UserId,
			"receive_account": input.ReceiveAccount,
			"action":          input.Action,
			"status":          input.Status,
		}
		duplicate = map[string]any{
			"status": input.Status,
		}
		if input.Status == 0 {
			set["code"] = input.Code
			set["auth_type"] = input.AuthType
			set["action"] = input.Action
			set["expired_at"] = input.ExpiredAt
			duplicate["code"] = input.Code
			duplicate["auth_type"] = input.AuthType
			duplicate["status"] = input.Status
			duplicate["expired_at"] = input.ExpiredAt
		}
	} else if input.Action == "password" {
		set = map[string]any{
			"user_id":         input.UserId,
			"receive_account": input.ReceiveAccount,
			"auth_type":       input.AuthType,
			"code":            input.Code,
			"action":          input.Action,
			"expired_at":      input.ExpiredAt,
		}
		duplicate = map[string]any{
			"code":       input.Code,
			"expired_at": input.ExpiredAt,
			"auth_type":  input.AuthType,
		}
	}

	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["userAuthCode"],
		Action: query_tool.DUPLICATE,
		Set:    map[string]any{},
	})

	_, err := r.GetMasterDB().QueryExecute(query, params...)
	return err
}

func (r *userRepository) UserAuthorization(userId int, authType string) error {
	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["userInformation"],
		Action: query_tool.UPDATE,
		Set: map[string]any{
			"auth":      1,
			"auth_type": authType,
		},
		Where: map[string]any{
			"user_id": userId,
		},
	})
	_, err := r.GetMasterDB().QueryExecute(query, params...)
	return err
}

func (r *userRepository) GetAvailableAuthCode(userId int, action string) *GetAvailableAuthCodeOutput {
	list := []user_entity.UserAuthCodeEntity{}

	where := map[string]any{
		"user_id": userId,
		"action":  action,
		"expired_at": query_tool.CompareColumn{
			CompareType: query_tool.GREATER,
			Value:       time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	if action == "auth" {
		where["status"] = 0
	}

	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["userAuthCode"],
		Action: query_tool.SELECT,
		Where:  where,
	})

	r.GetSlaveDB().QuerySelect(&list, query, params...)

	if len(list) == 0 {
		return nil
	}

	data := list[0]

	return &GetAvailableAuthCodeOutput{
		UserId:         data.UserId,
		ReceiveAccount: data.ReceiveAccount,
		Action:         data.Action,
		Status:         data.Status,
		Code:           data.Code,
		AuthType:       data.AuthType,
	}
}
