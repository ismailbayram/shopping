package application

import domain "github.com/ismailbayram/shopping/pkg/users/domain/models"

type UserRepository interface {
	Create(*domain.User) error
	Update(*domain.User) error
	GetByID(uint) (error, *domain.User)
	GetByToken(string) (error, *domain.User)
	All() (error, []domain.User)
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

func (us *UserService) GetUserByID(id uint) (error, *domain.User) {
	return us.userRepo.GetByID(id)
}

func (us *UserService) IsVerified(user *domain.User) bool {
	return true
}
