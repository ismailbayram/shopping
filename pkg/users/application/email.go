package application

import domain "github.com/ismailbayram/shopping/pkg/users/domain/models"

type EmailRepository interface {
	Create(uint, string) (*domain.Email, error)
	Verify(*domain.Email) error
	GetByEmail(string) (*domain.Email, error)
	GetPrimaryOfUser(*domain.User) (*domain.Email, error)
}
