package datastore

import (
	"fmt"

	"github.com/Le0tk0k/go-rest-api/domain/model"
	"github.com/jinzhu/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

type CategoryRepository interface {
	Store(category *model.Category) error
	FindAll(categories []*model.Category) ([]*model.Category, error)
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (categoryRepository *categoryRepository) Store(category *model.Category) error {
	return categoryRepository.db.Save(category).Error
}

func (categoryRepository *categoryRepository) FindAll(categories []*model.Category) ([]*model.Category, error) {
	err := categoryRepository.db.Find(&categories).Error
	if err != nil {
		return nil, fmt.Errorf("SQL Error", err)
	}

	return categories, nil
}
