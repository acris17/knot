package cmd

import (
	"fmt"
	"local/knot/internal/email"

	"github.com/emersion/go-imap/v2/imapclient"
)

func Boxes(client *imapclient.Client) {
	mailboxNames, err := email.ListMailboxes(client)
	if err != nil {
		fmt.Printf("Error: Could not get mailbox names: %v\n", err)
	}

	for _, name := range mailboxNames {
		fmt.Println(name)
	}
}
