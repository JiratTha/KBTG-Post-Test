package tax_controller

import (
	"github.com/JiratTha/assessment-tax/tax/allowance-calculation"
	struc "github.com/JiratTha/assessment-tax/tax/personal"
	"github.com/JiratTha/assessment-tax/tax/tax-calculation"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// TaxCalculationsPost handles the POST /tax/calculations route
// @Summary Calculate taxes
// @Description Calculates taxes based on total income, withholding tax, and allowances.
// @Tags tax
// @Accept  json
// @Produce  json
// @Success 200 {object} model.TaxResponse  "Returns the calculated tax amount"
// @Failure 400 {string} string "Invalid input"
// @Router /tax/calculations [post]
func TaxCalculationsPost(c echo.Context) error {
	var personal struc.Personnel
	if err := c.Bind(&personal); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	err := c.Validate(personal)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
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
	allowance, _, _ := allowance_calculation.AllowanceCalculation(personal)
	personal.TotalIncome -= allowance
	totalTax := tax_calculation.TaxCalculation(personal.TotalIncome, personal.Wht)
	return c.JSON(http.StatusOK, totalTax)
}
