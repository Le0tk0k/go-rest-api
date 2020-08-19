package database

import (
	"github.com/Le0tk0k/go-rest-api/domain"
)

type ArticleRepository struct {
	SqlHandler
}

func (articleRepository *ArticleRepository) FindByID(id int) (article domain.Article, err error) {
	if err = articleRepository.Find(&article, id).Error; err != nil {
		return
	}
	return
}

func (articleRepository *ArticleRepository) Store(a domain.Article) (article domain.Article, err error) {
	if err = articleRepository.Create(&a).Error; err != nil {
		return
	}
	article = a
	return
}

func (articleRepository *ArticleRepository) Update(a domain.Article) (article domain.Article, err error) {
	if err = articleRepository.Save(&a).Error; err != nil {
		return
	}
	article = a
	return
}

func (articleRepository *ArticleRepository) Delete(article domain.Article) (err error) {
	if err = articleRepository.Delete(&article).Error; err != nil {
		return
	}
	return
}

func (articleRepository *ArticleRepository) FindAll() (articles domain.Articles, err error) {
	if err = articleRepository.Find(&articles).Error; err != nil {
		return
	}
}
