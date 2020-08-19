package controllers

import (
	"net/http"
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
	cg := &domain.Category{}
	if err := c.Bind(cg); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	category, err := controller.Interactor.Add(*cg)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, category)
}

func (controller *CategoryController) GetCategories(c echo.Context) error {
	categories, err := controller.Interactor.Categories()

	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, categories)
}

func (controller *CategoryController) GetCategory(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	category, err := controller.Interactor.CategoryByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, category)
}

func (controller *CategoryController) UpdateCategory(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	category := &domain.Category{ID: id}

	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	if err := controller.Interactor.Update(category); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, category)
}

func (controller *CategoryController) DeleteCategory(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	category := &domain.Category{ID: id}
	if err := controller.Interactor.DeleteById(cateogyr); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, category)
}
