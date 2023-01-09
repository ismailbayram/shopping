package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_GetFullName(t *testing.T) {
	user := User{
		ID:        0,
		Email:     "iso@iso.com",
		FirstName: "ismail",
		LastName:  "bayram",
		IsActive:  false,
		Phone:     "",
	}
	assert.Equal(t, "ismail bayram", user.GetFullName())
}
