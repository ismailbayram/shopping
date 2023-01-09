package domain

import userDomain "github.com/ismailbayram/shopping/pkg/users/domain/models"

type Order struct {
	ID     uint
	Number string
	User   userDomain.User
}
