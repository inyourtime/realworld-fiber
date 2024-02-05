package uc

import (
	"realworld-go-fiber/core/domain"
	"realworld-go-fiber/core/port"
	"realworld-go-fiber/core/util"
	"realworld-go-fiber/core/util/exception"
	"time"
)

func (uu *userUsecase) Register(arg port.RegisterParams) (domain.User, error) {
	reqUser, err := domain.NewUser(arg.User)
	if err != nil {
		return domain.User{}, err
	}

	user, err := uu.property.repo.User().Create(reqUser)
	if err != nil {
		return domain.User{}, err
	}

	user.Token, _, err = uu.property.tokenMaker.CreateToken(user.ID, 2*time.Hour)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (uu *userUsecase) Login(arg port.LoginParams) (domain.User, error) {
	existing, err := uu.property.repo.User().FilterUser(domain.User{Email: arg.User.Email})
	if err != nil {
		return domain.User{}, err
	}
	if len(existing) < 1 {
		return domain.User{}, exception.Validation().AddError("exception", "email or password invalid")
	}

	user := existing[0]
	if err := util.CheckPassword(arg.User.Password, user.Password); err != nil {
		return domain.User{}, exception.Validation().AddError("exception", "email or password invalid")
	}

	user.Token, _, err = uu.property.tokenMaker.CreateToken(user.ID, 2*time.Hour)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (uu *userUsecase) Current(arg port.AuthParams) (domain.User, error) {
	if arg.Payload == nil {
		return domain.User{}, exception.New(exception.TypePermissionDenied, "token payload not provided", nil)
	}

	existing, err := uu.property.repo.User().FilterUser(domain.User{ID: arg.Payload.UserID})
	if err != nil {
		return domain.User{}, err
	}
	if len(existing) < 1 {
		return domain.User{}, exception.New(exception.TypePermissionDenied, "no user found", nil)
	}

	user := existing[0]
	user.Token = arg.Token

	return user, nil
}
