package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(8)
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err)

	err = CheckPassword(password, hashedPassword)
	assert.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashedPassword)
	assert.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	differentHashedPassword, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, differentHashedPassword)
	assert.NotEqual(t, hashedPassword, differentHashedPassword)
}
