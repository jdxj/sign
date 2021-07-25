package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

func Send(text string) error {
	mc := tgbotapi.NewMessage(chatID, text)
	_, err := client.Send(mc)
	return err
}
