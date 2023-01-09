package domain

import "fmt"

type User struct {
	ID        uint
	Email     string
	FirstName string
	LastName  string
	IsActive  bool
	Phone     string
}

func (u *User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}
