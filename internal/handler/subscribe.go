package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/YuraSahanovskyi/weather-api/internal/dto"
	"github.com/YuraSahanovskyi/weather-api/internal/domain"
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
		_ = c.AbortWithError(http.StatusBadRequest, fmt.Errorf("can`t parse frequency: %w", err))
		return
	}

	if err = subscription.AddSubscription(domain.Subscription{
		Email:     input.Email,
		City:      input.City,
		Frequency: *frequency,
	}); err != nil {
		switch{
		case errors.As(err, &subscription.SubscriptionConflictError{}):
			_ = c.AbortWithError(http.StatusConflict, err)
		default:
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	c.Status(http.StatusOK)
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
