package router

import (
	Personnel_controller "github.com/JiratTha/assessment-tax/Personnel/Personnel-controller"
	tax_controller "github.com/JiratTha/assessment-tax/tax/tax-controller"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/personnel", Personnel_controller.GetPersonnel)
	//e.GET("/allowance", Personnel_controller.GetAllowance)
	e.POST("/tax/calculations", tax_controller.TaxCalculationsPost)
}
