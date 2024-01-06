package tools

import (
	"fmt"
	"log/slog"
	"net/smtp"
	"os"
	"regexp"
)

type Email struct {
	smtpHost string
	smtpPort string
	password string
	to       []string
	from     string
}

func NewEmail(smtpHost, password, smtpPort, from string) *Email {
	return &Email{
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		password: password,
		from:     from,
	}
}

func NewDefaultEmail() (*Email, error) {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	password := os.Getenv("SMTP_PASSWORD")
	email := os.Getenv("SMPT_USER")
	client := NewEmail(host, password, port, email)

	if !client.isEmail(email) {
		return nil, fmt.Errorf("invalid email %s", email)
	}

	return client, nil
}

func (e *Email) SetFrom(from string) error {
	if !e.isEmail(from) {
		return fmt.Errorf("invalid email %s", from)
	}

	e.from = from

	return nil
}

func (e *Email) AddTo(to string) error {
	if !e.isEmail(to) {
		slog.Error("invalid email", "email", to)

		return fmt.Errorf("invalid email %s", to)
	}

	e.to = append(e.to, to)

	return nil
}

func (e Email) isEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(emailRegex, email)

	if err != nil {
		slog.Error("Error when validate email", err)

		return false
	}

	return match
}

func (e Email) Send(message string, to ...string) error {
	for _, t := range to {
		err := (&e).AddTo(t)

		if err != nil {
			return err
		}
	}

	if e.from == "" {
		return fmt.Errorf("no email from")
	}

	if len(e.to) == 0 {
		return fmt.Errorf("no email to send")
	}

	auth := smtp.PlainAuth("", e.from, e.password, e.smtpHost)
	smtpConn := fmt.Sprintf("%s:%s", e.smtpHost, e.smtpPort)

	if err := smtp.SendMail(smtpConn, auth, e.from, e.to, []byte(message)); err != nil {
		slog.Error("Error when send email", err)

		return err
	}

	return nil
}
