package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/YuraSahanovskyi/weather-api/internal/domain"
	"github.com/YuraSahanovskyi/weather-api/internal/dto"
	"github.com/YuraSahanovskyi/weather-api/internal/service/subscription"
	"github.com/gin-gonic/gin"
)

func Subscribe(c *gin.Context) {
	input, err := parseInput(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	frequency, err := domain.ParseFrequency(input.Frequency)
	if err != nil {
		respondError(c, http.StatusBadRequest, "frequency is invalid")
		return
	}

	if err = subscription.AddSubscription(domain.Subscription{
		Email:     input.Email,
		City:      input.City,
		Frequency: *frequency,
	}); err != nil {
		switch {
		case errors.As(err, &subscription.SubscriptionConflictError{}):
			respondError(c, http.StatusConflict, "subscription for this email already exists")
		default:
			respondError(c, http.StatusInternalServerError, "internal server error")
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "confirm subscription with link in email"})
}

func parseInput(c *gin.Context) (*dto.SubscriptionDto, error) {
	var input dto.SubscriptionDto
	contentType := c.ContentType()

	var err error
	switch {
	case strings.HasPrefix(contentType, "application/json"):
		err = c.ShouldBindJSON(&input)
	case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
		err = c.ShouldBind(&input)
	default:
		return nil, fmt.Errorf("unsupported content type: %s", contentType)
	}

	if err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}

	return &input, nil
}
