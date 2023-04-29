package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestUser_GetFullName(t *testing.T) {
	user := User{
		ID:        0,
		Email:     "iso@iso.com",
		FirstName: "ismail",
		LastName:  "bayram",
		IsActive:  false,
	}
	assert.Equal(t, "ismail bayram", user.GetFullName())
}

func TestUser_SetPassword(t *testing.T) {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s%s", "PASSWORD", os.Getenv("SECRET_KEY"))))
	expectedPassword := hex.EncodeToString(h.Sum(nil))

	user := User{}
	user.SetPassword("PASSWORD")
	assert.Equal(t, expectedPassword, user.Password)
}

func TestUser_CheckPassword(t *testing.T) {
	user := User{}
	user.SetPassword("PASSWORD")

	assert.Equal(t, ErrorWrongPassword, user.CheckPassword("wron"))
	assert.Nil(t, user.CheckPassword("PASSWORD"))
}
