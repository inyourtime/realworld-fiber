package model

import (
	"realworld-go-fiber/core/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email            string `gorm:"uniqueIndex;not null"`
	Username         string `gorm:"uniqueIndex;not null"`
	Password         string `gorm:"not null"`
	Bio              *string
	Image            *string
	Followings       []*User   `gorm:"many2many:user_follows"`
	FavoriteArticles []Article `gorm:"many2many:favorite_articles"`
}

func (data User) ToDomain() domain.User {
	return domain.User{
		ID:        data.ID,
		Email:     data.Email,
		Username:  data.Username,
		Password:  data.Password,
		Bio:       data.Bio,
		Image:     data.Image,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func AsUser(arg domain.User) User {
	return User{
		Model: gorm.Model{
			ID:        arg.ID,
			CreatedAt: arg.CreatedAt,
			UpdatedAt: arg.UpdatedAt,
		},
		Email:    arg.Email,
		Username: arg.Username,
		Password: arg.Password,
		Image:    arg.Image,
		Bio:      arg.Bio,
	}
}
