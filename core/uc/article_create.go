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

	assigningTags := []domain.Tag{}
	createdArticle := domain.Article{}

	err = u.property.repo.Atomic(func(r port.Repository) error {
		created, err := r.Article().CreateArticle(article)
		if err != nil {
			return err
		}
		createdArticle = created

		if len(arg.Article.TagNames) == 0 {
			return nil
		}

		existing, err := r.Article().FilterTags(map[string]interface{}{"name": arg.Article.TagNames})
		if err != nil {
			return err
		}

		assigningTags = append(assigningTags, existing...)

		tags := FilterTagNotExist(existing, arg.Article.TagNames)
		if len(tags) > 0 {
			craetedTags, err := r.Article().AddTags(port.AddTagsPayload{Tags: tags})
			if err != nil {
				return err
			}
			assigningTags = append(assigningTags, craetedTags...)
		}

		if err := r.Article().AssignTags(port.AssignTagsParams{
			Article: created,
			Tags:    assigningTags,
		}); err != nil {
			return err
		}

		createdArticle.TagNames = []string{}
		for _, tag := range assigningTags {
			createdArticle.TagNames = append(createdArticle.TagNames, tag.Name)
		}

		return nil
	})
	if err != nil {
		return domain.Article{}, err
	}

	return createdArticle, nil
}

func FilterTagNotExist(existing []domain.Tag, incoming []string) []domain.Tag {
	result := []domain.Tag{}
	existMap := map[string]domain.Tag{}
	for _, tag := range existing {
		existMap[tag.Name] = tag
	}
	for _, tag := range incoming {
		if _, exist := existMap[tag]; exist {
			continue
		}
		result = append(result, domain.Tag{Name: tag})
	}
	return result
}
