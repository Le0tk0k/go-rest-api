package datastore

import (
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
	article := model.Article{}
	err := articleRepository.db.First(&article, id).Error
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (articleRepository *articleRepository) Store(article *model.Article) error {
	return articleRepository.db.Save(article).Error
}

func (articleRepository *articleRepository) Update(article *model.Article) error {
	return articleRepository.db.Model(&model.Article{ID: article.ID}).Updates(article).Error
}

func (articleRepository *articleRepository) Delete(article *model.Article) error {
	return articleRepository.db.Delete(article).Error
}

func (articleRepository *articleRepository) FindAll() ([]*model.Article, error) {
	articles := []*model.Article{}

	err := articleRepository.db.Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}
