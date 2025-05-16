package subscription

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/YuraSahanovskyi/weather-api/internal/db"
	"github.com/YuraSahanovskyi/weather-api/internal/domain"
	"github.com/YuraSahanovskyi/weather-api/internal/service/email"
	"gorm.io/gorm"
)

type SubscriptionConflictError struct {
}

func (e SubscriptionConflictError) Error() string {
	return "subscription for this email already exists"
}

func AddSubscription(subscription domain.Subscription) error {
	token, err := generateToken()
	if err != nil {
		return err
	}
	subscriptionEntity := db.Subscription{
		Email:     subscription.Email,
		City:      subscription.City,
		Frequency: subscription.Frequency.String(),
		Confirmed: false,
		Token: token,
	}
	if err := db.DB.Create(&subscriptionEntity).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return SubscriptionConflictError{}
		}
		return err
	}
	if err := email.SendConfirmEmail(subscription.Email, token); err != nil {
		log.Println("failed to send email: ", err)
		return err
	}
	return nil
}

func generateToken() (string, error) {
	const n int = 32
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("unable to generate token: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}
