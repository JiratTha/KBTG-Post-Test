package calculation

import (
	Personnel_model "github.com/JiratTha/assessment-tax/Personnel/Personnel-model"
	_ "github.com/JiratTha/assessment-tax/db"
)

func TaxCalculation(totalIncome float64) (taxAmount Personnel_model.TaxResponse) {
	taxAmount = Personnel_model.TaxResponse{}
	var PersonnelDeduction = 60000.0
	totalIncome -= PersonnelDeduction
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

	return taxAmount
}
