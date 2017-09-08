package models

import (
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/olliecoleman/alloy/app/services"
	"github.com/lib/pq"
	"github.com/markbates/pop/nulls"
	"github.com/pkg/errors"
)

type AdminUser struct {
	ID           int64
	Name         nulls.String      `db:"name"`
	Email        nulls.String      `db:"email"`
	PasswordHash nulls.String      `db:"password_hash"`
	InsertedAt   time.Time         `db:"inserted_at"`
	UpdatedAt    time.Time         `db:"updated_at"`
	Password     string            `db:"-"`
	Errors       map[string]string `db:"-"`
}

func NewAdminUser(name, email, password string) *AdminUser {
	return &AdminUser{
		Name:     nulls.NewString(name),
		Email:    nulls.NewString(downcase(email)),
		Password: password,
	}
}

func GetAdminUser(id int64) (*AdminUser, error) {
	sql := `
		SELECT id, name, email, password_hash
		FROM admin_users
		WHERE id=$1
	`

	adminUser := AdminUser{}
	err := services.DB.Get(&adminUser, sql, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &adminUser, nil
}

func GetAdminByEmail(email string) (*AdminUser, error) {
	sql := `
		SELECT id, name, email, password_hash, inserted_at, updated_at 
		FROM admin_users
		WHERE email=$1
	`

	adminUser := AdminUser{}
	err := services.DB.Get(&adminUser, sql, email)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &adminUser, nil
}

func (u *AdminUser) CheckAuth(password string) bool {
	hashedP := []byte(u.PasswordHash.String)
	p := []byte(password)

	err := bcrypt.CompareHashAndPassword(hashedP, p)
	if err != nil {
		return false
	}

	return true
}

func (u *AdminUser) Create() error {
	u.InsertedAt = time.Now()
	u.UpdatedAt = time.Now()
	valid := u.Validate()

	if !valid {
		return errors.New("Invalid email/password")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.WithStack(err)
	}
	u.PasswordHash = nulls.NewString(string(hash))

	stmt, err := services.DB.PrepareNamed(`
		INSERT INTO admin_users (name, email, password_hash, inserted_at, updated_at)
		VALUES 			  (:name, :email, :password_hash, :inserted_at, :updated_at)
		RETURNING id
	`)

	if err != nil {
		return errors.WithStack(err)
	}

	err = stmt.Get(&u.ID, u)
	if err != nil {
		if pgerr, ok := err.(*pq.Error); ok {
			if pgerr.Code == "23505" && pgerr.Constraint == "admin_users_email_index" {
				return ErrAlreadyTaken
			}
		}

		return errors.WithStack(err)
	}

	return nil
}

func (u *AdminUser) Validate() bool {
	u.Errors = make(map[string]string)
	re := regexp.MustCompile(".+@.+\\..+")

	email := strings.TrimSpace(u.Email.String)
	password := strings.TrimSpace(u.Password)
	matched := re.Match([]byte(email))

	if matched == false {
		u.Errors["Email"] = "Please enter a valid email address"
	}

	if len(password) <= 6 {
		u.Errors["Password"] = "Password must be atleast 6 characters"
	}

	return len(u.Errors) == 0
}
