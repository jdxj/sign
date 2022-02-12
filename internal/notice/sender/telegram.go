package sender

import (
	"errors"
	"fmt"
	"strconv"

	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	telegramMessenger = "telegram-messenger"
)

var (
	tb *bot.BotAPI
)

var (
	ErrInvalidTelegramID = errors.New("invalid telegram id")
	ErrSendFailed        = errors.New("send failed")
)

type telegram struct {
}

func (t *telegram) Send(letter *Letter) error {
	if letter == nil {
		return nil
	}

	telegramID, err := strconv.ParseInt(letter.Recipient, 10, 64)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidTelegramID, letter.Recipient)
	}

	var text string
	if letter.Subject != "" {
		text += fmt.Sprintf("[%s]\n", letter.Subject)
	}
	if letter.Content != "" {
		text += letter.Content
	}
	_, err = tb.Send(bot.NewMessage(telegramID, text))
	if err != nil {
		return fmt.Errorf("%w, err: %s", ErrSendFailed, err)
	}
	return nil
}
