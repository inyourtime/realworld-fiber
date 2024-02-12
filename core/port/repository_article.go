package port

import "realworld-go-fiber/core/domain"

type AddTagsPayload struct {
	Tags []domain.Tag
}

type AssignTagsParams struct {
	Article domain.Article
	Tags    []domain.Tag
}

type ArticleRepository interface {
	CreateArticle(arg domain.Article) (domain.Article, error)
	FilterTags(condition interface{}) ([]domain.Tag, error)
	AddTags(arg AddTagsPayload) ([]domain.Tag, error)
	AssignTags(arg AssignTagsParams) error
	FilterArticles(condition interface{}) ([]domain.Article, error)
}
