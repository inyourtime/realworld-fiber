package port

import "realworld-go-fiber/core/util/token"

type Usecase interface {
	TokenMaker() token.Maker
	User() UserUsecase
	Article() ArticleUsecase
}
