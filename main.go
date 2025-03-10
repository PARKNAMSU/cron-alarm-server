package main

import (
	"fmt"

	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/encrypt_tool"
)

func main() {
	data := "aaaa"

	s := "aaaaaaa"

	a, _ := encrypt_tool.Encrypt([]byte(data), s)

	b, _ := encrypt_tool.Encrypt([]byte(data), s)

	ad, _ := encrypt_tool.Decrypt(a, s)

	bd, _ := encrypt_tool.Decrypt(b, s)

	fmt.Println(string(ad))
	fmt.Println(string(bd))

}
