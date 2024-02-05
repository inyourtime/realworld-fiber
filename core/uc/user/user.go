package uc

import (
	"realworld-go-fiber/core/port"
	"realworld-go-fiber/core/util"
	"realworld-go-fiber/core/util/token"
)

type userUsecaseProperty struct {
	config     util.Config
	logger     port.Logger
	tokenMaker token.Maker
	userRepo   port.UserRepository
}

type userUsecase struct {
	property userUsecaseProperty
}

func NewUserUsecase(config util.Config, logger port.Logger, ur port.UserRepository) (port.UserUsecase, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	property := userUsecaseProperty{
		config:     config,
		logger:     logger,
		tokenMaker: tokenMaker,
		userRepo:   ur,
	}

	return &userUsecase{property: property}, nil
}
