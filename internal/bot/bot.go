package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/jdxj/sign/internal/logger"
)

var (
	client *tgbotapi.BotAPI
	chatID int64
)

func Init(token string, chat int64) {
	var err error
	client, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalln(err)
	}
	chatID = chat
}

func Send(text string) {
	mc := tgbotapi.NewMessage(chatID, text)
	// todo: 并发执行?
	_, err := client.Send(mc)
	if err != nil {
		logger.Errorf("Send err: %s, text: %s", err, text)
	}
}
