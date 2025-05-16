package subscription

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/YuraSahanovskyi/weather-api/internal/db"
	"gorm.io/gorm"
)

type TokenNotFoundError struct {
	token string
}

func (e TokenNotFoundError) Error() string {
	return fmt.Sprintf("token %s not found", e.token)
}

type AlreadyConfirmedError struct {
	token string
}

func (e AlreadyConfirmedError) Error() string {
	return fmt.Sprintf("token %s is already confirmed", e.token)
}

type InvalidTokenError struct {
	token string
}

func (e InvalidTokenError) Error() string {
	return fmt.Sprintf("token %s is invalid", e.token)
}




func ConfirmSubscription(token string) error {
	if !isValidHexToken(token) {
		return InvalidTokenError{token: token}
	}

	var subscription db.Subscription
	err := db.DB.Where("token = ?", token).First(&subscription).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return TokenNotFoundError{token: token}
	}
	if err != nil {
		return fmt.Errorf("failed to query subscription: %w", err)
	}

	if subscription.Confirmed {
		return AlreadyConfirmedError{token: token}
	}

	subscription.Confirmed = true
	if err := db.DB.Save(&subscription).Error; err != nil {
		return fmt.Errorf("failed to confirm subscription: %w", err)
	}
	return nil
}

var hexTokenRegex = regexp.MustCompile(`^[a-f0-9]{64}$`)

func isValidHexToken(token string) bool {
	return hexTokenRegex.MatchString(token)
}