package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type DiscordMessageType int

const (
	MESSAGE_LOGGING DiscordMessageType = 0
)

type DiscordInput struct {
	WebhookUrl  string
	Message     string
	Title       string
	MessageType DiscordMessageType
}

type webhookPayload struct {
	Content string `json:"content"`
}

func SendDiscordMessage(input DiscordInput) error {
	content := ""

	switch input.MessageType {
	case MESSAGE_LOGGING:
		{
			content = strings.Join([]string{
				"**" + input.Title + "**\n",
				"```" + input.Message + "```",
			}, "\n")
		}
	default:
		content = strings.Join([]string{
			input.Title,
			input.Message,
		}, "\n")
	}

	payload := webhookPayload{
		Content: content,
	}

	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		return fmt.Errorf("[DISCORD]fail to generate payload: %v", err)
	}

	// HTTP 요청 생성
	resp, err := http.Post(input.WebhookUrl, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("[DISCORD]Fail to request1: %v", err)
	}

	defer resp.Body.Close()

	// 응답 확인
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("[DISCORD]Fail to request2: %s", string(body))
	}

	log.Println("[DISCORD]✅ Success to send!!")

	return nil
}
