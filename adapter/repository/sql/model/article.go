package model

import (
	"realworld-go-fiber/core/domain"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Slug           string `gorm:"uniqueIndex;not null"`
	Title          string `gorm:"not null"`
	Description    string `gorm:"not null"`
	Body           string `gorm:"not null"`
	Author         User   `gorm:"foreignKey:AuthorID"`
	AuthorID       uint   `gorm:"not null"`
	Tags           []Tag  `gorm:"many2many:article_tags"`
	FavoritedUsers []User `gorm:"many2many:favorite_articles"`
	Comments       []Comment
}

type Tag struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func (data Tag) ToDomain() domain.Tag {
	return domain.Tag{
		ID:   data.ID,
		Name: data.Name,
	}
}

func (data Article) ToDomain() domain.Article {
	return domain.Article{
		ID:          data.ID,
		Slug:        data.Slug,
		Title:       data.Title,
		Description: data.Description,
		Body:        data.Body,
		Author:      data.Author.ToDomain(),
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func AsArticle(arg domain.Article) Article {
	return Article{
		Model: gorm.Model{
			ID:        arg.ID,
			CreatedAt: arg.CreatedAt,
			UpdatedAt: arg.UpdatedAt,
		},
		Slug:        arg.Slug,
		Title:       arg.Title,
		Description: arg.Description,
		Body:        arg.Body,
		AuthorID:    arg.Author.ID,
		Author:      AsUser(arg.Author),
	}
}

func AsTag(arg domain.Tag) Tag {
	return Tag{
		Model: gorm.Model{
			ID: arg.ID,
		},
		Name: arg.Name,
	}
}
