package Personnel_controller

import (
	Personnel_model "github.com/JiratTha/assessment-tax/Personnel/Personnel-model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetPersonnel(c echo.Context) error {
	req := new(Personnel_model.Personnel)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// Call the GetPersonnelData function from the Personnel_model package with user input
	personnel, err := Personnel_model.GetPersonnelData(Personnel_model.Personnel{
		TotalIncome: req.TotalIncome,
		Wht:         req.Wht,
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
