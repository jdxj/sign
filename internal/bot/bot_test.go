package bot

import (
	"fmt"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TestBot(t *testing.T) {
	client, err := tgbotapi.NewBotAPI("")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	updateConfig := tgbotapi.NewUpdate(0)
	updates, err := client.GetUpdates(updateConfig)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	for _, update := range updates {
		if update.Message == nil {
			continue
		}

		fmt.Printf("%#v\n", *update.Message.Chat)
	}
}

var (
	token       = ""
	chat  int64 = 0
)

func TestSend(t *testing.T) {
	Init(token, chat)

	Send("test")
}
