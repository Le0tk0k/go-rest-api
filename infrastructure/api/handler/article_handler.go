package handler

import (
	"github.com/Le0tk0k/go-rest-api/infrastructure/datastore"
	"github.com/labstack/echo"
)

type articleHandler struct {
	articleRepository datastore.ArticleRepository
}

type ArticleHandler interface {
	CreateArticle(c echo.Context) error
	GetArticles(c echo.Context) error
}

func NewArticleHandler()
