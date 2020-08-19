package controllers

import (
	"net/http"
	"strconv"

	"github.com/Le0tk0k/go-rest-api/domain"
	"github.com/Le0tk0k/go-rest-api/interfaces/database"

	"github.com/Le0tk0k/go-rest-api/usecase"

	"github.com/labstack/echo"
)

type ArticleController struct {
	Interactor usecase.ArticleInteractor
}

func NewArticleController(sqlHandler database.SqlHandler) *ArticleController {
	return &ArticleController{
		Interactor: usecase.ArticleInteractor{
			ArticleRepository: &database.ArticleRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ArticleController) CreateArticle(c echo.Context) error {
	a := &domain.Article{}
	if err := c.Bind(a); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	article, err := controller.Interactor.Add(*a)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, article)
}

func (controller *ArticleController) GetArticles(c echo.Context) error {
	articles, err := controller.Interactor.Articles()

	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, articles)
}

func (controller *ArticleController) GetArticle(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	article, err := controller.Interactor.ArticleByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, article)
}

func (controller *ArticleController) UpdateArticle(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	article := &domain.Article{ID: id}

	if err := c.Bind(article); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	if err := controller.Interactor.Update(article); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, article)
}

func (controller *ArticleController) DeleteArticle(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	article := &domain.Article{ID: id}
	if err := controller.Interactor.DeleteById(article); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, article)
}
