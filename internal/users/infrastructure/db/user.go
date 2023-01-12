package infrastructure

import (
	domain "github.com/ismailbayram/shopping/internal/users/domain/models"
	"gorm.io/gorm"
)

type UserDBRepository struct {
	DB *gorm.DB
}

func (ur *UserDBRepository) Create(domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (ur *UserDBRepository) Update(domain.User) error {
	return nil
}

func (ur *UserDBRepository) GetByID(uint) (domain.User, error) {
	return domain.User{}, nil
}

func (ur *UserDBRepository) GetByEmail(string) (domain.User, error) {
	return domain.User{}, nil
}

func (ur *UserDBRepository) GetByToken(string) (domain.User, error) {
	return domain.User{}, nil
}

func (ur *UserDBRepository) All() ([]domain.User, error) {
	return nil, nil
}
