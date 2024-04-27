package setting

import (
	"github.com/JiratTha/assessment-tax/admin/model"

	"github.com/JiratTha/assessment-tax/db"
	"log"
)

// SettingPersonnelDeduction using SetPersonnelDeductPost
// @Summary SetPersonnelDeduct
// @Description Admin set personnel deduct
// @Tags admin
// @Accept  json
// @Produce  json
// @Param   personnelDeduct_body  body  _model.Admin  true  "Set Personal Deduct"
// @Success 200 {object} _model.AdminResponse  "Returns new personal Deduct"
// @Failure 400 {string} string "err"
// @Router /deductions/personal [post]
// @Router /deductions/personal [post]
func SettingPersonnelDeduction(personnelDeduction model.Admin) (newPersonnelDeduction model.AdminResponse) {
	var personnelDeduct model.Admin
	if personnelDeduction.Amount > 100000.0 {
		personnelDeduct.Amount = 100000.0
	} else {
		personnelDeduct.Amount = personnelDeduction.Amount
	}
	_, err := db.DB.Exec(`UPDATE project1."personnel_deduction" SET amount=$1 WHERE personnel_deduction='personnelDeduction'`, personnelDeduct.Amount)
	if err != nil {
		log.Println(err)
	}
	newPersonnelDeduction.PersonalDeduction = personnelDeduct.Amount
	return newPersonnelDeduction
}
