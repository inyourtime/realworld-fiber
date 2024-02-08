package domain

import (
	"strings"
	"time"
)

type Article struct {
	ID            uint
	Slug          string
	Title         string
	Description   string
	Body          string
	Author        User
	AuthorID      uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	TagNames      []string
	IsFavorite    bool
	FavoriteCount int
}

func (article *Article) SetTitle(value string) {
	article.Title = value
	article.Slug = strings.ToLower(strings.ReplaceAll(value, " ", "-"))
}

func NewArticle(arg Article) Article {
	article := Article{
		Description: arg.Description,
		Body:        arg.Body,
		Author:      arg.Author,
	}
	article.SetTitle(arg.Title)
	return article
}

type Tag struct {
	ID   uint
	Name string
}
