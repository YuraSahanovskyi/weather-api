package subscription

import (
	"log"

	"github.com/YuraSahanovskyi/weather-api/internal/db"
	"github.com/YuraSahanovskyi/weather-api/internal/domain"
	"github.com/YuraSahanovskyi/weather-api/internal/service/email"
	"github.com/YuraSahanovskyi/weather-api/internal/service/weather"
)

func ProcessSubscribers(frequency domain.Frequency) {
	var subs []db.Subscription
	if err := db.DB.Where("confirmed = ? AND frequency = ?", true, frequency.String()).Find(&subs).Error; err != nil {
		log.Println("faled to get subscribers:", err)
	}
	for _, sub := range subs {
		weather, err := weather.GetWeather(sub.City)
		if err != nil {
			log.Println("failed to get weather:", err)
		}
		if err := email.SendWeatherEmail(sub.Email, sub.City, *weather, sub.Token); err != nil {
			log.Println("failed to send email:", err)
		}
	}

}
