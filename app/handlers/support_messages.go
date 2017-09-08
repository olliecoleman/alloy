package handlers

import (
	"net/http"
	"time"

	"github.com/olliecoleman/alloy/app/mailer"
	"github.com/olliecoleman/alloy/app/models"
	"github.com/olliecoleman/alloy/app/views"
	"github.com/markbates/pop/nulls"
	"github.com/pkg/errors"
)

type recaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []int     `json:"error-codes"`
}

func NewSupportMessage(w http.ResponseWriter, r *http.Request) error {
	v := views.New(r)
	message := models.SupportMessage{}

	v.Vars["Message"] = message
	v.Render(w, "support-messages/new")
	return nil
}

func CreateSupportMessage(w http.ResponseWriter, r *http.Request) error {
	v := views.New(r)
	message := models.SupportMessage{}
	message.Name = nulls.NewString(r.FormValue("name"))
	message.Email = nulls.NewString(r.FormValue("email"))
	message.Subject = nulls.NewString(r.FormValue("subject"))
	message.Content = nulls.NewString(r.FormValue("content"))

	if message.Validate() {
		err := message.Create()

		if err != nil {
			return StatusError{Code: 500, Err: errors.Wrapf(err, "Message: %v", message)}
		}

		go func() {
			mailer.NewSupportMail(&message)
			mailer.NewSupportNotification(&message)
		}()

		views.SuccessFlash(w, r, "Thank you for contacting us. We will get back to you within 2 business days")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	v.Vars["Message"] = message
	v.Render(w, "support-messages/new")
	return nil
}
