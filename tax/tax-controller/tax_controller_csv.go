package tax_controller

import (
	"encoding/csv"
	Personnel_model "github.com/JiratTha/assessment-tax/Personnel/model"
	"github.com/JiratTha/assessment-tax/db"
	struc "github.com/JiratTha/assessment-tax/tax/model"
	tax_calculation "github.com/JiratTha/assessment-tax/tax/tax-calculation"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"strconv"
)

func TaxCalculationCSVPost(c echo.Context) error {

	file, err := c.FormFile("taxFile")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "The file retrieval was unsuccessful")
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	reader := csv.NewReader(src)
	var taxes []struc.TaxResponseCSVDataStruct

	if _, err = reader.Read(); err != nil {
		return err
	}
	var loopNumber = 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "The file reading was unsuccessful")
		}

		totalIncome, err := strconv.ParseFloat(record[0], 64)
		wht, err := strconv.ParseFloat(record[1], 64)
		donation, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid input value")
		}
		var personalDeduction float64
		var allowanceAmountDomation Personnel_model.Allowance
		_ = db.DB.Get(&personalDeduction, `SELECT  amount FROM project1."personnel_deduction" `)
		_ = db.DB.Get(&allowanceAmountDomation, `SELECT  amount FROM project1."allowance" WHERE allowance_type=$1`, "donation")
		//personalDeduct, _ := model.GetPersonalDeduct()
		//donationDeduct, _ := model.GetDonationDeduct()
		totalIncomeDeductPersonal := totalIncome - personalDeduction
		totalIncomeDeductDonation := totalIncomeDeductPersonal - allowanceAmountDomation.Amount
		tax := tax_calculation.TaxCalculation(totalIncomeDeductDonation, wht)

		taxes = append(
			taxes,
			struc.TaxResponseCSVDataStruct{
				TotalIncomeCSV: totalIncome,
				TaxCSV:         tax.Tax,
			},
		)

		loopNumber += 1
	}

	return c.JSON(http.StatusOK, struc.TaxResponseCSVStruct{Taxes: taxes})
}
