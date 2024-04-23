package tax_controller

import (
	struc "github.com/JiratTha/assessment-tax/Personnel/Personnel-model"
	"github.com/JiratTha/assessment-tax/tax/calculation"
	"github.com/labstack/echo/v4"
	"net/http"
)

func TaxCalculationsPost(c echo.Context) error {
	var personnelIncome struc.Personnel
	if err := c.Bind(&personnelIncome); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	allowance, _, _ := calculation.AllowanceCalculation(personnelIncome)
	personnelIncome.TotalIncome -= allowance
	totalTax := calculation.TaxCalculation(personnelIncome.TotalIncome, personnelIncome.Wht)
	return c.JSON(http.StatusOK, totalTax)
}
