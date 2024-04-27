package allowance_calculation

import (
	"github.com/JiratTha/assessment-tax/Personnel/model"
	"github.com/JiratTha/assessment-tax/db"
)

// AllowanceCalculation calculate total allowance , seperate Allowance type and Limit max allowance amount by database
func AllowanceCalculation(allowance model.Personnel) (totalAllowance float64, totalDonation float64, totalKreciept float64) {
	var allowanceAmountDonation model.Allowance
	var allowanceAmountKReceipt model.Allowance
	var totalAllowanceAmount = 0.0
	var donationAmount = 0.0
	var kReceiptAmount = 0.0

	_ = db.DB.Get(&allowanceAmountDonation, `SELECT  amount FROM project1."allowance" WHERE allowance_type=$1`, "donation")
	_ = db.DB.Get(&allowanceAmountKReceipt, `SELECT  amount FROM project1."allowance" WHERE allowance_type=$1`, "k-receipt")

	for _, i := range allowance.Allowance {
		totalAllowanceAmount += i.Amount
		if i.AllowanceType == "donation" {
			donationAmount += i.Amount
			if donationAmount > allowanceAmountDonation.Amount {
				donationAmount = allowanceAmountDonation.Amount
			}
		}
		if i.AllowanceType == "k-receipt" {
			kReceiptAmount += i.Amount
			if kReceiptAmount > allowanceAmountKReceipt.Amount {
				kReceiptAmount = allowanceAmountKReceipt.Amount
			}
		}
	}

	totalAllowance = donationAmount + kReceiptAmount
	return totalAllowance, donationAmount, kReceiptAmount
}
