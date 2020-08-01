package datastore

import (
	"fmt"

	"github.com/Le0tk0k/go-rest-api/domain/model"
	"github.com/Le0tk0k/go-rest-api/usecase/repository"
	"github.com/jinzhu/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) repository.ArticleRepository {
	return &articleRepository{db}
}

func (articleRepository *articleRepository) FindByID(id int) (*model.Article, error) {

}

func (articleRepository *articleRepository) Store(article *model.Article) error {
	return articleRepository.db.Save(article).Error
}

func (articleRepository *articleRepository) Update(article *model.Article) error {

}

func (articleRepository *articleRepository) Delete(article *model.Article) error {

}

func (articleRepository *articleRepository) FindAll(articles []*model.Article) ([]*model.Article, error) {
	err := articleRepository.db.Find(&articles).Error
	if err != nil {
		return nil, fmt.Errorf("SQL Error", err)
	}
	return articles, nil
}
