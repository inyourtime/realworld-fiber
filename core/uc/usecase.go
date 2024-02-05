package uc

import (
	"realworld-go-fiber/core/port"
	"realworld-go-fiber/core/util"
	"realworld-go-fiber/core/util/token"
)

type usecaseProperty struct {
	config     util.Config
	tokenMaker token.Maker
	repo       port.Repository
	logger     port.Logger
}

type usecases struct {
	property       usecaseProperty
	userUsecase    port.UserUsecase
	articleUsecase port.ArticleUsecase
}

func NewUsecase(config util.Config, repo port.Repository, logger port.Logger) (port.Usecase, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}
	property := usecaseProperty{
		config:     config,
		repo:       repo,
		tokenMaker: tokenMaker,
		logger:     logger,
	}
	uc := usecases{
		property:    property,
		userUsecase: NewUserUsecase(property),
	}
	return &uc, nil
}

func (u *usecases) TokenMaker() token.Maker {
	return u.property.tokenMaker
}

func (u *usecases) Article() port.ArticleUsecase {
	return u.articleUsecase
}

func (u *usecases) User() port.UserUsecase {
	return u.userUsecase
}
