package main

import (
	Personnel_model "github.com/JiratTha/assessment-tax/Personnel/Personnel-model"
	"github.com/JiratTha/assessment-tax/db"
	"github.com/JiratTha/assessment-tax/router"
	"github.com/labstack/echo/v4"
	echoswagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	// Use environment variables for database connection
	DatabaseUrl := os.Getenv("DatabaseUrl")
	if DatabaseUrl == "" {
		DatabaseUrl = "host=localhost port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable"
		// Initialize the database
		err := db.InitDB(DatabaseUrl)
		if err != nil {
			e.Logger.Fatal("Error fetching allowance:", err)
		}
		// Fetch allowance data after initializing the database
		allowance, err := Personnel_model.GetAllowanceData()
		if err != nil {
			e.Logger.Fatal("Error fetching allowance:", allowance)
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

	router.InitRoutes(e)

	e.GET("/swagger/*", echoswagger.WrapHandler)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy")
	})

	e.Logger.Fatal(e.Start(":" + PORT))
}
