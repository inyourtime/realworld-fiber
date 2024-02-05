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
	err := r.db.Model(&user).Updates(user).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return domain.User{}, exception.New(exception.TypeValidation, "Email or Username are already existing", err)
		}
		return domain.User{}, err
	}

	updated, err := r.FindOne(domain.User{ID: user.ID})
	if err != nil {
		return domain.User{}, err
	}

	return updated, nil
}

/**
 * FilterUser is a method of the userRepo struct that filters users based on a given condition.
 *
 * Parameters:
 * - condition: The condition to filter users by.
 * 		example: domain.User{Email: "some@mail.co"} or map[string]interface{}{"email": "some@mail.co"}
 *
 * Returns:
 * - []domain.User: A slice of domain.User objects that match the condition.
 * - error: An error if the filtering process encounters any issues.
 */
func (r *userRepo) FilterUser(condition interface{}) ([]domain.User, error) {
	users := []model.User{}
	err := r.db.Where(condition).Find(&users).Error
	if err != nil {
		return []domain.User{}, err
	}
	result := []domain.User{}
	for _, user := range users {
		result = append(result, user.ToDomain())
	}
	return result, nil
}

func (r *userRepo) FindOne(condition interface{}) (domain.User, error) {
	users, err := r.FilterUser(condition)
	if err != nil {
		return domain.User{}, err
	}
	if len(users) == 0 {
		return domain.User{}, exception.New(exception.TypeNotFound, "User not found", nil)
	}
	return users[0], nil
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
