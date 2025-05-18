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
		respondError(c, http.StatusBadRequest, "city not specified")
		return
	}
	data, err := weather.GetWeather(city)
	if err != nil {
		if errors.As(err, &weather.CityNotFoundError{}) {
			respondError(c, http.StatusNotFound, "city not found")
		} else {
			respondError(c, http.StatusInternalServerError, "internal server error")
		}
		return
	}
	c.JSON(http.StatusOK, data)
}
