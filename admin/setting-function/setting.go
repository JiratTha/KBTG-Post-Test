package setting

import (
	"github.com/JiratTha/assessment-tax/Personnel/model"
	"github.com/JiratTha/assessment-tax/db"
	"log"
)

// SettingPersonnelDeduction function for admin set new personal deduction amount and cannot set more than 100,000
func SettingPersonnelDeduction(personnelDeduction model.Admin) (newPersonnelDeduction model.AdminResponse) {
	var personnelDeductReq model.Admin
	if personnelDeduction.Amount > 100000 {
		personnelDeductReq.Amount = 100000
	} else {
		personnelDeductReq.Amount = personnelDeduction.Amount
	}
	_, err := db.DB.Exec(`UPDATE project1."personnel_deduction" SET amount=$1 WHERE personnel_deduction='personnelDeduction'`, personnelDeductReq.Amount)
	if err != nil {
		log.Println(err)
	}
	newPersonnelDeduction.PersonalDeduction = personnelDeductReq.Amount
	return newPersonnelDeduction
}

// SettingKReceipt function for admin set new k-receipt amount and cannot set more than 100,000
func SettingKReceipt(kReceipt model.Admin) (newKReceipt model.AdminResponse) {
	var kReceiptReq model.Admin
	if kReceipt.Amount > 100000 {
		kReceiptReq.Amount = 100000
	} else {
		kReceiptReq.Amount = kReceipt.Amount
	}
	_, err := db.DB.Exec(`UPDATE project1."allowance" SET amount=$1 WHERE allowance_type='k-receipt'`, kReceiptReq.Amount)
	if err != nil {
		log.Println(err)
	}
	newKReceipt.KReceipt = kReceiptReq.Amount
	return newKReceipt
}
