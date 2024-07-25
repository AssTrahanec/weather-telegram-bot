package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	config "github.com/asstrahanec/weather-telegram-bot/configs"
	"net/http"
	"time"
)

type WeatherInfo struct {
	Temp      float64 `json:"temp"`
	Condition string  `json:"condition"`
	Wind      float64 `json:"wind"`
	Humidity  int     `json:"humidity"`
}

func GetWeather(city string) (*WeatherInfo, error) {
	apiKey := config.WeatherAPIToken
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=ru", city, apiKey)

	return fetchWeather(url)
}

func GetWeatherByCoords(lat, lon float64) (*WeatherInfo, error) {
	apiKey := config.WeatherAPIToken
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%.4f&lon=%.4f&appid=%s&units=metric&lang=ru", lat, lon, apiKey)
	fmt.Println(url)
	return fetchWeather(url)
}

func fetchWeather(url string) (*WeatherInfo, error) {
	client := &http.Client{
		Timeout: 6 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 || resp.StatusCode == 400 {
		return nil, errors.New("invalid input")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var result struct {
		Main struct {
			Temp     float64 `json:"temp"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
		Wind struct {
			Speed float64 `json:"speed"`
		} `json:"wind"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	weatherInfo := &WeatherInfo{
		Temp:      result.Main.Temp,
		Condition: result.Weather[0].Description,
		Humidity:  result.Main.Humidity,
		Wind:      result.Wind.Speed,
	}

	return weatherInfo, nil
}
