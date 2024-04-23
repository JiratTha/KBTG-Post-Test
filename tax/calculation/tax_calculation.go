package calculation

import (
	Personnel_model "github.com/JiratTha/assessment-tax/Personnel/Personnel-model"
	"github.com/JiratTha/assessment-tax/db"
	_ "github.com/JiratTha/assessment-tax/db"
)

func TaxCalculation(totalIncome float64, wht float64) (taxAmount Personnel_model.TaxResponse) {
	taxAmount = Personnel_model.TaxResponse{}
	var personnelDeduction float64
	_ = db.DB.Get(&personnelDeduction, `SELECT  amount FROM project1."personnel_deduction" `)
	totalIncome -= personnelDeduction
	if totalIncome <= 150000.0 {
		taxAmount.Tax = 0.0
	} else if totalIncome <= 500000.0 {
		taxAmount.Tax = (totalIncome - 150000) * 0.1
	} else if totalIncome <= 1000000.0 {
		taxAmount.Tax = 35000 + ((totalIncome - 500000) * 0.15)
	} else if totalIncome <= 2000000.0 {
		taxAmount.Tax = 110000 + ((totalIncome - 1000000) * 0.2)
	} else {
		taxAmount.Tax = 310000 + ((totalIncome - 2000000) * 0.35)
	}
	taxAmount.Tax -= wht
	if taxAmount.Tax < 0 {
		taxAmount.Refund = -taxAmount.Tax
		taxAmount.Tax = 0
	}
	return taxAmount
}
