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

type UpdateUserParams struct {
	User    domain.User
	AuthArg AuthParams
}

type ProfileParams struct {
	AuthArg  AuthParams
	Username string
}

type UserUsecase interface {
	Register(arg RegisterParams) (domain.User, error)
	Login(arg LoginParams) (domain.User, error)
	Update(arg UpdateUserParams) (domain.User, error)
	Current(arg AuthParams) (domain.User, error)

	Profile(arg ProfileParams) (domain.User, error)
	Follow(arg ProfileParams) (domain.User, error)
	UnFollow(arg ProfileParams) (domain.User, error)
}
