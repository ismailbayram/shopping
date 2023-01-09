package domain

type Email struct {
	User       *User
	Email      string
	IsVerified bool
	IsPrimary  bool
}
