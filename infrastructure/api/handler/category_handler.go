package handler

import (
	"net/http"
	"strconv"

	"github.com/Le0tk0k/go-rest-api/domain/model"
	"github.com/Le0tk0k/go-rest-api/usecase/repository"
	"github.com/labstack/echo"
)

type categoryHandler struct {
	categoryRepository repository.CategoryRepository
}

type CategoryHandler interface {
	CreateCategory(c echo.Context) error
	GetCategories(c echo.Context) error
	GetCategory(c echo.Context) error
	UpdateCategory(c echo.Context) error
	DeleteCategory(c echo.Context) error
}

func NewCategoryHandler(cR repository.CategoryRepository) CategoryHandler {
	return &categoryHandler{categoryRepository: cR}
}

func (cH *categoryHandler) CreateCategory(c echo.Context) error {
	cg := &model.Category{}
	if err := c.Bind(cg); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	if err := cH.categoryRepository.Store(cg); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	c.JSON(http.StatusOK, cg)
}

func (cH *categoryHandler) GetCategories(c echo.Context) error {
	cg, err := cH.categoryRepository.FindAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	c.JSON(http.StatusOK, cg)
}

func (cH *categoryHandler) GetCategory(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	cg, err := cH.categoryRepository.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	c.JSON(http.StatusOK, cg)
}

func (cH *categoryHandler) UpdateCategory(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	cg := &model.Category{ID: id}

	if err := c.Bind(cg); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	if err := cH.categoryRepository.Update(cg); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	c.JSON(http.StatusOK, cg)
}

func (cH *categoryHandler) DeleteCategory(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	if err := cH.categoryRepository.Delete(&model.Category{ID: id}); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	c.JSON(http.StatusCreated, "success")
}
