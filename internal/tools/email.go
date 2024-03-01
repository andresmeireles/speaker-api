package tools

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/smtp"
	"os"
	"regexp"
	"strings"
)

type EmailService interface {
	SetFrom(from string) error
	Send(message string, to string) error
}

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
		to:       make([]string, 0),
	}
}

func (e *Email) SetFrom(from string) error {
	if !e.isEmail(from) {
		return fmt.Errorf("invalid email %s", from)
	}

	e.from = from

	return nil
}

func (e *Email) isEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(emailRegex, email)

	if err != nil {
		slog.Error("Error when validate email", err)

		return false
	}

	return match
}

func (e *Email) Send(message string, to string) error {
	e.to = []string{to}

	if os.Getenv("APP_MODE") == "dev" {
		return e.sendDevEmail(message)
	}

	client, err := e.setupClient()
	if err != nil {
		return err
	}

	if err = client.Mail(e.from); err != nil {
		slog.Error("Error when set mail producer", err)

		return err
	}

	for _, addr := range e.to {
		if err = client.Rcpt(addr); err != nil {
			slog.Error("Error when set recipient", err)

			return err
		}
	}

	wc, err := client.Data()
	if err != nil {
		slog.Error("Error when create data", err)

		return err
	}

	emailMessage := e.formatEmail(message)
	if _, err = wc.Write([]byte(emailMessage)); err != nil {
		slog.Error("Error when write data", err)

		return err
	}

	err = wc.Close()
	if err != nil {
		slog.Error("Error when close data", err)

		return err
	}

	return nil
}

func (e *Email) setupClient() (*smtp.Client, error) {
	smtpConn := fmt.Sprintf("%s:%s", e.smtpHost, e.smtpPort)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         e.smtpHost,
	}

	conn, err := tls.Dial("tcp", smtpConn, tlsConfig)
	if err != nil {
		slog.Error("Error on tls dial", err)

		return nil, err
	}

	client, err := smtp.NewClient(conn, e.smtpHost)
	if err != nil {
		slog.Error("Error when create client", err)

		return nil, err
	}

	auth := smtp.PlainAuth("", e.from, e.password, e.smtpHost)
	if err = client.Auth(auth); err != nil {
		slog.Error("Error when auth", err)

		return nil, err
	}

	return client, nil
}

func (e Email) sendDevEmail(message string) error {
	err := smtp.SendMail(e.smtpHost+":"+e.smtpPort, nil, e.from, e.to, []byte(e.formatEmail(message)))

	if err != nil {
		slog.Error("Error when send email", err)

		return err
	}

	slog.Info("Email sent")

	return nil
}

func (e Email) formatEmail(message string) string {
	return fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n%s",
		e.from,
		strings.Join(e.to, ","),
		"Code login",
		message,
	)
}
