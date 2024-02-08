package port

import "realworld-go-fiber/core/domain"

type CreateArticleParams struct {
	Article domain.Article
	AuthArg AuthParams
}

type ArticleUsecase interface {
	CreateArticle(arg CreateArticleParams) (domain.Article, error)
}
