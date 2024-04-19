package router

import (
	Personnel_controller "github.com/JiratTha/assessment-tax/Personnel/Personnel-controller"
	Personnel_model "github.com/JiratTha/assessment-tax/Personnel/Personnel-model"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/personnel", Personnel_controller.GetPersonnel)
	e.GET("/allowance", Personnel_controller.GetAllowance)
}

func InitDBRoutes(e *echo.Echo) {
	err := Personnel_model.InitDB("host=localhost port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable")
	if err != nil {
		e.Logger.Fatal("Error initializing database:", err)
	}

	// Now that the database connection is initialized, you can use other functions from the Personnel_model package
	allowance, err := Personnel_model.GetAllowanceData()
	if err != nil {
		e.Logger.Fatal("Error fetching allowance:", allowance)
	}

	// Use allowance...
}
