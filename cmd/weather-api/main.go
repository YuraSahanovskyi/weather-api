package main

import (
	"log"
	"time"

	"github.com/YuraSahanovskyi/weather-api/internal/config"
	"github.com/YuraSahanovskyi/weather-api/internal/db"
	"github.com/YuraSahanovskyi/weather-api/internal/handler"
	"github.com/YuraSahanovskyi/weather-api/internal/service/cron"
	"github.com/YuraSahanovskyi/weather-api/internal/service/email"
	"github.com/YuraSahanovskyi/weather-api/internal/service/weather"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	db.Init()
	email.InitSMTP()
	weather.ReadApiKey()

	go cron.StartSceduler()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handler.RegisterRoutes(r)

	log.Fatalln(r.Run("0.0.0.0:" + config.GetAppPort()))
}
