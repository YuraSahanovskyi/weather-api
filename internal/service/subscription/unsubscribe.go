package subscription

import (
	"errors"
	"fmt"

	"github.com/YuraSahanovskyi/weather-api/internal/db"
	"gorm.io/gorm"
)

func Unsubscribe(token string) error {
	if !isValidHexToken(token) {
		return InvalidTokenError{token: token}
	}

	var sub db.Subscription
	if err := db.DB.Where("token = ?", token).First(&sub).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return TokenNotFoundError{token: token}
		}
		return err
	}

	if err := db.DB.Delete(&sub).Error; err != nil {
		return fmt.Errorf("failed to delete subscription: %w", err)
	}

	return nil
}
