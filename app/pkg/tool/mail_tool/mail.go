package mail_tool

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendMail(to string, msg string) error {
	fmt.Println(os.Getenv("MAIL_EMAIL"))
	fmt.Println(os.Getenv("MAIL_PASSWORD"))
	headerSubject := "Subject: 테스트\r\n"
	headerBlank := "\r\n"
	body := "메일 테스트입니다\r\n"
	msgByte := []byte(headerSubject + headerBlank + body)

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
