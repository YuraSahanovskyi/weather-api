package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/YuraSahanovskyi/weather-api/internal/domain"
)

type weatherApiResponse struct {
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

type CityNotFoundError struct {
	city string
}

func (c CityNotFoundError) Error() string {
	return fmt.Sprintf("city %s not found", c.city)
}

var apiKey string

func ReadApiKey() {
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is not set")
	}
}

const API_URL string = "http://api.weatherapi.com/v1/current.json?key=%s&q=%s"

func GetWeather(city string) (*domain.Weather, error) {

	url := fmt.Sprintf(API_URL, apiKey, city)
	res, err := http.Get(url)

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
	var weather weatherApiResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &domain.Weather{
		Temperature: weather.Current.TempC,
		Humidity:    weather.Current.Humidity,
		Description: weather.Current.Condition.Text,
	}, nil
}

const NOT_FOUND_CODE int = 1006

func handleApiError(city string, body []byte) error {
	var res errorApiResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}
	log.Println(res)
	switch res.Error.Code {
	case NOT_FOUND_CODE:
		return CityNotFoundError{city: city}
	default:
		return fmt.Errorf("bad request")
	}

}
