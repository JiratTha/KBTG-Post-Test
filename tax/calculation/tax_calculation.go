package calculation

import (
	Personnel_model "github.com/JiratTha/assessment-tax/Personnel/Personnel-model"
	"github.com/JiratTha/assessment-tax/db"

)


func AllowanceCalculation(allowance Personnel_model.Personnel) (totalAllowance float64, totalDonation float64, totalKreciept float64) {
	var allowanceAmountDomation Personnel_model.Allowance
	var allowanceAmountKreceipt Personnel_model.Allowance
	var totalAllowanceAmount = 0.0
	var donationAmount = 0.0
	var krecieptAmount = 0.0
	_ = db.DB.Get(&allowanceAmountDomation, `SELECT  amount FROM project1."allowance" WHERE allowance_type=$1`, "donation")
	_ = db.DB.Get(&allowanceAmountKreceipt, `SELECT  amount FROM project1."allowance" WHERE allowance_type=$1`, "k-receipt")

	for _, i := range allowance.Allowance {
		totalAllowanceAmount += i.Amount
		if i.AllowanceType == "donation" {
			donationAmount += i.Amount
			if donationAmount > allowanceAmountDomation.Amount {
				donationAmount = allowanceAmountDomation.Amount
			}
		}
		if i.AllowanceType == "k-receipt" {
			krecieptAmount += i.Amount
			if krecieptAmount > allowanceAmountKreceipt.Amount {
				krecieptAmount = allowanceAmountKreceipt.Amount
			}
		}
	}

	totalAllowance = donationAmount + krecieptAmount
	return totalAllowance, donationAmount, krecieptAmount
}

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
