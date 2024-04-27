package tax_calculation

import (
	"github.com/JiratTha/assessment-tax/db"
	"github.com/JiratTha/assessment-tax/tax/model"
)

// TaxCalculation calculate tax , receive personal deduction from database
func TaxCalculation(totalIncome float64, wht float64) (taxResponse model.TaxResponse) {
	taxAmount := model.TaxLevel{}
	var personnelDeduction float64
	_ = db.DB.Get(&personnelDeduction, `SELECT  amount FROM project1."personnel_deduction" `)
	totalIncome -= personnelDeduction
	response := model.TaxResponse{
		Tax:    taxAmount.Tax,
		Refund: taxAmount.Refund,
		TaxLevel: []model.TaxLevel{
			{Level: "0-150,000", Tax: 0.0},
			{Level: "150,001-500,000", Tax: 0.0},
			{Level: "500,001-1,000,000", Tax: 0.0},
			{Level: "1,000,001-2,000,000", Tax: 0.0},
			{Level: "2,000,001 ขึ้นไป", Tax: 0.0},
		},
	}

	if totalIncome <= 150000.0 {
		taxAmount.Level = "0-150000"
		taxAmount.Tax = 0.0
	} else if totalIncome <= 500000.0 {
		taxAmount.Level = "150,001-500,000"
		taxAmount.Tax = (totalIncome - 150000) * 0.1
		response.TaxLevel[1].Tax = taxAmount.Tax

	} else if totalIncome <= 1000000.0 {
		taxAmount.Level = "500,001-1,000,000"
		taxAmount.Tax = 35000 + ((totalIncome - 500000) * 0.15)
		response.TaxLevel[1].Tax = 35000
		response.TaxLevel[2].Tax = taxAmount.Tax - 35000
	} else if totalIncome <= 2000000.0 {
		taxAmount.Level = "1,000,001-2,000,000"
		taxAmount.Tax = 110000 + ((totalIncome - 1000000) * 0.2)
		response.TaxLevel[1].Tax = 35000
		response.TaxLevel[2].Tax = 75000.0
		response.TaxLevel[3].Tax = taxAmount.Tax - 75000 - 35000
	} else {
		taxAmount.Level = "2,000,001 ขึ้นไป"
		taxAmount.Tax = 310000 + ((totalIncome - 2000000) * 0.35)
		response.TaxLevel[1].Tax = 35000
		response.TaxLevel[2].Tax = 75000.0
		response.TaxLevel[3].Tax = 200000.0
		response.TaxLevel[4].Tax = taxAmount.Tax - 200000 - 75000 - 35000
	}

	taxAmount.Tax -= wht
	if taxAmount.Tax < 0 {
		taxAmount.Refund = -taxAmount.Tax
		taxAmount.Tax = 0
	}
	response.Tax = taxAmount.Tax
	response.Refund = taxAmount.Refund
	return response
}
