package handler

import (
	"errors"
	"net/http"

	"github.com/YuraSahanovskyi/weather-api/internal/service/subscription"
	"github.com/gin-gonic/gin"
)

func Unsubscribe(c *gin.Context) {
	token := c.Param("token")
	if err := subscription.Unsubscribe(token); err != nil {
		switch {
		case errors.As(err, &subscription.InvalidTokenError{}):
			respondError(c, http.StatusBadRequest, "token is invalid")
		case errors.As(err, &subscription.TokenNotFoundError{}):
			respondError(c, http.StatusNotFound, "token not found")
		default:
			respondError(c, http.StatusInternalServerError, "internal server error")
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "unsubscribed successfuly"})
}
