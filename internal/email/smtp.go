package email

import (
	"fmt"
	"net/smtp"
	"strings"
)

type SmtpConfig struct {
	Server string
	Port   string
	Email  string
	Pass   string
}

func CreateMessage(from string, to string, subject string, body string) []byte {
	message := strings.Join([]string{
		fmt.Sprintf("From: %s", from),
		fmt.Sprintf("To: %s", to),
		fmt.Sprintf("Subject: %s", subject),
		"",
		body,
	}, "\r\n")

	return []byte(message)
}

func Send(message []byte, recipients []string, cfg SmtpConfig) error {
	auth := smtp.PlainAuth("", cfg.Email, cfg.Pass, cfg.Server)
	serverAddr := fmt.Sprintf("%v:%v", cfg.Server, cfg.Port)

	err := smtp.SendMail(serverAddr, auth, cfg.Email, recipients, message)
	if err != nil {
		return err
	}

	return nil
}
