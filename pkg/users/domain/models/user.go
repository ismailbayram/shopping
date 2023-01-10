package domain

import (
	"errors"
	"fmt"
)

var (
	ErrorUserNotFound      = errors.New("User not found.")
	ErrorUserAlreadyExists = errors.New("User with this e-mail already exists.")
)

type User struct {
	ID         uint
	Email      string
	FirstName  string
	LastName   string
	IsActive   bool
	IsVerified bool
	Password   string
}

func (u *User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}
