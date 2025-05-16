package cron

import (
	"log"

	"github.com/YuraSahanovskyi/weather-api/internal/domain"
	"github.com/YuraSahanovskyi/weather-api/internal/service/subscription"
	"github.com/robfig/cron/v3"
)

func StartSceduler() {
	c := cron.New()
	if _, err := c.AddFunc("@hourly", func() {
		log.Println("running hourly job")
		subscription.ProcessSubscribers(domain.Hourly)
	}); err != nil {
		log.Fatal(err)
	}

	// daily emails are sent at 9:00 AM
	if _, err := c.AddFunc("0 9 * * *", func() {
		log.Println("running daily job")
		subscription.ProcessSubscribers(domain.Daily)
	}); err != nil {
		log.Fatal(err)
	}

	c.Start()
}
