package usecase

import "github.com/Le0tk0k/go-rest-api/domain"

type CategoryRepository interface {
	FindByID(id int) (domain.Category, error)
	Store(domain.Category) (domain.Category, error)
	Update(domain.Category) (domain.Category, error)
	Delete(domain.Category) error
	FindAll() (domain.Categories, error)
}
