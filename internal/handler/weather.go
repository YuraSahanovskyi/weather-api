package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/YuraSahanovskyi/weather-api/internal/service"
	"github.com/gin-gonic/gin"
)

func GetWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city is required"})
		return
	}
	weather, err := service.GetWeather(city)
	if err != nil {
		log.Println("[ERROR]", err)
		if errors.As(err, &service.CityNotFound{}) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		_ = c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, weather)
}
