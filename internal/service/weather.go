package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/YuraSahanovskyi/weather-api/internal/model"
)

type weatherAPIResponse struct {
	Current struct {
		TempC     float64 `json:"temp_c"`
		Humidity  int     `json:"humidity"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

type errorApiResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type CityNotFound struct {
	city string
}

func (c CityNotFound) Error() string {
	return fmt.Sprintf("City %s not found", c.city)
}

func GetWeather(city string) (*model.Weather, error) {
	key := os.Getenv("API_KEY")
	if key == "" {
		panic("API_KEY env variable not found")
	}
	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", key, city))
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, handleApiError(city, body)
	}
	var weather weatherAPIResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &model.Weather{
		Temperature: weather.Current.TempC,
		Humidity:    weather.Current.Humidity,
		Description: weather.Current.Condition.Text,
	}, nil
}

func handleApiError(city string, body []byte) error {
	var res errorApiResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}
	log.Println(res)
	switch res.Error.Code {
	case 1006:
		return CityNotFound{city: city}
	default:
		return fmt.Errorf("bad request")
	}

}
