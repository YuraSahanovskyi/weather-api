package main

import (
	"log"

	"github.com/YuraSahanovskyi/weather-api/internal/config"
	"github.com/YuraSahanovskyi/weather-api/internal/db"
	"github.com/YuraSahanovskyi/weather-api/internal/handler"
	"github.com/YuraSahanovskyi/weather-api/internal/service/cron"
	"github.com/YuraSahanovskyi/weather-api/internal/service/email"
	"github.com/YuraSahanovskyi/weather-api/internal/service/weather"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	db.Init()
	email.InitSMTP()
	weather.ReadApiKey()
	
	go cron.StartSceduler()

	r := gin.Default()

	handler.RegisterRoutes(r)

	log.Fatalln(r.Run(":" + config.GetAppPort()))
}
