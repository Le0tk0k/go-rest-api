package datastore

import (
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
	category := model.Category{}
	err := categoryRepository.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (categoryRepository *categoryRepository) Store(category *model.Category) error {
	return categoryRepository.db.Save(category).Error
}

func (categoryRepository *categoryRepository) Update(category *model.Category) error {
	return categoryRepository.db.Model(&model.Category{ID: category.ID}).Updates(category).Error
}

func (categoryRepository *categoryRepository) Delete(category *model.Category) error {
	return categoryRepository.db.Delete(category).Error
}

func (categoryRepository *categoryRepository) FindAll() ([]*model.Category, error) {
	categories := []*model.Category{}

	err := categoryRepository.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}
