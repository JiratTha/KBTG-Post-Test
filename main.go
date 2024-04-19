package main

import (
	"github.com/JiratTha/assessment-tax/router"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.InitDBRoutes(e)
	router.InitRoutes(e)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
