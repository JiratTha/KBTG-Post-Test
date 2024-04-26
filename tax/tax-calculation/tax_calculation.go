package tax_calculation

import (
	Personnel_model "github.com/JiratTha/assessment-tax/Personnel/model"
	"github.com/JiratTha/assessment-tax/db"
	"github.com/JiratTha/assessment-tax/tax/model"
)

// AllowanceCalculation calculate total allowance , seperate Allowance type and Limit max allowance amount by database
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

// TaxCalculation calculate tax , receive Personnel deduction from database
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
