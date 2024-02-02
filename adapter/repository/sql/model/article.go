package model

import "gorm.io/gorm"

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
