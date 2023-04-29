package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)

var (
	ErrorGeneral           = errors.New("Something went wrong, please try again.")
	ErrorUserNotFound      = errors.New("User not found.")
	ErrorUserAlreadyExists = errors.New("User with this e-mail already exists.")
	ErrorWrongPassword     = errors.New("Password is invalid.")
	ErrorPasswordUnmatched = errors.New("Passwords do not match.")
	ErrorUserNotVerified   = errors.New("Please check your mail inbox to verify your account.")
)

type User struct {
	ID         uint
	Email      string
	FirstName  string
	LastName   string
	IsActive   bool
	IsVerified bool
	IsAdmin    bool
	Password   string
	Token      string
}

func (u *User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (u *User) SetPassword(password string) {
	u.Password = hashPassword(password)
}

func (u *User) CheckPassword(password string) error {
	if hashPassword(password) != u.Password {
		return ErrorWrongPassword
	}
	return nil
}

func hashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s%s", password, os.Getenv("SECRET_KEY"))))
	return hex.EncodeToString(h.Sum(nil))
}
