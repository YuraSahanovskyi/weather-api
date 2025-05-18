package config

import (
	"log"
	"os"
)

type config struct {
	appHost string
	appPort string
}

var appConfig config

func Init() {
	appHost := os.Getenv("HOST")
	appPort := os.Getenv("PORT")
	if appHost == "" || appPort == "" {
		log.Fatal("app host or app port are not found")
	}
	appConfig = config{
		appHost: appHost,
		appPort: appPort,
	}
}

func GetAppHost() string {
	return appConfig.appHost
}

func GetAppPort() string {
	return appConfig.appPort
}
