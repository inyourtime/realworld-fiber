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

func (r *articleRepo) FilterTags(condition interface{}) ([]domain.Tag, error) {
	tags := []model.Tag{}
	err := r.db.Where(condition).Find(&tags).Error
	if err != nil {
		return []domain.Tag{}, err
	}
	result := []domain.Tag{}
	for _, tag := range tags {
		result = append(result, tag.ToDomain())
	}
	return result, nil
}

func (r *articleRepo) CreateArticleTransaction(arg domain.Article) (domain.Article, []domain.Tag, error) {
	article := model.AsArticle(arg)
	tags := []model.Tag{}
	existingTags := []model.Tag{}
	assigningTags := []model.Tag{}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&article).Error; err != nil {
			if err == gorm.ErrDuplicatedKey {
				return exception.New(exception.TypeValidation, "Slug is already existing", err)
			}
			return err
		}
		if len(arg.TagNames) == 0 {
			return nil
		}

		if err := tx.Where(map[string]interface{}{"name": arg.TagNames}).Find(&existingTags).Error; err != nil {
			return err
		}

		existMap := map[string]model.Tag{}
		for _, tag := range existingTags {
			existMap[tag.Name] = tag
		}

		for _, tag := range arg.TagNames {
			if _, exist := existMap[tag]; exist {
				continue
			}
			tags = append(tags, model.Tag{Name: tag})
		}
		if len(tags) > 0 {
			if err := tx.Create(&tags).Error; err != nil {
				return err
			}
		}

		assigningTags = append(existingTags, tags...)
		if err := tx.Model(&article).Association("Tags").Append(&assigningTags); err != nil {
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
