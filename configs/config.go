// pkg/config/config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	WeatherAPIToken  string
	TelegramAPIToken string
)

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	WeatherAPIToken = os.Getenv("WEATHER_API_TOKEN")
	if WeatherAPIToken == "" {
		log.Fatal("WEATHER_API_TOKEN not set in .env file")
	}
	TelegramAPIToken = os.Getenv("TELEGRAM_API_TOKEN")
	if TelegramAPIToken == "" {
		log.Fatal("TELEGRAM_API_TOKEN not set in .env file")
	}
}
