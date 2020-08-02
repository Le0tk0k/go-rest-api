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
	CreateCategory(c echo.Context)
	GetCategories(c echo.Context)
	GetCategory(c echo.Context)
	UpdateCategory(c echo.Context)
	DeleteCategory(c echo.Context)
}

func NewCategoryHandler(cR repository.CategoryRepository) CategoryHandler {
	return &categoryHandler{categoryRepository: cR}
}

func (cH *categoryHandler) CreateCategory(c echo.Context) {
	cg := &model.Category{}
	if err := c.Bind(cg); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := cH.categoryRepository.Store(cg); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, cg)
}

func (cH *categoryHandler) GetCategories(c echo.Context) {
	cg, err := cH.categoryRepository.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, cg)
}

func (cH *categoryHandler) GetCategory(c echo.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	cg, err := cH.categoryRepository.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, cg)
}

func (cH *categoryHandler) UpdateCategory(c echo.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	cg := &model.Category{ID: id}

	if err := c.Bind(cg); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := cH.categoryRepository.Update(cg); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, cg)
}

func (cH *categoryHandler) DeleteCategory(c echo.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := cH.categoryRepository.Delete(&model.Category{ID: id}); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "success")
}
