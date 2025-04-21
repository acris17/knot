package email

import (
	"errors"
	"fmt"
	"slices"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
)

type ImapConfig struct {
	Server string
	Port   string
	Email  string
	Pass   string
}

func NewImapClient(cfg ImapConfig) (*imapclient.Client, error) {
	serverAddr := fmt.Sprintf("%v:%v", cfg.Server, cfg.Port)

	client, err := imapclient.DialTLS(serverAddr, nil)
	if err != nil {
		return nil, err
	}

	if err := client.Login(cfg.Email, cfg.Pass).Wait(); err != nil {
		return nil, err
	}

	return client, nil
}

func ListMailboxes(client *imapclient.Client) ([]string, error) {
	mailboxes, err := client.List("", "*", nil).Collect()
	if err != nil {
		return nil, err
	}

	var mailboxNames []string

	for _, mailbox := range mailboxes {
		mailboxNames = append(mailboxNames, mailbox.Mailbox)
	}

	return mailboxNames, nil
}

func PullRecentMessages(client *imapclient.Client) ([]*imapclient.FetchMessageBuffer, error) {
	inbox, err := client.Select("INBOX", nil).Wait()
	if err != nil {
		return nil, err
	}

	if inbox.NumMessages == 0 {
		return nil, errors.New("inbox is empty")
	}

	var seqSet imap.SeqSet
	var start uint32 = 1

	// Get five most recent messages
	if inbox.NumMessages > 5 {
		start = inbox.NumMessages - 4
	}

	seqSet.AddRange(start, inbox.NumMessages)

	textSection := &imap.FetchItemBodySection{
		Specifier: imap.PartSpecifierText,
	}

	fetchOptions := &imap.FetchOptions{
		Envelope: true,
		BodySection: []*imap.FetchItemBodySection{
			textSection,
		},
	}

	messages, err := client.Fetch(seqSet, fetchOptions).Collect()
	if err != nil {
		return nil, err
	}

	// Order by most recent mail (descending)
	var ordered []*imapclient.FetchMessageBuffer

	for _, message := range slices.Backward(messages) {
		ordered = append(ordered, message)
	}

	return ordered, nil
}
