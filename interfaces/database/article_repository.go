package database

import (
	"github.com/Le0tk0k/go-rest-api/domain"
)

type ArticleRepository struct {
	SqlHandler
}

func (articleRepository *ArticleRepository) FindByID(id int) (*domain.Article, error) {
	article := domain.Article{}
	err := articleRepository.First(&article, id).Error
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (articleRepository *ArticleRepository) Store(article *domain.Article) error {
	return articleRepository.Save(article).Error
}

func (articleRepository *ArticleRepository) Update(article *domain.Article) error {
	return articleRepository.Model(&domain.Article{ID: article.ID}).Updates(article).Error
}

func (articleRepository *ArticleRepository) Delete(article *domain.Article) error {
	return articleRepository.Delete(article).Error
}

func (articleRepository *ArticleRepository) FindAll() ([]*domain.Article, error) {
	articles := []*domain.Article{}

	err := articleRepository.Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}
