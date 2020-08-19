package datasbase

import (
	"github.com/Le0tk0k/go-rest-api/domain"
	"github.com/Le0tk0k/go-rest-api/usecase/repository"
	"github.com/jinzhu/gorm"
)

type categoryRepository struct {
	SqlHandler
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &categoryRepository{db}
}

func (categoryRepository *categoryRepository) FindByID(id int) (*domain.Category, error) {
	category := domain.Category{}
	err := categoryRepository.First(&category, id).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (categoryRepository *categoryRepository) Store(category *domain.Category) error {
	return categoryRepository.Save(category).Error
}

func (categoryRepository *categoryRepository) Update(category *domain.Category) error {
	return categoryRepository.Model(&domain.Category{ID: category.ID}).Updates(category).Error
}

func (categoryRepository *categoryRepository) Delete(category *domain.Category) error {
	return categoryRepository.Delete(category).Error
}

func (categoryRepository *categoryRepository) FindAll() ([]*domain.Category, error) {
	categories := []*domain.Category{}

	err := categoryRepository.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}
