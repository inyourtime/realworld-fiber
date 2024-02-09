package port

import "realworld-go-fiber/core/domain"

type AddTagsPayload struct {
	Tags []string
}

type ArticleRepository interface {
	// FilterTags(condition interface{}) ([]domain.Tag, error)
	CreateArticleTransaction(article domain.Article) (domain.Article, []domain.Tag, error)
}
