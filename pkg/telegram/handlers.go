package telegram

import (
	"fmt"
	"github.com/asstrahanec/weather-telegram-bot/pkg/weather"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	//newMessage := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	weatherInfo, err := weather.GetWeather(message.Text)
	if err != nil {
		log.Printf("Failed to get weather: %v", err)
		return err
	}
	response := formatWeatherResponse(weatherInfo)
	return b.sendMessage(message.Chat.ID, response)
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case "start":
		return b.handleStartCommand(message)
	case "weather":
		return b.handleWeatherCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}

}
func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	return b.sendMessage(message.Chat.ID, "Напишите /weather, чтобы узнать погоду.")
}

func (b *Bot) handleWeatherCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Поделитесь геолокацией, чтобы узнать погоду или напишите название города")
	button := tgbotapi.NewKeyboardButton("Поделиться геолокацией")
	button.RequestLocation = true
	keyboard := tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{button})
	msg.ReplyMarkup = keyboard
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleLocation(message *tgbotapi.Message) error {
	lat := message.Location.Latitude
	lon := message.Location.Longitude
	fmt.Println(lat, lon)
	weatherInfo, err := weather.GetWeatherByCoords(lat, lon)
	if err != nil {
		log.Printf("Failed to get weather: %v", err)
		return err
	}

	response := formatWeatherResponse(weatherInfo)
	return b.sendMessage(message.Chat.ID, response)
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	return b.sendMessage(message.Chat.ID, "Я не знаю такой команды.")
}
func (b *Bot) sendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := b.bot.Send(msg)
	return err
}
func formatWeatherResponse(weatherInfo *weather.WeatherInfo) string {
	return fmt.Sprintf("Погода в вашем городе:\nТемпература: %.2f°C\nПогодные условия: %s\nВлажность: %d%%\nВетер: %.2f м/с\n",
		weatherInfo.Temp, weatherInfo.Condition, weatherInfo.Humidity, weatherInfo.Wind)
}
