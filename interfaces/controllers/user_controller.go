package controllers

import (
	"strconv"

	"github.com/Le0tk0k/go-rest-api/domain"
	"github.com/Le0tk0k/go-rest-api/interfaces/database"
	"github.com/Le0tk0k/go-rest-api/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) CreateUser(c Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

func (controller *UserController) GetUsers(c Context) (err error) {
	users, err := controller.Interactor.Users()

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

func (controller *UserController) GetUser(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

func (controller *UserController) UpdateUser(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	u := domain.User{ID: id}
	c.Bind(&u)

	user, err := controller.Interactor.Update(u)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

func (controller *UserController) DeleteUser(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := domain.User{ID: id}

	err = controller.Interactor.DeleteById(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}
