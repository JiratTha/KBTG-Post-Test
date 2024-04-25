package router

import (
	admin_controller "github.com/JiratTha/assessment-tax/admin/admin-controller"
	tax_controller "github.com/JiratTha/assessment-tax/tax/tax-controller"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/tax/calculations", tax_controller.TaxCalculationsPost)
	e.POST("/admin/deductions/personal", admin_controller.SetPersonnelDeductPost)
}
