package infrastructure

import (
	"github.com/Le0tk0k/go-rest-api/interfaces/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	e := echo.New()

	articleController := controllers.NewArticleController(NewSqlHandler())
	categoryController := controllers.NewCategoryController(NewSqlHandler())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/articles", func(c echo.Context) error { return articleController.GetArticles(c) })
	e.GET("/articles/:id", func(c echo.Context) error { return articleController.GetArticle(c) })
	e.POST("/articles", func(c echo.Context) error { return articleController.CreateArticle(c) })
	e.PUT("/articles/:id", func(c echo.Context) error { return articleController.UpdateArticle(c) })
	e.DELETE("/articles/:id", func(c echo.Context) error { return articleController.DeleteArticle(c) })

	e.GET("/categories", func(c echo.Context) error { return categoryController.GetCategories(c) })
	e.GET("/categories/:id", func(c echo.Context) error { return categoryController.GetCategory(c) })
	e.POST("/categories", func(c echo.Context) error { return categoryController.CreateCategory(c) })
	e.PUT("/categories/:id", func(c echo.Context) error { return categoryController.UpdateCategory(c) })
	e.DELETE("/categories/:id", func(c echo.Context) error { return categoryController.DeleteCategory(c) })

	e.Logger.Fatal(e.Start(":8080"))
}
