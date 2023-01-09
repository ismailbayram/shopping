package application

import domain "github.com/ismailbayram/shopping/pkg/users/domain/models"

type EmailRepository interface {
	Create(uint, string) (error, *domain.Email)
	Verify(*domain.Email) error
	GetByEmail(string) (error, *domain.Email)
	GetByUser(*domain.User) (error, []domain.Email)
}
