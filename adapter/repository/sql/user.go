package sql

import (
	"realworld-go-fiber/adapter/repository/sql/model"
	"realworld-go-fiber/core/domain"
	"realworld-go-fiber/core/port"
	"realworld-go-fiber/core/util/exception"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(arg domain.User) (domain.User, error) {
	user := model.AsUser(arg)
	err := r.db.Create(&user).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return domain.User{}, exception.New(exception.TypeValidation, "Email or Username are already existing", err)
		}
		return domain.User{}, err
	}
	return user.ToDomain(), nil
}

func (r *userRepo) Update(arg domain.User) (domain.User, error) {
	user := model.AsUser(arg)
	err := r.db.Save(&user).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return domain.User{}, exception.New(exception.TypeValidation, "Email or Username are already existing", err)
		}
		return domain.User{}, err
	}
	return user.ToDomain(), nil
}

func (r *userRepo) FindOne(condition domain.User) (domain.User, error) {
	user := model.User{}
	cond := model.AsUser(condition)
	err := r.db.Where(cond).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, exception.New(exception.TypeNotFound, "User not found", err)
		}
		return domain.User{}, err
	}
	return user.ToDomain(), nil
}

func (r *userRepo) Follow(arg1 domain.User, arg2 domain.User) error {
	a := model.AsUser(arg1)
	b := model.AsUser(arg2)
	return r.db.Model(a).Association("Followings").Append(b)
}

func (r *userRepo) UnFollow(arg1 domain.User, arg2 domain.User) error {
	a := model.AsUser(arg1)
	b := model.AsUser(arg2)
	return r.db.Model(a).Association("Followings").Delete(b)
}
