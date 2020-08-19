package database

import (
	"github.com/Le0tk0k/go-rest-api/domain"
)

type CategoryRepository struct {
	SqlHandler
}

func (categoryRepository *CategoryRepository) FindByID(id int) (*domain.Category, error) {
	category := domain.Category{}
	err := categoryRepository.First(&category, id).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (categoryRepository *CategoryRepository) Store(category *domain.Category) error {
	return categoryRepository.Save(category).Error
}

func (categoryRepository *CategoryRepository) Update(category *domain.Category) error {
	return categoryRepository.Model(&domain.Category{ID: category.ID}).Updates(category).Error
}

func (categoryRepository *CategoryRepository) Delete(category *domain.Category) error {
	return categoryRepository.Delete(category).Error
}

func (categoryRepository *CategoryRepository) FindAll() ([]*domain.Category, error) {
	categories := []*domain.Category{}

	err := categoryRepository.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}
