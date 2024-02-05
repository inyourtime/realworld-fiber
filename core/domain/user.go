package domain

import (
	"realworld-go-fiber/core/util"
	"time"
)

type User struct {
	ID         uint
	Email      string
	Username   string
	Password   string
	Bio        *string
	Image      *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IsFollowed bool
	Token      string
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return nil
}

func NewUser(arg User) (User, error) {
	user := User{
		Email:    arg.Email,
		Username: arg.Username,
		Bio:      arg.Bio,
		Image:    arg.Image,
	}

	if err := user.SetPassword(arg.Password); err != nil {
		return User{}, err
	}

	return user, nil
}
