package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	errInvalidCity = errors.New("invalid city")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, "Failed to fetch the weather. Try again later.")
	switch err {
	case errInvalidCity:
		msg.Text = "Неверно введен город!"
	default:
		msg.Text = "Failed to fetch the weather. Try again later."
	}
	b.bot.Send(msg)

}
