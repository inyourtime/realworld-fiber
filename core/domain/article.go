package domain

import "time"

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
