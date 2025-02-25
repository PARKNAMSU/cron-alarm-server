package main

import (
	"nspark-cron-alarm.com/cron-alarm-server/app/internal"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal/repository"
)

func main() {
	repository.RepositoryLoad()
	internal.GetApp()
}
