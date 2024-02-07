package uc

import (
	"realworld-go-fiber/core/domain"
	"realworld-go-fiber/core/port"
	"realworld-go-fiber/core/util/exception"
)

func (uu *userUsecase) Profile(arg port.ProfileParams) (domain.User, error) {
	user, err := uu.property.repo.User().FindOne(domain.User{Username: arg.Username})
	if err != nil {
		return domain.User{}, err
	}

	if arg.AuthArg.Payload == nil {
		return user, nil
	}

	follows, err := uu.property.repo.User().FilterFollow(domain.User{ID: arg.AuthArg.Payload.UserID}, user)
	if err != nil {
		return domain.User{}, err
	}

	if len(follows) > 0 {
		user.IsFollowed = true
	}

	return user, nil
}

func (uu *userUsecase) Follow(arg port.ProfileParams) (domain.User, error) {
	if arg.AuthArg.Payload == nil {
		return domain.User{}, exception.New(exception.TypePermissionDenied, "authentication required", nil)
	}

	user, err := uu.property.repo.User().FindOne(domain.User{Username: arg.Username})
	if err != nil {
		return domain.User{}, err
	}

	err = uu.property.repo.User().Follow(domain.User{ID: arg.AuthArg.Payload.UserID}, user)
	if err != nil {
		return domain.User{}, err
	}

	user.IsFollowed = true
	return user, nil
}

func (uu *userUsecase) UnFollow(arg port.ProfileParams) (domain.User, error) {
	if arg.AuthArg.Payload == nil {
		return domain.User{}, exception.New(exception.TypePermissionDenied, "authentication required", nil)
	}

	user, err := uu.property.repo.User().FindOne(domain.User{Username: arg.Username})
	if err != nil {
		return domain.User{}, err
	}

	err = uu.property.repo.User().UnFollow(domain.User{ID: arg.AuthArg.Payload.UserID}, user)
	if err != nil {
		return domain.User{}, err
	}

	user.IsFollowed = false
	return user, nil
}
