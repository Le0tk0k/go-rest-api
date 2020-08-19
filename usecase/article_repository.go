package usecase

import "github.com/Le0tk0k/go-rest-api/domain"

type ArticleRepository interface {
	FindByID(id int) (domain.Article, error)
	Store(domain.Article) (domain.Article, error)
	Update(domain.Article) (domain.Article, error)
	DeleteByID(domain.Article) error
	FindAll() (domain.Articles, error)
}
