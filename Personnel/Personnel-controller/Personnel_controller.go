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
	return c.JSON(http.StatusOK, req)
}

//func GetAllowance(c echo.Context) error {
//	allowance, err := Personnel_model.GetAllowanceData()
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, "Error fetching allowance")
//	}
//	return c.JSON(http.StatusOK, allowance)
//}
