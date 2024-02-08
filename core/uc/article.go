package uc

import "realworld-go-fiber/core/port"

type articleUsecase struct {
	property usecaseProperty
}

func NewArticleUsecase(property usecaseProperty) port.ArticleUsecase {
	return &articleUsecase{property: property}
}
