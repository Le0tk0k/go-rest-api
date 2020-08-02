package handler

import (
	"net/http"
	"strconv"

	"github.com/Le0tk0k/go-rest-api/domain/model"
	"github.com/Le0tk0k/go-rest-api/usecase/repository"
	"github.com/labstack/echo/v4"
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
	c.Bind(cg)
	cH.categoryRepository.Store(cg)

	c.JSON(http.StatusOK, cg)
}

func (cH *categoryHandler) GetCategories(c echo.Context) {
	cg, _ := cH.categoryRepository.FindAll()

	c.JSON(http.StatusOK, cg)
}

func (cH *categoryHandler) GetCategory(c echo.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	cg, _ := cH.categoryRepository.FindByID(id)

	c.JSON(http.StatusOK, cg)
}

func (cH *categoryHandler) UpdateCategory(c echo.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	cg := &model.Category{ID: id}

	c.Bind(cg)
	cH.categoryRepository.Update(cg)

	c.JSON(http.StatusOK, cg)
}

func (cH *categoryHandler) DeleteCategory(c echo.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	cH.categoryRepository.Delete(&model.Category{ID: id})
}
