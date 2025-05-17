package handler

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/weather", GetWeather)
		api.POST("/subscribe", Subscribe)
		api.GET("/confirm/:token", ConfirmSubscription)
		api.GET("/unsubscribe/:token", Unsubscribe)
	}
}
