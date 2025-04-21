package main

import (
	"local/knot/internal/app"
	"local/knot/internal/email"
	"log"
)

func main() {
	login := app.Login()

	imapServer, err := email.GetImapServer(login.Provider)
	if err != nil {
		log.Fatalf("Error: Could not get imap server domain name: %v", err)
	}

	smtpServer, err := email.GetSmtpServer(login.Provider)
	if err != nil {
		log.Fatalf("Error: Could not get smtp server domain name: %v", err)
	}

	imapCfg := email.ImapConfig{
		Server: imapServer,
		Port:   email.ImapPort,
		Email:  login.Email,
		Pass:   login.Pass,
	}

	imapClient, err := email.NewImapClient(imapCfg)
	if err != nil {
		log.Fatalf("Error: Could not create imap client: %v", err)
	}

	defer imapClient.Close()
	defer imapClient.Logout()

	smtpCfg := email.SmtpConfig{
		Server: smtpServer,
		Port:   email.SmtpPort,
		Email:  login.Email,
		Pass:   login.Pass,
	}

	a := app.NewApp(imapClient, smtpCfg)
	a.Run()
}
