package infrastructure

import (
	"errors"
	"github.com/ismailbayram/shopping/internal/users/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type UserDB struct {
	ID         uint      `gorm:"primarykey"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime"`
	Email      string    `gorm:"not null;uniqueIndex"`
	FirstName  string    `gorm:"not null"`
	LastName   string    `gorm:"not null"`
	IsActive   bool      `gorm:"not null"`
	IsVerified bool      `gorm:"not null"`
	IsAdmin    bool      `gorm:"not null"`
	Password   string    `gorm:"not null"`
	Token      string    `gorm:"not null;uniqueIndex"`
}

func (UserDB) TableName() string {
	return "user_users"
}

type UserDBRepository struct {
	db *gorm.DB
}

func NewUserDBRepository(db *gorm.DB) *UserDBRepository {
	return &UserDBRepository{db: db}
}

func (ur *UserDBRepository) Create(user models.User) (models.User, error) {
	userDB := ToUserDB(user)

	result := ur.db.Create(&userDB)
	if result.Error != nil {
		logrus.WithFields(logrus.Fields{
			"user":  user,
			"error": result.Error,
		}).Error("UserDB.Create")
		return models.User{}, models.ErrorGeneral
	}

	return ToUser(userDB), nil
}

func (ur *UserDBRepository) Update(user models.User) error {
	userDB := ToUserDB(user)
	if err := ur.db.UpdateColumns(&userDB).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user":  user,
			"error": err,
		}).Error("UserDB.Update")
		return models.ErrorGeneral
	}
	return nil
}

func (ur *UserDBRepository) GetByID(id uint) (models.User, error) {
	var userDB UserDB
	result := ur.db.Where("id = ?", id).First(&userDB)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.User{}, models.ErrorUserNotFound
		}
		logrus.WithFields(logrus.Fields{
			"id":    id,
			"error": result.Error,
		}).Error("UserDB.GetByID")
		return models.User{}, models.ErrorGeneral
	}

	return ToUser(userDB), nil
}

func (ur *UserDBRepository) GetByEmail(email string) (models.User, error) {
	var userDB UserDB
	result := ur.db.Where("email = ?", email).First(&userDB)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.User{}, models.ErrorUserNotFound
		}
		logrus.WithFields(logrus.Fields{
			"email": email,
			"error": result.Error,
		}).Error("UserDB.GetByEmail")
		return models.User{}, models.ErrorGeneral
	}

	return ToUser(userDB), nil
}

func (ur *UserDBRepository) GetByToken(token string) (models.User, error) {
	var userDB UserDB
	result := ur.db.Where("token = ?", token).First(&userDB)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.User{}, models.ErrorUserNotFound
		}
		logrus.WithFields(logrus.Fields{
			"token": token,
			"error": result.Error,
		}).Error("UserDB.GetByToken")
		return models.User{}, models.ErrorGeneral
	}

	return ToUser(userDB), nil
}

func (ur *UserDBRepository) All(filters map[string]interface{}) ([]models.User, error) {
	var userDBs []UserDB

	//db.Where(&User{Name: "iso", Email: "b@s.com})
	result := ur.db.Order("id asc").Where(filters).Find(&userDBs)
	if result.Error != nil {
		logrus.WithFields(logrus.Fields{
			"filters": filters,
			"error":   result.Error,
		}).Error("UserDB.All")
		return nil, models.ErrorGeneral
	}

	users := make([]models.User, len(userDBs))

	for i, uDB := range userDBs {
		users[i] = ToUser(uDB)
	}

	return users, result.Error
}
