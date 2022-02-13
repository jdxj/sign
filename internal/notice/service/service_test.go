package service

import (
	"fmt"
	"testing"

	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TestSendNotice(t *testing.T) {
	client, err := bot.NewBotAPI("")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	mc := bot.NewMessage(548250375, "abc")
	m, err := client.Send(mc)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("m: %+v\n", m)
}
