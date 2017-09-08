package mailer

import (
	"strconv"

	gomail "gopkg.in/gomail.v2"

	"github.com/gobuffalo/envy"
	"github.com/pkg/errors"
)

func setupSMTP() gomail.SendCloser {
	var err error
	port, err := strconv.Atoi(envy.Get("SMTP_PORT", "1025"))
	if err != nil {
		panic(err)
	}

	d := gomail.NewDialer(
		envy.Get("SMTP_HOST", "127.0.0.1"),
		port,
		envy.Get("SMTP_USERNAME", ""),
		envy.Get("SMTP_PASSWORD", ""),
	)

	s, err = d.Dial()
	if err != nil {
		panic(err)
	}
	return s
}

func sendWithSMTP(m *Message) error {
	message := gomail.NewMessage()
	message.SetHeader("From", m.From)
	message.SetHeader("To", m.To)
	message.SetHeader("Reply-To", m.ReplyTo)
	message.SetHeader("Subject", m.Subject)
	message.SetBody("text/html", m.Body)

	if m.Attachment != nil {
		message.Attach(m.Attachment.Path, gomail.Rename(m.Attachment.Name))
	}

	err := gomail.Send(s, message)
	if err != nil {
		return errors.Wrapf(err, "Could not send email to %s: %+v", m.To, err)
	}

	return nil
}
