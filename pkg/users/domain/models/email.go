package domain

import "errors"

var (
	ErrorEmailNotFound = errors.New("Email not found.")
)

type Email struct {
	User       *User
	Email      string
	IsVerified bool
	IsPrimary  bool
}
