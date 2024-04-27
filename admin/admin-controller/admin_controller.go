package admin_controller

import (
	struc "github.com/JiratTha/assessment-tax/admin/model"
	"github.com/JiratTha/assessment-tax/admin/setting-function"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// SetPersonalDeductPost using func SettingPersonalDeduction
// @Summary Set personal deduction
// @Description Set personal deduction. personal deduction must be greater than 10000 and cannot greater than 100000
// @Tags tax
// @Accept  json
// @Produce  json
// @Param   admin_body  body      model.Admin  true  "personal deduct Request"
// @Success 200 {object} model.AdminResponse  "Returns new personal deduct amount"
// @Failure 400 {string} string "Invalid input"
// @Router  /admin/deductions/personal [post]
func SetPersonalDeductPost(c echo.Context) error {
	var reqPersonalAmount struc.Admin
	if err := c.Bind(&reqPersonalAmount); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	err := c.Validate(reqPersonalAmount)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if reqPersonalAmount.Amount <= 10000 {
		return echo.NewHTTPError(http.StatusBadRequest, "personal Deduction must be greater than 10000")
	}
	newPersonalDeduct := setting.SettingPersonalDeduction(reqPersonalAmount)
	return c.JSON(http.StatusOK, newPersonalDeduct)
}

// SetKReceiptPost using func SettingKReceipt
// @Summary Set K-receipt
// @Description Set K-receipt. K-receipt must be greater than 0 and cannot greater than 100000
// @Tags tax
// @Accept  json
// @Produce  json
// @Param   admin_body  body      model.Admin  true  "k-receipt Request"
// @Success 200 {object} model.AdminResponse  "Returns new k-receipt deduct amount"
// @Failure 400 {string} string "Invalid input"
// @Router /admin/deductions/k-receipt [post]
func SetKReceiptPost(c echo.Context) error {
	var reqKReceiptAmount struc.Admin
	if err := c.Bind(&reqKReceiptAmount); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	err := c.Validate(reqKReceiptAmount)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if reqKReceiptAmount.Amount <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "K-receipt must be greater than 0")
	}
	newKReceipt := setting.SettingKReceipt(reqKReceiptAmount)
	return c.JSON(http.StatusOK, newKReceipt)
}
