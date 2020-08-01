package repository

import "github.com/Le0tk0k/go-rest-api/domain/model"

type ArticleRepository interface {
	FindByID(id int) (*model.Article, error)
	Store(article *model.Article) error
	Update(article *model.Article) error
	Delete(article *model.Article) error
	FindAll() ([]*model.Article, error)
}
