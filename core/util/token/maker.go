package token

import (
	"realworld-go-fiber/core/util/exception"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Maker interface {
	CreateToken(userID uint, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}

type Payload struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func NewPayload(userID uint, duration time.Duration) (*Payload, error) {
	now := time.Now()
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, exception.New(exception.TypeInternal, "failed generate token id", err)
	}
	payload := &Payload{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
		},
	}
	return payload, nil
}
