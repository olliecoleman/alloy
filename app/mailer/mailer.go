package mailer

import (
	"bytes"
	"fmt"
	"log"

	"github.com/gobuffalo/envy"
	"github.com/olliecoleman/alloy/app/views"
	"github.com/pkg/errors"
	"github.com/pressly/douceur/inliner"

	gomail "gopkg.in/gomail.v2"
)

type Message struct {
	To         string
	From       string
	Subject    string
	ReplyTo    string
	Body       string
	Attachment *Attachment
	Vars       map[string]interface{}
}

type Attachment struct {
	Name       string
	Path       string
	Downloaded bool
}

var s gomail.SendCloser
var host string

func init() {
	host = envy.Get("MAILER_HOST", "http://alloydev.me")
	s = setupSMTP()
}

func NewMail(to, subject string) *Message {
	from := envy.Get("MAILER_FROM", "support@alloydev.me")
	replyTo := envy.Get("MAILER_FROM", "support@alloydev.me")

	return &Message{
		To:      to,
		From:    from,
		ReplyTo: replyTo,
		Subject: subject,
		Vars: map[string]interface{}{
			"Meta": map[string]interface{}{
				"Env": envy.Get("ENVIRONMENT", "development"),
			},
		},
	}
}

func (m *Message) Send(templateName string) error {
	if m.To == "" || m.From == "" {
		return fmt.Errorf("invalid To (%s) or From (%s) address", m.To, m.From)
	}

	t, err := views.GetTemplate(templateName)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	err = t.ExecuteTemplate(buf, "base", m.Vars)
	if err != nil {
		return err
	}

	m.Body = buf.String()
	if m.Body == "" {
		return errors.New("empty body")
	}

	html, err := inliner.Inline(m.Body)
	if err != nil {
		log.Printf("email inlining failed: %+v\n", err)
	}

	m.Body = html

	log.Printf("Sending email \nTo: %s\n From: %s\n Subject: %s\n", m.To, m.From, m.Subject)
	return sendWithSMTP(m)
}
