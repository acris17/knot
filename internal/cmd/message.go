package cmd

import (
	"fmt"
	"local/knot/internal/email"
	"strconv"
	"strings"
	"time"

	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/jaytaylor/html2text"
)

func List(messages []*imapclient.FetchMessageBuffer) {
	for i, message := range messages {
		if message.Envelope == nil {
			continue
		}

		subject := message.Envelope.Subject
		date := message.Envelope.Date.Format(time.DateTime)

		fmt.Printf("%v %v %v\n", i, date, subject)
	}
}

func Read(args []string, messages []*imapclient.FetchMessageBuffer) {
	if len(args) == 0 || len(messages) == 0 {
		return
	}

	index, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}

	first := 0
	last := len(messages) - 1

	// Ensure the given index can be used to access messages
	if index < first || index > last {
		return
	}
	choice := messages[index]

	if len(choice.BodySection) == 0 {
		return
	}

	section := choice.BodySection[0]
	opts := html2text.Options{
		TextOnly:  true,
		OmitLinks: true,
	}

	text, err := html2text.FromString(string(section.Bytes), opts)
	if err != nil {
		fmt.Printf("Error: Could not parse mail html: %v\n", err)
		return
	}

	fmt.Println(text)
}

func Send(to string, subject string, body string, cfg email.SmtpConfig) {
	if strings.TrimSpace(to) == "" {
		return
	}
	if strings.TrimSpace(subject) == "" {
		return
	}
	if strings.TrimSpace(body) == "" {
		return
	}

	message := email.CreateMessage(cfg.Email, to, subject, body)
	recipients := []string{to}

	if err := email.Send(message, recipients, cfg); err != nil {
		fmt.Printf("Error: Could not send message: %v\n", err)
	}

	fmt.Println("Info: Message sent successfully")
}
