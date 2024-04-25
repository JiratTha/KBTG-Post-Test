package tax_controller

import (
	struc "github.com/JiratTha/assessment-tax/Personnel/model"
	"github.com/JiratTha/assessment-tax/tax/calculation"
	"github.com/labstack/echo/v4"
	"net/http"
)

// TaxCalculationsPost using TaxCalculation and AllowanceCalculation
func TaxCalculationsPost(c echo.Context) error {
	var personal struc.Personnel
	if err := c.Bind(&personal); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	if personal.TotalIncome < 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Total Income must be greater than zero")
	}
	if personal.Wht < 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Wht must be greater than zero")
	}
	allowanceData := personal.Allowance
	for i := range allowanceData {
		if allowanceData[i].Amount < 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Allowance Amount must be greater than zero")
		}
	}
	allowance, _, _ := calculation.AllowanceCalculation(personal)
	personal.TotalIncome -= allowance
	totalTax := calculation.TaxCalculation(personal.TotalIncome, personal.Wht)
	return c.JSON(http.StatusOK, totalTax)
}
