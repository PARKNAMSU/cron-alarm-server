package mail_tool

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

func SendMail(to string, msg string) error {
	fmt.Println(os.Getenv("MAIL_EMAIL"))
	fmt.Println(os.Getenv("MAIL_PASSWORD"))
	fromMsg := "From: " + os.Getenv("MAIL_EMAIL") + "\r"
	toMsg := "To: " + to + "\r"
	body := msg + "\r"
	msgByte := []byte(strings.Join([]string{fromMsg, toMsg, body}, "\n"))

	err := smtp.SendMail(
		"smtp.naver.com:587",
		smtp.PlainAuth(
			"",
			os.Getenv("MAIL_EMAIL"),
			os.Getenv("MAIL_PASSWORD"),
			"smtp.naver.com",
		),
		os.Getenv("MAIL_EMAIL"),
		[]string{to},
		msgByte,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
