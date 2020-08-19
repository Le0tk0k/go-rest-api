package controllers

import (
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

func (controller *ArticleController) CreateArticle(c echo.Context) (err error) {
	a := domain.Article{}
	c.Bind(&a)
	article, err := controller.Interactor.Add(a)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, article)
	return
}

func (controller *ArticleController) GetArticles(c echo.Context) (err error) {
	articles, err := controller.Interactor.Articles()

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, articles)
	return
}

func (controller *ArticleController) GetArticle(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := controller.Interactor.ArticleById(id)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, article)
	return
}

func (controller *ArticleController) UpdateArticle(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	a := domain.Article{ID: id}
	c.Bind(&a)

	article, err := controller.Interactor.Update(a)

	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, article)
	return
}

func (controller *ArticleController) DeleteArticle(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	article := domain.Article{ID: id}

	err := controller.Interactor.DeleteById(article)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, article)
	return
}
