package Personnel_controller

import (
	Personnel_model "github.com/JiratTha/assessment-tax/Personnel/Personnel-model"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetPersonnel handles the GET users route
func GetPersonnel(c echo.Context) error {
	// Call the GetPersonnel function from the Personnel_model package
	personnel, err := Personnel_model.GetPersonnelData(Personnel_model.Personnel{
		TotalIncome: 0,
		Wht:         0,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching personnel")
	}
	return c.JSON(http.StatusOK, personnel)
}

func GetAllowance(c echo.Context) error {
	allowance, err := Personnel_model.GetAllowanceData()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching allowance")
	}
	return c.JSON(http.StatusOK, allowance)
}
