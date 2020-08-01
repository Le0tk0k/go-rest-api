package datastore

import (
	"fmt"

	"github.com/Le0tk0k/go-rest-api/domain/model"
	"github.com/jinzhu/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

type ArticleRepository interface {
	Store(article *model.Article) error
	FindAll(articles []*model.Article) ([]*model.Article, error)
}

func NewArticleRepository(db *gorm.DB) articleRepository {
	return &articleRepository{db}
}

func (articleRepository *articleRepository) Store(article *model.Article) {
	return articleRepository.db.Save(article).Error
}

func (articleRepository *articleRepository) FindAll(articles []*model.Article) ([]*model.Article, error) {
	err := articleRepository.db.Find(&articles).Error
	if err != nil {
		return nil, fmt.Errorf("SQL Error", err)
	}
	return articles, nil
}
