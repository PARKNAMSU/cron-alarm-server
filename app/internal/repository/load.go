package repository

import (
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/interface/repository_impl"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository/user_repository"
)

func RepositoryLoad() {
	loadRepos := []repository_impl.RepositoryImpl{
		user_repository.Repository,
	}
	for _, repo := range loadRepos {
		repo.InitRepository()
	}
}
