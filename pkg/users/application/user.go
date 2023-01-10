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

type EmailSender interface {
	SendWelcomeEmail(string)
}

type UserCache interface {
	GetUserIDByVerificationToken(string) *uint
}

type UserService struct {
	repo        UserRepository
	emailSender EmailSender
	cache       UserCache
}

func NewUserService(userRepo UserRepository, emailSender EmailSender, cache UserCache) *UserService {
	return &UserService{
		repo:        userRepo,
		emailSender: emailSender,
		cache:       cache,
	}
}

func (us *UserService) GetByID(id uint) (*domain.User, error) {
	return us.repo.GetByID(id)
}

func (us *UserService) Register(email string, password string, firstName string, lastName string) error {
	existed, _ := us.repo.GetByEmail(email)
	if existed != nil {
		return domain.ErrorUserAlreadyExists
	}

	user, err := us.repo.Create(&domain.User{
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
	go us.emailSender.SendWelcomeEmail(user.Email)
	return nil
}

func (us *UserService) Verify(token string) error {
	userID := us.cache.GetUserIDByVerificationToken(token)
	if userID == nil {
		return domain.ErrorUserNotFound
	}

	user, err := us.repo.GetByID(*userID)
	if err != nil {
		log.Println(err)
		return domain.ErrorGeneral
	}

	user.IsVerified = true
	err = us.repo.Update(user)
	if err != nil {
		log.Println(err)
		return domain.ErrorGeneral
	}
	return nil
}
