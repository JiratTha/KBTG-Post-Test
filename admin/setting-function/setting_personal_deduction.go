package setting

import (
	"github.com/JiratTha/assessment-tax/admin/model"

	"github.com/JiratTha/assessment-tax/db"
	"log"
)

// SettingPersonalDeduction using SetPersonnelDeductPost
func SettingPersonalDeduction(personalDeduction model.Admin) (newPersonalDeduction model.AdminResponse) {
	var personalDeduct model.Admin
	if personalDeduction.Amount > 100000.0 {
		personalDeduct.Amount = 100000.0
	} else {
		personalDeduct.Amount = personalDeduction.Amount
	}
	_, err := db.DB.Exec(`UPDATE project1."personal_deduction" SET amount=$1 WHERE personal_deduction='personalDeduction'`, personalDeduct.Amount)
	if err != nil {
		log.Println(err)
	}
	newPersonalDeduction.PersonalDeduction = personalDeduct.Amount
	return newPersonalDeduction
}
