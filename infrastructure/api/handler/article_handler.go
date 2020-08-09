package handler

import (
	"net/http"
	"strconv"

	"github.com/Le0tk0k/go-rest-api/usecase/repository"
	"github.com/labstack/echo"
)

type articleHandler struct {
	articleRepository  repository.ArticleRepository
	categoryRepository repository.CategoryRepository
}

type ArticleHandler interface {
	CreateArticle(c echo.Context) error
	GetArticles(c echo.Context) error
	GetArticle(c echo.Context) error
	UpdateArticle(c echo.Context) error
	DeleteArticle(c echo.Context) error
}

func NewArticleHandler(aR repository.ArticleRepository, cR repository.CategoryRepository) ArticleHandler {
	return &articleHandler{articleRepository: aR, categoryRepository: cR}
}

func (aH *articleHandler) CreateArticle(c echo.Context) error {
	a := &model.Article{}
	if err := c.Bind(a); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	if err := aH.articleRepository.Store(a); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, a)
}

func (aH *articleHandler) GetArticles(c echo.Context) error {
	a, err := aH.articleRepository.FindAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, a)
}

func (aH *articleHandler) GetArticle(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	a, err := aH.articleRepository.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, a)
}

func (aH *articleHandler) UpdateArticle(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	a := &model.Article{ID: id}

	if err := c.Bind(a); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	if err := aH.articleRepository.Update(a); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, a)
}

func (aH *articleHandler) DeleteArticle(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	if err := aH.articleRepository.Delete(&model.Article{ID: id}); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, "success")
}
