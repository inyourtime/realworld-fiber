package sql

import (
	"realworld-go-fiber/core/port"

	"gorm.io/gorm"
)

type sqlRepo struct {
	db          *gorm.DB
	logger      port.Logger
	userRepo    port.UserRepository
	articleRepo port.ArticleRepository
}

func NewSQLRepository(db *gorm.DB, logger port.Logger) port.Repository {
	return &sqlRepo{
		db:          db,
		logger:      logger,
		userRepo:    NewUserRepository(db),
		articleRepo: NewArticleRepository(db),
	}
}

func (r *sqlRepo) User() port.UserRepository {
	return r.userRepo
}

func (r *sqlRepo) Article() port.ArticleRepository {
	return r.articleRepo
}
