package repository

import "github.com/Le0tk0k/go-rest-api/domain/model"

type CategoryRepository interface {
	FindByID(id int) (*model.Category, error)
	Store(category *model.Category) error
	Update(category *model.Category) error
	Delete(category *model.Category) error
	FindAll() ([]*model.Category, error)
}
