package emails

import (
	"errors"
	"net/smtp"
	"project/config"

	"github.com/darkoatanasovski/htmltags"
	"github.com/jordan-wright/email"
)

type SendOptions struct {
	From     string
	FromName string // unused for now
	To       []string
	Bcc      []string
	Cc       []string
	Subject  string
	HtmlBody string
	TextBody string
}

// Send sends an email
func Send(options SendOptions) error {
	// drvr := os.Getenv("MAIL_DRIVER")
	host := config.MailHost
	port := config.MailPort
	user := config.MailUsername
	pass := config.MailPassword
	addr := host + ":" + port

	if options.From == "" {
		return errors.New("from is required")
	}

	if len(options.To) == 0 {
		return errors.New("to is required")
	}

	if options.Subject == "" {
		return errors.New("subject is required")
	}

	if options.HtmlBody == "" {
		return errors.New("html is required")
	}

	if options.TextBody == "" {
		nodes, errStripped := htmltags.Strip(options.HtmlBody, []string{}, true)

		if errStripped == nil {
			options.TextBody = nodes.ToString() // returns stripped HTML string
		}
	}

	e := email.NewEmail()
	e.From = options.From
	e.To = options.To
	e.Bcc = options.Bcc
	e.Cc = options.Cc
	e.Subject = options.Subject
	e.Text = []byte(options.TextBody)
	e.HTML = []byte(options.HtmlBody)
	var auth smtp.Auth
	if user == "" {
		auth = nil
	} else {
		auth = smtp.PlainAuth("", user, pass, host)
	}

	err := e.Send(addr, auth)

	if err != nil {
		// cfmt.Infoln(err.Error())
		config.LogStore.ErrorWithContext("Error at Send", err.Error())
		return err
	}

	return nil
}
