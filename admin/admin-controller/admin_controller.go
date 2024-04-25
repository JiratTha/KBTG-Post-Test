package admin_controller

import (
	struc "github.com/JiratTha/assessment-tax/Personnel/model"
	"github.com/JiratTha/assessment-tax/admin/setting"
	"github.com/labstack/echo/v4"
	"net/http"
)

// SetPersonnelDeductPost using func SettingPersonnelDeduction
func SetPersonnelDeductPost(c echo.Context) error {
	var reqPersonalAmount struc.Admin
	if err := c.Bind(&reqPersonalAmount); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	if reqPersonalAmount.Amount <= 10000 {
		return echo.NewHTTPError(http.StatusBadRequest, "Personal Deduction must be greater than 10000")
	}
	newPersonalDeduct := setting.SettingPersonnelDeduction(reqPersonalAmount)
	return c.JSON(http.StatusOK, newPersonalDeduct)
}

// SetKReceiptPost using func SettingKReceipt
func SetKReceiptPost(c echo.Context) error {
	var reqKReceiptAmount struc.Admin
	if err := c.Bind(&reqKReceiptAmount); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	if reqKReceiptAmount.Amount <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "K-receipt must be greater than 0")
	}
	newKReceipt := setting.SettingKReceipt(reqKReceiptAmount)
	return c.JSON(http.StatusOK, newKReceipt)
}
