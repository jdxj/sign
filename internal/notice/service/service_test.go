package service

import (
	"fmt"
	"os"
	"testing"

	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/jdxj/sign/configs"
	"github.com/jdxj/sign/internal/pkg/logger"
)

func TestMain(t *testing.M) {
	logger.Init("./notice.log")
	os.Exit(t.Run())
}

func TestSendNotice(t *testing.T) {
	client, err := bot.NewBotAPI(configs.BotKey)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	mc := bot.NewMessage(configs.MyTelID, "abc")
	m, err := client.Send(mc)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("m: %+v\n", m)
}
