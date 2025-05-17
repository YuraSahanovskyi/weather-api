package email

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/YuraSahanovskyi/weather-api/internal/config"
	"github.com/YuraSahanovskyi/weather-api/internal/domain"
)

type smtpCredentials struct {
	host     string
	port     string
	user     string
	password string
	from     string
}

var smtpCred smtpCredentials

func InitSMTP() {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	from := os.Getenv("SMTP_FROM")
	if host == "" || port == "" || user == "" || password == "" || from == "" {
		log.Fatal("smtp credentials not found")
	}
	smtpCred = smtpCredentials{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		from:     from,
	}
}

func SendConfirmEmail(toEmail string, token string) error {

	to := []string{toEmail}

	confirmLink := fmt.Sprintf("%s:%s/api/confirm/%s", config.GetAppHost(), config.GetAppPort(), token)

	subject := "Confirm Your Subscription"
	body := fmt.Sprintf(`Hello!

Please confirm your subscription by clicking the link below:
%s

If you did not request this, please ignore this email.
`, confirmLink)

	message := []byte(fmt.Sprintf("From: %s\r\n", smtpCred.from) +
		fmt.Sprintf("To: %s\r\n", to[0]) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body)

	auth := smtp.PlainAuth("", smtpCred.user, smtpCred.password, smtpCred.host)

	err := smtp.SendMail(smtpCred.host+":"+smtpCred.port, auth, smtpCred.from, to, message)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	log.Println("Email sent successfully!")
	return nil
}

func SendWeatherEmail(toEmail string, city string, weather domain.Weather, token string) error {
	to := []string{toEmail}

	unsubscribeLink := fmt.Sprintf("%s:%s/api/unsubscribe/%s", config.GetAppHost(), config.GetAppPort(), token)

	subject := fmt.Sprintf("Weather update for %s", city)
	body := fmt.Sprintf(`Hello!

Here's the latest weather update for %s:

Temperature: %.1f°C
Humidity: %d%%
Description: %s

If you no longer wish to receive these emails, you can unsubscribe using the link below:
%s
`, city, weather.Temperature, weather.Humidity, weather.Description, unsubscribeLink)

	message := []byte(fmt.Sprintf("From: %s\r\n", smtpCred.from) +
		fmt.Sprintf("To: %s\r\n", to[0]) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body)

	auth := smtp.PlainAuth("", smtpCred.user, smtpCred.password, smtpCred.host)

	err := smtp.SendMail(smtpCred.host+":"+smtpCred.port, auth, smtpCred.from, to, message)
	if err != nil {
		return fmt.Errorf("error sending weather email: %w", err)
	}

	log.Println("Weather email sent to", toEmail)
	return nil
}
