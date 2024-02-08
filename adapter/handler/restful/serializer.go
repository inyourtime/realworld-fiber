package restful

import (
	"realworld-go-fiber/core/domain"
	"time"
)

type User struct {
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Bio      *string `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

type UserResponse struct {
	User User `json:"user"`
}

func serializeUser(arg domain.User) User {
	return User{
		Email:    arg.Email,
		Username: arg.Username,
		Bio:      arg.Bio,
		Image:    arg.Image,
		Token:    arg.Token,
	}
}

type Profile struct {
	Username  string  `json:"username"`
	Bio       *string `json:"bio"`
	Image     *string `json:"image"`
	Following bool    `json:"following"`
}

type ProfileResponse struct {
	Profile Profile `json:"profile"`
}

func serializeProfile(arg domain.User) Profile {
	return Profile{
		Username:  arg.Username,
		Image:     arg.Image,
		Bio:       arg.Bio,
		Following: arg.IsFollowed,
	}
}

type Article struct {
	Slug        string   `json:"slug"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Body        string   `json:"body"`
	Tags        []string `json:"tagList"`
	Author      Profile  `json:"author"`
	Favorited   bool     `json:"favorited"`
	Favorites   int      `json:"favoritesCount"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}

type ArticleResponse struct {
	Article Article `json:"article"`
}

func serializeArticle(arg domain.Article) Article {
	return Article{
		Slug:        arg.Slug,
		Title:       arg.Title,
		Description: arg.Description,
		Body:        arg.Body,
		Tags:        arg.TagNames,
		Author:      serializeProfile(arg.Author),
		Favorited:   arg.IsFavorite,
		Favorites:   arg.FavoriteCount,
		CreatedAt:   arg.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   arg.UpdatedAt.Format(time.RFC3339),
	}
}
