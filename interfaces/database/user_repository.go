package database

import (
	"github.com/Le0tk0k/go-rest-api/domain"
)

type UserRepository struct {
	SqlHandler
}

func (userRepository *UserRepository) FindByID(id int) (user domain.User, err error) {
	if err = userRepository.Find(&user, id).Error; err != nil {
		return
	}
	return
}

func (userRepository *UserRepository) Store(u domain.User) (user domain.User, err error) {
	if err = userRepository.Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (userRepository *UserRepository) Update(u domain.User) (user domain.User, err error) {
	if err = userRepository.Save(&u).Error; err != nil {
		return
	}
	user = u
	return
}

func (userRepository *UserRepository) DeleteByID(user domain.User) (err error) {
	if err = userRepository.Delete(&user).Error; err != nil {
		return
	}
	return
}

func (userRepository *UserRepository) FindAll() (users domain.Users, err error) {
	if err = userRepository.Find(&users).Error; err != nil {
		return
	}
	return
}
