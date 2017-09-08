package models

import (
	"errors"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var (
	ErrAlreadyTaken = errors.New("is already taken")
)

func downcase(str string) string {
	return strings.TrimSpace(strings.ToLower(str))
}

func newUUID() string {
	return uuid.NewV4().String()
}
