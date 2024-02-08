package uc

import (
	"realworld-go-fiber/core/domain"
	"realworld-go-fiber/core/port"
	"realworld-go-fiber/core/util/exception"
)

func (u *articleUsecase) CreateArticle(arg port.CreateArticleParams) (domain.Article, error) {
	if arg.AuthArg.Payload == nil {
		return domain.Article{}, exception.New(exception.TypePermissionDenied, "authentication required", nil)
	}

	author, err := u.property.repo.User().FindOne(domain.User{ID: arg.AuthArg.Payload.UserID})
	if err != nil {
		return domain.Article{}, err
	}

	article := domain.NewArticle(domain.Article{
		Author:      author,
		Title:       arg.Article.Title,
		Description: arg.Article.Description,
		Body:        arg.Article.Body,
	})
	article.TagNames = arg.Article.TagNames

	created, tags, err := u.property.repo.Article().CreateArticleTransaction(article)
	if err != nil {
		return domain.Article{}, err
	}

	created.TagNames = []string{}
	for _, tag := range tags {
		created.TagNames = append(created.TagNames, tag.Name)
	}

	return created, nil
}
