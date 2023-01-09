package application

import (
	domain "github.com/ismailbayram/shopping/pkg/users/domain/models"
)

type UserRepository interface {
	Create(*domain.User) error
	Update(*domain.User) error
	GetByID(uint) (*domain.User, error)
	GetByToken(string) (*domain.User, error)
	All() ([]domain.User, error)
}

type UserService struct {
	userRepo  UserRepository
	emailRepo EmailRepository
}

func NewUserService(userRepo UserRepository, emailRepo EmailRepository) *UserService {
	return &UserService{
		userRepo:  userRepo,
		emailRepo: emailRepo,
	}
}

func (us *UserService) GetByID(id uint) (*domain.User, error) {
	return us.userRepo.GetByID(id)
}

func (us *UserService) IsVerified(user *domain.User) bool {
	email, err := us.emailRepo.GetPrimaryOfUser(user)
	if err != nil {
		// TODO: log
		return false
	}

	return email.IsVerified
}
