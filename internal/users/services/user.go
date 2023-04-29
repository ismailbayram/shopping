package services

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ismailbayram/shopping/internal/users/models"
	"os"
)

type UserRepository interface {
	Create(models.User) (models.User, error)
	Update(models.User) error
	GetByID(uint) (models.User, error)
	GetByEmail(string) (models.User, error)
	GetByToken(string) (models.User, error)
	All(map[string]interface{}) ([]models.User, error)
}

type EmailSender interface {
	SendWelcomeEmail(string, string)
}

type UserCache interface {
	GetUserIDByVerificationToken(string) (uint, error)
	SetUserVerificationToken(string, uint)
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

func (us *UserService) GetByID(id uint) (models.User, error) {
	return us.repo.GetByID(id)
}

func (us *UserService) GetByToken(token string) (models.User, error) {
	return us.repo.GetByToken(token)
}

func (us *UserService) Login(email string, password string) (string, error) {
	user, err := us.repo.GetByEmail(email)
	if err != nil {
		return "", models.ErrorUserNotFound
	}

	if !user.IsVerified {
		return "", models.ErrorUserNotVerified
	}

	if err := user.CheckPassword(password); err != nil {
		return "", err
	}

	return user.Token, nil
}

func (us *UserService) Register(email string, password string, firstName string, lastName string) error {
	existed, _ := us.repo.GetByEmail(email)
	if existed.ID != uint(0) {
		return models.ErrorUserAlreadyExists
	}

	user := models.User{
		Email:      email,
		FirstName:  firstName,
		LastName:   lastName,
		IsActive:   true,
		IsVerified: false,
		IsAdmin:    false,
		Token:      generateToken(email),
	}
	user.SetPassword(password)

	createdUser, err := us.repo.Create(user)
	if err != nil {
		return err
	}
	go us.emailSender.SendWelcomeEmail(createdUser.Token, createdUser.Email)
	return nil
}

func (us *UserService) Verify(token string) error {
	user, err := us.repo.GetByToken(token)
	if err != nil {
		return models.ErrorUserNotFound
	}

	user.IsVerified = true
	err = us.repo.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) ChangePassword(user models.User, newPassword string) error {
	user.SetPassword(newPassword)
	if err := us.repo.Update(user); err != nil {
		return err
	}
	// TODO: changing password inform mail
	return nil
}

//func (us *UserService) generateAndSetVerificationToken(userID uint) string {
//	h := sha256.New()
//	h.Write([]byte(fmt.Sprintf("%d%s", userID, os.Getenv("SECRET_KEY"))))
//	token := hex.EncodeToString(h.Sum(nil))
//
//	us.cache.SetUserVerificationToken(token, userID)
//
//	return token
//}

func generateToken(email string) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s%s", email, os.Getenv("SECRET_KEY"))))
	return hex.EncodeToString(h.Sum(nil))
}
