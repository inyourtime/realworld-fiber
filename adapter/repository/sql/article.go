package sql

import (
	"realworld-go-fiber/core/port"

	"gorm.io/gorm"
)

type articleRepo struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) port.ArticleRepository {
	return &articleRepo{db: db}
}
