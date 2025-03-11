package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/mail_tool"
)

func main() {
	mail_tool.SendMail("skatn7979@gmail.com", "test123")
}
