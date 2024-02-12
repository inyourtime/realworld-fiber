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

func (r *articleRepo) FilterArticles(condition interface{}) ([]domain.Article, error) {
	articles := []model.Article{}
	err := r.db.
		Preload("Author", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Username", "Bio", "Image")
		}).
		Where(condition).Find(&articles).Error
	if err != nil {
		return []domain.Article{}, err
	}
	result := []domain.Article{}
	for _, article := range articles {
		result = append(result, article.ToDomain())
	}
	return result, nil
}

func (r *articleRepo) CreateArticle(arg domain.Article) (domain.Article, error) {
	article := model.AsArticle(arg)
	err := r.db.Create(&article).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return domain.Article{}, exception.New(exception.TypeValidation, "Slug is already existing", err)
		}
		return domain.Article{}, err
	}
	return article.ToDomain(), nil
}

func (r *articleRepo) FilterTags(condition interface{}) ([]domain.Tag, error) {
	tags := []model.Tag{}
	result := []domain.Tag{}
	err := r.db.Where(condition).Find(&tags).Error
	if err != nil {
		return []domain.Tag{}, err
	}
	for _, tag := range tags {
		result = append(result, tag.ToDomain())
	}
	return result, nil
}

func (r *articleRepo) AddTags(arg port.AddTagsPayload) ([]domain.Tag, error) {
	tags := []model.Tag{}
	for _, tag := range arg.Tags {
		tags = append(tags, model.AsTag(tag))
	}
	err := r.db.Create(&tags).Error
	if err != nil {
		return []domain.Tag{}, err
	}

	result := []domain.Tag{}
	for _, tag := range tags {
		result = append(result, tag.ToDomain())
	}
	return result, nil
}

func (r *articleRepo) AssignTags(arg port.AssignTagsParams) error {
	article := model.AsArticle(arg.Article)
	tags := []model.Tag{}
	for _, tag := range arg.Tags {
		tags = append(tags, model.AsTag(tag))
	}
	return r.db.Model(&article).Association("Tags").Append(&tags)
}
