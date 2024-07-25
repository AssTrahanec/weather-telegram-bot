package main

import (
	config "github.com/asstrahanec/weather-telegram-bot/configs"
	"github.com/asstrahanec/weather-telegram-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	config.Load()
	bot, err := tgbotapi.NewBotAPI(config.TelegramAPIToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
