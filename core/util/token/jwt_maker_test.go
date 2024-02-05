package token

import (
	"realworld-go-fiber/core/util"
	"realworld-go-fiber/core/util/exception"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJwtMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	assert.NoError(t, err)

	var userID uint = 99
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(userID, duration)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	assert.NotZero(t, payload.ID)
	assert.Equal(t, userID, payload.UserID)

	assert.WithinDuration(t, issuedAt, payload.IssuedAt.Time, time.Second)
	assert.WithinDuration(t, expiredAt, payload.ExpiresAt.Time, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	assert.NoError(t, err)

	token, payload, err := maker.CreateToken(99, -time.Minute)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	assert.Error(t, err)

	fail, ok := err.(*exception.Exception)
	assert.True(t, ok)
	assert.Equal(t, exception.TypeTokenExpired, fail.Type)
	assert.Nil(t, payload)
}

func TestInvalidJWTToken(t *testing.T) {
	payload, err := NewPayload(99, time.Minute)
	assert.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	assert.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	assert.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)

	fail, ok := err.(*exception.Exception)
	assert.True(t, ok)
	assert.Equal(t, exception.TypeTokenInvalid, fail.Type)
	assert.Nil(t, payload)
}
