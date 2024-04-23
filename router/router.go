package router

import (
	tax_controller "github.com/JiratTha/assessment-tax/tax/tax-controller"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/tax/calculations", tax_controller.TaxCalculationsPost)
}
