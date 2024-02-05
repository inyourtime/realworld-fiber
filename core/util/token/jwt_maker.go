package token

import (
	"errors"
	"fmt"
	"realworld-go-fiber/core/util/exception"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (maker JWTMaker) CreateToken(userID uint, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, exception.New(exception.TypeTokenInvalid, "invalid token", nil)
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, exception.New(exception.TypeTokenExpired, "token expired", err)
		}
		return nil, exception.New(exception.TypeTokenInvalid, "invalid token", err)
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, exception.New(exception.TypeTokenInvalid, "invalid token", nil)
	}

	return payload, nil
}
