package router

import (
	admin_controller "github.com/JiratTha/assessment-tax/admin/admin-controller"
	"github.com/JiratTha/assessment-tax/middleware"
	tax_controller "github.com/JiratTha/assessment-tax/tax/tax-controller"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/tax/calculations", tax_controller.TaxCalculationsPost)
	e.POST("/tax/calculations/upload-csv", tax_controller.TaxCalculationCSVPost)
	//Admin controller
	g := e.Group("/admin")
	g.Use(middleware.BasicAuthMiddleware)
	g.POST("/deductions/personal", admin_controller.SetPersonalDeductPost)
	g.POST("/deductions/k-receipt", admin_controller.SetKReceiptPost)

}
