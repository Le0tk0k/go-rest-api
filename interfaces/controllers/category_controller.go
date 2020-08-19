package controllers

import (
	"strconv"

	"github.com/Le0tk0k/go-rest-api/domain"
	"github.com/Le0tk0k/go-rest-api/interfaces/database"
	"github.com/Le0tk0k/go-rest-api/usecase"
	"github.com/labstack/echo"
)

type CategoryController struct {
	Interactor usecase.CategoryInteractor
}

func NewCategoryController(sqlHandler database.SqlHandler) *CategoryController {
	return &CategoryController{
		Interactor: usecase.CategoryInteractor{
			CategoryRepository: &database.CategoryRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *CategoryController) CreateCategory(c echo.Context) error {
	cg := domain.Category{}
	c.Bind(&cg)
	category, err := controller.Interactor.Add(cg)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, article)
	return
}

func (controller *CategoryController) GetCategories(c echo.Context) error {
	categories, err := controller.Interactor.Categories()

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, articles)
	return
}

func (controller *CategoryController) GetCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := controller.Interactor.CategoryByID(id)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, articles)
	return
}

func (controller *CategoryController) UpdateCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cg := domain.Category{ID: id}
	c.Bind(&cg)

	category, err := controller.Interactor.Update(cg)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, category)
	return
}

func (controller *CategoryController) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	category := domain.Category{ID: id}

	err := controller.Interactor.DeleteById(category)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, category)
	return
}
