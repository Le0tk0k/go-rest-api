package usecase

import "github.com/Le0tk0k/go-rest-api/domain"

type CategoryRepository interface {
	FindByID(id int) (*domain.Category, error)
	Store(category *domain.Category) error
	Update(category *domain.Category) error
	Delete(category *domain.Category) error
	FindAll() ([]*domain.Category, error)
}
