package port

import (
	"realworld-go-fiber/core/domain"
)

type UserRepository interface {
	Create(m domain.User) (domain.User, error)
	Update(m domain.User) (domain.User, error)
	FilterUser(condition interface{}) ([]domain.User, error)
	FindOne(condition interface{}) (domain.User, error)
	Follow(a domain.User, b domain.User) error
	UnFollow(a domain.User, b domain.User) error
}
