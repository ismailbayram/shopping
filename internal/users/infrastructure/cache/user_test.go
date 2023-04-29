package infrastructure

import (
	"github.com/ismailbayram/shopping/internal/users/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUserCache(t *testing.T) {
	cache := NewUserCache()
	assert.Empty(t, cache.verificationTokens)
}

func TestUserCache_SetUserVerificationToken(t *testing.T) {
	cache := NewUserCache()
	cache.SetUserVerificationToken("token", 1)
	assert.Equal(t, uint(1), cache.verificationTokens["token"])
}

func TestUserCache_GetUserIDByVerificationToken(t *testing.T) {
	cache := NewUserCache()
	cache.SetUserVerificationToken("token", 1)

	userID, err := cache.GetUserIDByVerificationToken("xx")
	assert.Equal(t, domain.ErrorUserNotFound, err)
	assert.Zero(t, userID)

	userID, err = cache.GetUserIDByVerificationToken("token")
	assert.Nil(t, err)
	assert.Equal(t, uint(1), userID)
}
