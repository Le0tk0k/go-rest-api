package database

import (
	"github.com/Le0tk0k/go-rest-api/domain"
)

type CategoryRepository struct {
	SqlHandler
}

func (categoryRepository *CategoryRepository) FindByID(id int) (category domain.Category, err error) {
	if err = categoryRepository.Find(&category, id).Error; err != nil {
		return
	}
	return
}

func (categoryRepository *CategoryRepository) Store(cg domain.Category) (category domain.Category, err error) {
	if err = categoryRepository.Create(&cg).Error; err != nil {
		return
	}
	category = cg
	return
}

func (categoryRepository *CategoryRepository) Update(cg domain.Category) (category domain.Category, err error) {
	if err = categoryRepository.Save(&cg).Error; err != nil {
		return
	}
	category = cg
	return
}

func (categoryRepository *CategoryRepository) DeleteByID(category domain.Category) (err error) {
	if err = categoryRepository.Delete(&category).Error; err != nil {
		return
	}
	return
}

func (categoryRepository *CategoryRepository) FindAll() (categories domain.Categories, err error) {
	if err = categoryRepository.Find(&categories).Error; err != nil {
		return
	}
	return
}
