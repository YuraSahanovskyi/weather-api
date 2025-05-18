package templates

import (
	"embed"
	htmltmpl "html/template"
	"log"
	"text/template"
)

//go:embed email/*.html email/*.txt
var emailTemplatesFS embed.FS

var (
	ConfirmTemplate    *htmltmpl.Template
	WeatherTemplate    *htmltmpl.Template
	SMTPHeaderTemplate *template.Template
)

func LoadTemplates() {
	var err error

	ConfirmTemplate, err = htmltmpl.ParseFS(emailTemplatesFS, "email/confirm.html")
	if err != nil {
		log.Fatalf("failed to parse confirm template: %v", err)
	}

	WeatherTemplate, err = htmltmpl.ParseFS(emailTemplatesFS, "email/weather.html")
	if err != nil {
		log.Fatalf("failed to parse weather template: %v", err)
	}

	SMTPHeaderTemplate, err = template.ParseFS(emailTemplatesFS, "email/email_header.txt")
	if err != nil {
		log.Fatalf("failed to parse smtp header template: %v", err)
	}
}
