package main

import (
	"github.com/Le0tk0k/go-rest-api/infrastructure/api/handler"
	"github.com/Le0tk0k/go-rest-api/infrastructure/datastore"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	conn := datastore.NewMySqlDb()
	defer conn.Close()
	articleRepo := datastore.NewArticleRepository(conn)
	categoryRepo := datastore.NewCategoryRepository(conn)

	aH := handler.NewArticleHandler(articleRepo, categoryRepo)
	cH := handler.NewCategoryHandler(categoryRepo)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/articles", aH.GetArticles)
	e.GET("/articles/:id", aH.GetArticle)
	e.POST("/articles", aH.CreateArticle)
	e.PUT("/articles/:id", aH.UpdateArticle)
	e.DELETE("/articles/:id", aH.DeleteArticle)

	e.GET("/categories", cH.GetCategories)
	e.GET("/categories/:id", cH.GetCategory)
	e.POST("/categories", cH.CreateCategory)
	e.PUT("/categories/:id", cH.UpdateCategory)
	e.DELETE("/categories/:id", cH.DeleteCategory)

	e.Logger.Fatal(e.Start(":8080"))
}
