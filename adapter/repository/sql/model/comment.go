package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body      string  `gorm:"not null"`
	Author    User    `gorm:"foreignKey:AuthorID"`
	AuthorID  uint    `gorm:"not null"`
	Article   Article `gorm:"foreignKey:ArticleID"`
	ArticleID uint    `gorm:"not null"`
}
