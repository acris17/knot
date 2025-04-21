package app

import (
	"fmt"
	"local/knot/internal/cmd"
	"local/knot/internal/email"
	"os"

	"github.com/emersion/go-imap/v2/imapclient"
)

type App struct {
	imapClient *imapclient.Client
	smtpCfg    email.SmtpConfig
	messages   []*imapclient.FetchMessageBuffer
}

func NewApp(imapClient *imapclient.Client, smtpCfg email.SmtpConfig) App {
	empty := []*imapclient.FetchMessageBuffer{}

	return App{imapClient, smtpCfg, empty}
}

func (a *App) Run() {
	for {
		input := Input("knot> ")
		action := Parse(input)

		if len(action) == 0 {
			continue
		}

		command := action[0]
		args := action[1:]

		a.Dispatch(command, args)
	}
}

func (a *App) Dispatch(command string, args []string) {
	switch command {
	case "exit":
		a.imapClient.Logout()
		a.imapClient.Close()
		os.Exit(0)
	case "version":
		cmd.Version()
	case "boxes":
		cmd.Boxes(a.imapClient)
	case "list":
		cmd.List(a.messages)
	case "read":
		cmd.Read(args, a.messages)
	case "send":
		to := Input("to: ")
		subject := Input("subject: ")
		body := Input("body: ")
		ok := Input("ok? [y/n]: ")

		if ok != "y" {
			fmt.Println("Info: Will not send mail")
			break
		}

		cmd.Send(to, subject, body, a.smtpCfg)
	case "pull":
		messages, err := email.PullRecentMessages(a.imapClient)
		if err != nil {
			break
		}

		a.messages = messages
		fmt.Printf("Info: Pulled %v messages\n", len(messages))
	}
}
