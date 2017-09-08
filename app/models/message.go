package models

import (
	"regexp"
	"strings"
	"time"

	"github.com/olliecoleman/alloy/app/services"
	"github.com/markbates/pop/nulls"
	"github.com/pkg/errors"
)

type SupportMessage struct {
	ID         int64
	Name       nulls.String      `db:"name"`
	Email      nulls.String      `db:"email"`
	Subject    nulls.String      `db:"subject"`
	Content    nulls.String      `db:"content"`
	InsertedAt time.Time         `db:"inserted_at"`
	UpdatedAt  time.Time         `db:"updated_at"`
	Errors     map[string]string `db:"-"`
}

func ListMessages(offset, limit int) ([]*SupportMessage, int, error) {
	supportMessages := []*SupportMessage{}
	err := services.DB.Select(&supportMessages, `
		SELECT 
			id, name, email, subject, content, inserted_at 
		FROM support_messages 
		ORDER BY inserted_at DESC
		OFFSET $1 LIMIT $2
	`, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	var count int
	err = services.DB.Get(&count, `SELECT count(*) FROM support_messages`)
	if err != nil {
		return nil, 0, err
	}

	return supportMessages, count, nil
}

func (b *SupportMessage) Delete() error {
	_, err := services.DB.Exec("DELETE from support_messages WHERE id = $1", b.ID)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (m *SupportMessage) Validate() bool {
	m.Errors = make(map[string]string)
	re := regexp.MustCompile(".+@.+\\..+")

	email := m.Email.String
	matched := re.Match([]byte(email))

	if matched == false {
		m.Errors["Email"] = "Please enter a valid email address"
	}

	if strings.TrimSpace(m.Name.String) == "" {
		m.Errors["Name"] = "Please enter a valid name"
	}

	if strings.TrimSpace(m.Content.String) == "" {
		m.Errors["Content"] = "Please enter your message"
	}

	if strings.TrimSpace(m.Subject.String) == "" {
		m.Errors["Subject"] = "What is your query about?"
	}

	return len(m.Errors) == 0
}

func (m *SupportMessage) Create() error {
	m.InsertedAt = time.Now()
	m.UpdatedAt = time.Now()

	stmt, err := services.DB.PrepareNamed(`
		INSERT INTO support_messages (name, email, subject, content, inserted_at, updated_at)
		VALUES 			 			 (:name, :email, :subject, :content, :inserted_at, :updated_at)
		RETURNING id
	`)

	if err != nil {
		return errors.WithStack(err)
	}

	err = stmt.Get(&m.ID, m)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
