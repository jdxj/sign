package sender

import (
	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/jdxj/sign/internal/pkg/config"
)

var (
	Messengers = map[string]Messenger{}
)

type Letter struct {
	Sender    string
	Recipient string
	Subject   string
	Content   string
}

type Messenger interface {
	Send(*Letter) error
}

func Init(conf config.Bot) (err error) {
	tb, err = bot.NewBotAPI(conf.Token)
	if err != nil {
		return
	}
	Messengers[telegramMessenger] = &telegram{}
	return
}
