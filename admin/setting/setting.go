package setting

import (
	"github.com/JiratTha/assessment-tax/Personnel/model"
	"github.com/JiratTha/assessment-tax/db"
	"log"
)

// SettingPersonnelDeduction function for admin set new personal deduction and cannot set more than 100,000
func SettingPersonnelDeduction(personnelDeduction model.Admin) (newPersonnelDeduction model.AdminResponse) {
	var personnelDeduct model.Admin
	err := db.DB.Get(&personnelDeduct, `SELECT amount FROM project1."personnel_deduction" WHERE personnel_deduction='personnelDeduction'`)
	if err != nil {
		log.Println(err)
		return newPersonnelDeduction
	}
	if personnelDeduction.Amount > 100000.0 {
		personnelDeduct.Amount = 100000.0
	} else {
		personnelDeduct.Amount = personnelDeduction.Amount
	}
	_, err = db.DB.Exec(`UPDATE project1."personnel_deduction" SET amount=$1 WHERE personnel_deduction='personnelDeduction'`, personnelDeduct.Amount)
	if err != nil {

		return newPersonnelDeduction
	}
	newPersonnelDeduction.PersonalDeduction = personnelDeduct.Amount
	return newPersonnelDeduction
}