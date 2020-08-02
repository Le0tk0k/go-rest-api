package handler

import (
	"net/http"
	"strconv"

	"github.com/Le0tk0k/go-rest-api/domain/model"
	"github.com/Le0tk0k/go-rest-api/usecase/repository"
	"github.com/labstack/echo"
)

type articleHandler struct {
	articleRepository  repository.ArticleRepository
	categoryRepository repository.CategoryRepository
}

type ArticleHandler interface {
	CreateArticle(c echo.Context)
	GetArticles(c echo.Context)
	GetArticle(c echo.Context)
	UpdateArticle(c echo.Context)
	DeleteArticle(c echo.Context)
}

func NewArticleHandler(aR repository.ArticleRepository, cR repository.CategoryRepository) ArticleHandler {
	return &articleHandler{articleRepository: aR, categoryRepository: cR}
}

func (aH *articleHandler) CreateArticle(c echo.Context) {
	a := &model.Article{}
	if err := c.Bind(a); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := aH.articleRepository.Store(a); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, a)
}

func (aH *articleHandler) GetArticles(c echo.Context) {
	a, err := aH.articleRepository.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, a)
}

func (aH *articleHandler) GetArticle(c echo.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	a, err := aH.articleRepository.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, a)
}

func (aH *articleHandler) UpdateArticle(c echo.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	a := &model.Article{ID: id}

	if err := c.Bind(a); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := aH.articleRepository.Update(a); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, a)
}

func (aH *articleHandler) DeleteArticle(c echo.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := aH.articleRepository.Delete(&model.Article{ID: id}); err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "success")
}
