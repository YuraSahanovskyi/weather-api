package main

import (
	"log"

	"github.com/YuraSahanovskyi/weather-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	handler.RegisterRoutes(r)

	log.Fatalln(r.Run(":8080"))
}
