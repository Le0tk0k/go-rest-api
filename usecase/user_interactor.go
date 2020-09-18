package usecase

import "github.com/Le0tk0k/go-rest-api/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) UserById(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByID(id)
	return
}

func (interactor *UserInteractor) Users() (users domain.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Update(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Update(u)
	return
}

func (interactor *UserInteractor) DeleteById(user domain.User) (err error) {
	err = interactor.UserRepository.DeleteByID(user)
	return
}
