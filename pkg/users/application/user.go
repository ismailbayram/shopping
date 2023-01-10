package application

import (
	domain "github.com/ismailbayram/shopping/pkg/users/domain/models"
	"log"
)

type UserRepository interface {
	Create(*domain.User) (*domain.User, error)
	Update(*domain.User) error
	GetByID(uint) (*domain.User, error)
	GetByEmail(string) (*domain.User, error)
	GetByToken(string) (*domain.User, error)
	All() ([]domain.User, error)
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) Register(email string, password string, firstName string, lastName string) error {
	existed, _ := us.userRepo.GetByEmail(email)
	if existed != nil {
		return domain.ErrorUserAlreadyExists
	}

	_, err := us.userRepo.Create(&domain.User{
		Email:      email,
		FirstName:  firstName,
		LastName:   lastName,
		IsActive:   true,
		IsVerified: false,
		Password:   password,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	// TODO: send email
	return nil
}

func (us *UserService) GetByID(id uint) (*domain.User, error) {
	return us.userRepo.GetByID(id)
}
