package infrastructure

import "github.com/ismailbayram/shopping/internal/users/models"

type UserCache struct {
	verificationTokens map[string]uint
}

func NewUserCache() *UserCache {
	return &UserCache{
		verificationTokens: map[string]uint{},
	}
}

func (c *UserCache) SetUserVerificationToken(token string, userID uint) {
	c.verificationTokens[token] = userID
}

func (c *UserCache) GetUserIDByVerificationToken(token string) (uint, error) {
	userID, ok := c.verificationTokens[token]
	if !ok {
		return 0, models.ErrorUserNotFound
	}
	return userID, nil
}
