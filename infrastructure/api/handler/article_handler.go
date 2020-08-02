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
	c.Bind(a)
	aH.articleRepository.Store(a)

	c.JSON(http.StatusOK, a)
}

func (aH *articleHandler) GetArticles(c echo.Context) {
	a, _ := aH.articleRepository.FindAll()

	c.JSON(http.StatusOK, a)
}

func (aH *articleHandler) GetArticle(c echo.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	a, _ := aH.articleRepository.FindByID(id)

	c.JSON(http.StatusOK, a)
}

func (aH *articleHandler) UpdateArticle(c echo.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	a := &model.Article{ID: id}

	c.Bind(a)
	aH.articleRepository.Update(a)

	c.JSON(http.StatusOK, a)
}

func (aH *articleHandler) DeleteArticle(c echo.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	aH.articleRepository.Delete(&model.Article{ID: id})
}
