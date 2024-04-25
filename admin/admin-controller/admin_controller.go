package admin_controller

import (
	struc "github.com/JiratTha/assessment-tax/admin/model"
	"github.com/JiratTha/assessment-tax/admin/setting-function"
	"github.com/labstack/echo/v4"
	"net/http"
)

// SetPersonnelDeductPost using SettingPersonnelDeduction
func SetPersonnelDeductPost(c echo.Context) error {
	var admin struc.Admin
	if err := c.Bind(&admin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	newPersonalDeduct := setting_function.SettingPersonnelDeduction(admin)
	return c.JSON(http.StatusOK, newPersonalDeduct)
}
