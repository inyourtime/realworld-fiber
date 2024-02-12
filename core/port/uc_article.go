package port

import "realworld-go-fiber/core/domain"

type CreateArticleParams struct {
	Article domain.Article
	AuthArg AuthParams
}

type GetArticleParams struct {
	Article domain.Article
	AuthArg AuthParams
}

type ArticleUsecase interface {
	CreateArticle(arg CreateArticleParams) (domain.Article, error)
	GetArticle(arg GetArticleParams) (domain.Article, error)
}
