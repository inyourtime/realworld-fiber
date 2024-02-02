package domain

import "time"

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
