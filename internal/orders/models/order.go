package models

import userDomain "github.com/ismailbayram/shopping/internal/users/models"

type Order struct {
	ID     uint
	Number string
	User   userDomain.User
}
