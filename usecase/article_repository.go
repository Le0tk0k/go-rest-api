package usecase

import "github.com/Le0tk0k/go-rest-api/domain"

type ArticleRepository interface {
	FindByID(id int) (*domain.Article, error)
	Store(article *domain.Article) error
	Update(article *domain.Article) error
	Delete(article *domain.Article) error
	FindAll() ([]*domain.Article, error)
}
