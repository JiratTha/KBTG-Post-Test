package main

import (
	"github.com/JiratTha/assessment-tax/db"
	_ "github.com/JiratTha/assessment-tax/docs"
	"github.com/JiratTha/assessment-tax/router"
	"github.com/JiratTha/assessment-tax/util"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoswagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
)

func main() {
	e := echo.New()

	DatabaseUrl := os.Getenv("DatabaseUrl")
	if DatabaseUrl == "" {
		DatabaseUrl = "host=localhost port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable"
		err := db.InitDB(DatabaseUrl)
		if err != nil {
			e.Logger.Fatal("Error fetching allowance:", err)
		}
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080" // Default port if not specified
	}

	AdminUsername := os.Getenv("AdminUsername")
	if AdminUsername == "" {
		AdminUsername = "adminTax" // Default port if not specified
	}

	AdminPassword := os.Getenv("AdminPassword")
	if AdminPassword == "" {
		AdminPassword = "admin!" // Default port if not specified
	}
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	router.InitRoutes(e)
	e.GET("/swagger/*", echoswagger.WrapHandler)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy")
	})

	e.Logger.Fatal(e.Start(":" + PORT))

}
