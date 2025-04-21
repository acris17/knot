package email

import (
	"errors"
)

const (
	AppleProvider   = "apple"
	AppleImapServer = "imap.mail.me.com"
	AppleStmpServer = "smtp.mail.me.com"
)

const (
	GmailProvider   = "gmail"
	GmailImapServer = "imap.gmail.com"
	GmailStmpServer = "smtp.gmail.com"
)

const (
	ImapPort = "993"
	SmtpPort = "587"
)

func GetImapServer(provider string) (string, error) {
	var imapServer string

	switch provider {
	case AppleProvider:
		imapServer = AppleImapServer
	case GmailProvider:
		imapServer = GmailImapServer
	default:
		return "", errors.New("given provider is not apple or gmail")
	}

	return imapServer, nil
}

func GetSmtpServer(provider string) (string, error) {
	var smtpServer string

	switch provider {
	case AppleProvider:
		smtpServer = AppleStmpServer
	case GmailProvider:
		smtpServer = GmailStmpServer
	default:
		return "", errors.New("given provider is not apple or gmail")
	}

	return smtpServer, nil
}
