package uc

import (
	"realworld-go-fiber/core/domain"
	"realworld-go-fiber/core/port"
)

func (uu *userUsecase) Profile(arg port.ProfileParams) (domain.User, error) {
	return domain.User{}, nil
}

func (uu *userUsecase) Follow(arg port.ProfileParams) (domain.User, error) {

	return domain.User{}, nil
}

func (uu *userUsecase) UnFollow(arg port.ProfileParams) (domain.User, error) {

	return domain.User{}, nil
}
