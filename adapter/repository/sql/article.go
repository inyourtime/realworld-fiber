package sql

import (
	"realworld-go-fiber/adapter/repository/sql/model"
	"realworld-go-fiber/core/domain"
	"realworld-go-fiber/core/port"
	"realworld-go-fiber/core/util/exception"

	"gorm.io/gorm"
)

type articleRepo struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) port.ArticleRepository {
	return &articleRepo{db: db}
}

// func (r *articleRepo) FilterTags(condition interface{}) ([]domain.Tag, error) {
// 	tags := []model.Tag{}
// 	err := r.db.Where(condition).Find(&tags).Error
// 	if err != nil {
// 		return []domain.Tag{}, err
// 	}
// 	result := []domain.Tag{}
// 	for _, tag := range tags {
// 		result = append(result, tag.ToDomain())
// 	}
// 	return result, nil
// }

func (r *articleRepo) CreateArticle(tx *gorm.DB, arg *model.Article) error {
	if err := tx.Create(arg).Error; err != nil {
		if err == gorm.ErrDuplicatedKey {
			return exception.New(exception.TypeValidation, "Slug is already existing", err)
		}
		return err
	}
	return nil
}

func (r *articleRepo) FilterTags(tx *gorm.DB, condition interface{}, result *[]model.Tag) error {
	return tx.Where(condition).Find(result).Error
}

func (r *articleRepo) AddTags(tx *gorm.DB, arg *[]model.Tag) error {
	return tx.Create(arg).Error
}

func (r *articleRepo) AssignTags(tx *gorm.DB, article *model.Article, tags *[]model.Tag) error {
	return tx.Model(article).Association("Tags").Append(tags)
}

func (r *articleRepo) CreateArticleTransaction(arg domain.Article) (domain.Article, []domain.Tag, error) {
	article := model.AsArticle(arg)
	tags := []model.Tag{}
	existingTags := []model.Tag{}
	assigningTags := []model.Tag{}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.CreateArticle(tx, &article); err != nil {
			return err
		}
		if len(arg.TagNames) == 0 {
			return nil
		}

		if err := r.FilterTags(tx, map[string]interface{}{"name": arg.TagNames}, &existingTags); err != nil {
			return err
		}

		model.FilterTagNotExist(existingTags, arg.TagNames, &tags)

		if len(tags) > 0 {
			if err := r.AddTags(tx, &tags); err != nil {
				return err
			}
		}

		assigningTags = append(existingTags, tags...)
		if err := r.AssignTags(tx, &article, &assigningTags); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return domain.Article{}, []domain.Tag{}, err
	}

	resultTags := []domain.Tag{}
	for _, tag := range assigningTags {
		resultTags = append(resultTags, tag.ToDomain())
	}

	return article.ToDomain(), resultTags, nil
}
