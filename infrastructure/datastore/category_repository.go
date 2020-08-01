package datastore

import (
	"fmt"

	"github.com/Le0tk0k/go-rest-api/domain/model"
	"github.com/Le0tk0k/go-rest-api/usecase/repository"
	"github.com/jinzhu/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &categoryRepository{db}
}

func (categoryRepository *categoryRepository) FindByID(id int) (*model.Category, error) {

}

func (categoryRepository *categoryRepository) Store(category *model.Category) error {
	return categoryRepository.db.Save(category).Error
}

func (categoryRepository *categoryRepository) Update(article *model.Category) error {

}

func (categoryRepository *categoryRepository) Delete(article *model.Category) error {

}

func (categoryRepository *categoryRepository) FindAll(categories []*model.Category) ([]*model.Category, error) {
	err := categoryRepository.db.Find(&categories).Error
	if err != nil {
		return nil, fmt.Errorf("SQL Error", err)
	}

	return categories, nil
}
