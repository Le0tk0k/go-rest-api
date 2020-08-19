package usecase

import "github.com/Le0tk0k/go-rest-api/domain"

type CategoryInteractor struct {
	CategoryRepository CategoryRepository
}

func (interactor *CategoryInteractor) CategoryById(id int) (category *domain.Category, err error) {
	category, err = interactor.CategoryRepository.FindByID(id)
	return
}

func (interactor *CategoryInteractor) Categories() (categories *domain.Category, err error) {
	categories, err = interactor.CategoryRepository.FindAll()
	return
}

func (interactor *CategoryInteractor) Add(category domain.Category) (category *domain.Category, err error) {
	category, err = interactor.CategoryRepository.Store(category)
	return
}

func (interactor *CategoryInteractor) Update(category domain.Category) (category *domain.Category, err error) {
	category, err = interactor.CategoryRepository.Update(category)
	return
}

func (interactor *CategoryInteractor) DeleteById(category domain.Category) (err error) {
	err = interactor.CategoryRepository.DeleteById(category)
	return
}
