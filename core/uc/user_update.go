package uc

import (
	"realworld-go-fiber/core/domain"
	"realworld-go-fiber/core/port"
	"realworld-go-fiber/core/util/exception"
)

func (uu *userUsecase) Update(arg port.UpdateUserParams) (domain.User, error) {
	if arg.AuthArg.Payload == nil {
		return domain.User{}, exception.New(exception.TypePermissionDenied, "token payload not provided", nil)
	}

	payload := domain.User{
		ID:       arg.AuthArg.Payload.UserID,
		Username: arg.User.Username,
		Email:    arg.User.Email,
		Bio:      arg.User.Bio,
		Image:    arg.User.Image,
	}
	if arg.User.Password != "" {
		if err := payload.SetPassword(arg.User.Password); err != nil {
			return domain.User{}, err
		}
	}

	user, err := uu.property.repo.User().Update(payload)
	if err != nil {
		return domain.User{}, err
	}

	user.Token = arg.AuthArg.Token
	return user, nil
}
