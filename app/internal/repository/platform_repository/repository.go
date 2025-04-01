package platform_repository

import (
	"errors"
	"time"

	"nspark-cron-alarm.com/cron-alarm-server/app/internal/entity/user_entity"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/root_repository"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/database"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/common_tool"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/query_tool"
)

type PlatformRepositoryImpl interface {
	root_repository.RootRepositoryImpl
	InserPlatform(input InserPlatformInput) error
	UpdatePlatform(input UpdatePlatformInput) error
	GetPlatform(input GetPlatformInput) []GetPlatformOutput
}

type platformRepository struct {
	root_repository.RootRepository
}

var (
	repo *platformRepository
)

func NewRepository(masterDB *database.CustomDB, slaveDB *database.CustomDB) PlatformRepositoryImpl {
	if repo == nil {
		repo = &platformRepository{}
		repo.SetMasterDB(masterDB)
		repo.SetSlaveDB(slaveDB)
	}
	return repo
}

func (r *platformRepository) InserPlatform(input InserPlatformInput) error {
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

func (r *platformRepository) UpdatePlatform(input UpdatePlatformInput) error {
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

func (r *platformRepository) GetPlatform(input GetPlatformInput) []GetPlatformOutput {
	list := make([]user_entity.PlatformEntity, 0)

	where := map[string]any{}

	if input.IsGetUsable {
		where["expired_at"] = query_tool.CompareColumn{
			CompareType: query_tool.GREATER,
			Value:       time.Now().Format("2006-01-02 15:04:05"),
		}
		where["status"] = 1
	}
	if input.SearchType == GET_PLATFORM_USER_ID {
		where["user_id"] = *input.UserId
	} else if input.SearchType == GET_PLATFORM_API_KEY {
		where["api_key"] = *input.ApiKey
	} else if input.SearchType == GET_PLATFORM_HOST {
		where["hostname"] = *input.Hostname
	} else {
		return []GetPlatformOutput{}
	}

	query, params := query_tool.QueryBuilder(query_tool.QueryParams{
		Table:  database.MYSQL_TABLE["userApiKey"],
		Action: query_tool.SELECT,
		Where:  where,
	})

	r.GetSlaveDB().QuerySelect(list, query, params...)

	if len(list) == 0 {
		return []GetPlatformOutput{}
	}
	return common_tool.ArrayMap(list, func(data user_entity.PlatformEntity) GetPlatformOutput {
		return GetPlatformOutput{
			UserId:       data.UserId,
			ApiKey:       data.ApiKey,
			Hostname:     data.Hostname,
			Status:       data.Status,
			PlatformName: data.PlatformName,
		}
	})
}
