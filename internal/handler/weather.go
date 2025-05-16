package handler

import (
	"errors"
	"net/http"

	"github.com/YuraSahanovskyi/weather-api/internal/service/weather"
	"github.com/gin-gonic/gin"
)

func GetWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	data, err := weather.GetWeather(city)
	if err != nil {
		if errors.As(err, &weather.CityNotFoundError{}) {
			_ = c.AbortWithError(http.StatusNotFound, err)
		} else {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, data)
}
