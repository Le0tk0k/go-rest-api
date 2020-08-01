package main

import (
	"net/http"

	"github.com/Le0tk0k/go-rest-api/infrastructure/datastore"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	conn := datastore.NewMySqlDB()
	defer conn.Close()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
