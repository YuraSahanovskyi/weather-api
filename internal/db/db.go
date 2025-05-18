package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	err = DB.AutoMigrate(&Subscription{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	// Partial unique index to enforce unique email only for active subscriptions
	err = DB.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_active_email
		ON subscriptions (email)
		WHERE deleted_at IS NULL
	`).Error
	if err != nil {
		log.Fatalf("Failed to create partial index: %v", err)
	}

	log.Println("Connected to DB")
}
