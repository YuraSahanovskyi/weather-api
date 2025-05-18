package email

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/YuraSahanovskyi/weather-api/internal/config"
	"github.com/YuraSahanovskyi/weather-api/internal/domain"
	"github.com/YuraSahanovskyi/weather-api/internal/templates"
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

	templates.LoadTemplates()
}

func buildEmail(to, subject, body string) ([]byte, error) {
	var buf bytes.Buffer
	err := templates.SMTPHeaderTemplate.Execute(&buf, map[string]string{
		"From":    smtpCred.from,
		"To":      to,
		"Subject": subject,
		"Body":    body,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to render email headers: %w", err)
	}
	return buf.Bytes(), nil
}

func sendEmail(to string, subject string, body string) error {
	msg, err := buildEmail(to, subject, body)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", smtpCred.user, smtpCred.password, smtpCred.host)

	err = smtp.SendMail(smtpCred.host+":"+smtpCred.port, auth, smtpCred.from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	log.Println("Email sent successfully to", to)
	return nil
}

func SendConfirmEmail(toEmail string, token string) error {
	confirmLink := fmt.Sprintf("%s:%s/api/confirm/%s", config.GetAppHost(), config.GetAppPort(), token)

	var bodyBuf bytes.Buffer
	err := templates.ConfirmTemplate.Execute(&bodyBuf, map[string]string{
		"ConfirmLink": confirmLink,
	})
	if err != nil {
		return fmt.Errorf("failed to render confirm email template: %w", err)
	}

	return sendEmail(toEmail, "Confirm Your Subscription", bodyBuf.String())
}

func SendWeatherEmail(toEmail string, city string, weather domain.Weather, token string) error {
	unsubscribeLink := fmt.Sprintf("%s:%s/api/unsubscribe/%s", config.GetAppHost(), config.GetAppPort(), token)

	var bodyBuf bytes.Buffer
	err := templates.WeatherTemplate.Execute(&bodyBuf, map[string]interface{}{
		"City":            city,
		"Temperature":     weather.Temperature,
		"Humidity":        weather.Humidity,
		"Description":     weather.Description,
		"UnsubscribeLink": unsubscribeLink,
	})
	if err != nil {
		return fmt.Errorf("failed to render weather email template: %w", err)
	}

	return sendEmail(toEmail, fmt.Sprintf("Weather update for %s", city), bodyBuf.String())
}
