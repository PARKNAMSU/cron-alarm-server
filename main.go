package main

import (
	"nspark-cron-alarm.com/cron-alarm-server/src"
	"nspark-cron-alarm.com/cron-alarm-server/src/repository"
)

func main() {
	repository.RepositoryLoad()
	src.GetApp()
}
