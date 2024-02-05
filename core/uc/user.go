package uc

import (
	"realworld-go-fiber/core/port"
)

type userUsecase struct {
	property usecaseProperty
}

func NewUserUsecase(property usecaseProperty) port.UserUsecase {
	return &userUsecase{property: property}
}
