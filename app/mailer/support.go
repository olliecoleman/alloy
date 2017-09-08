package mailer

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/olliecoleman/alloy/app/models"
)

func NewSupportMail(message *models.SupportMessage) {
	m := NewMail(message.Email.String, "Thanks for contacting us.")
	m.Vars["Message"] = message.Content.String

	err := m.Send("mailer/support-messages/new")
	if err != nil {
		log.Printf("support email sending failed: %+v", err)
	}
}

func NewSupportNotification(message *models.SupportMessage) {
	m := NewMail(envy.Get("MAILER_FROM", "support@alloydev.me"), "New message: "+message.Subject.String)
	m.ReplyTo = message.Email.String
	m.Vars["Email"] = message.Email.String
	m.Vars["Name"] = message.Name.String
	m.Vars["Message"] = message.Content.String

	err := m.Send("mailer/support-messages/notification")
	if err != nil {
		log.Printf("support email sending failed: %+v", err)
	}
}
