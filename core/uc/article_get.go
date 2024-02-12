package uc

import (
	"realworld-go-fiber/core/domain"
	"realworld-go-fiber/core/port"
	"realworld-go-fiber/core/util/exception"
)

func (u *articleUsecase) GetArticle(arg port.GetArticleParams) (domain.Article, error) {
	if arg.Article.Slug == "" {
		return domain.Article{}, exception.New(exception.TypeValidation, "slug is required", nil)
	}

	articles, err := u.property.repo.Article().FilterArticles(map[string]interface{}{"slug": arg.Article.Slug})
	if err != nil {
		return domain.Article{}, err
	}

	if len(articles) == 0 {
		return domain.Article{}, exception.New(exception.TypeNotFound, "no article found", nil)
	}

	article := articles[0]

	if arg.AuthArg.Payload == nil {
		return article, nil
	}

	follows, err := u.property.repo.User().FilterFollow(domain.User{ID: arg.AuthArg.Payload.UserID}, article.Author)
	if err != nil {
		return domain.Article{}, err
	}

	if len(follows) > 0 {
		article.Author.IsFollowed = true
	}

	return article, nil
}
