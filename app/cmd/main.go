package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"nspark-cron-alarm.com/cron-alarm-server/app/internal"
)

func main() {
	internal.GetApp()
}
