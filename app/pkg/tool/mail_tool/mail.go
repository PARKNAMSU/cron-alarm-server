package mail_tool

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

// Base64 인코딩 함수
func encodeBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func readTemplate(title string, body string, path string) string {
	root, _ := os.Getwd()
	template, err := os.ReadFile(root + path)
	if err != nil {
		log.Println(path)
		return ""
	}

	templateStr := strings.Replace(string(template), "${body}", body, -1)
	templateStr = strings.Replace(templateStr, "${title}", title, -1)

	return encodeBase64(templateStr)
}

func SendAlarmMail(to string, msg string, title string) error {
	headers := fmt.Sprintf(
		`From: %s
To: %s
Subject: =?UTF-8?B?%s?=
Content-Type: text/html; charset="UTF-8"
Content-Transfer-Encoding: base64`+"\n",
		os.Getenv("MAIL_EMAIL"),
		to,
		encodeBase64(title),
	)

	body := readTemplate(title, msg, "/template/alarm_email.html") + "\r"
	msgByte := []byte(strings.Join([]string{headers, body}, "\n"))

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
