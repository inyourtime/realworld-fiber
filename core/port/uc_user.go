package port

import (
	"realworld-go-fiber/core/domain"
	"realworld-go-fiber/core/util/token"
)

type RegisterParams struct {
	User domain.User
}

type LoginParams struct {
	User domain.User
}

type AuthParams struct {
	Token   string
	Payload *token.Payload
}

type UserUsecase interface {
	Register(arg RegisterParams) (domain.User, error)
	Login(arg LoginParams) (domain.User, error)
	Update()
	Current(arg AuthParams) (domain.User, error)

	Profile()
	Follow()
	UnFollow()
}
