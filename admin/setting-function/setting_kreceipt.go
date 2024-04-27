package setting

import (
	"github.com/JiratTha/assessment-tax/admin/model"
	"github.com/JiratTha/assessment-tax/db"
	"log"
)

// SettingKReceipt using SetKReceiptPost
// @Summary SetKReceipt
// @Description Admin set k-receipt
// @Tags admin
// @Accept  json
// @Produce  json
// @Param   KReceipt_body  body  _model.Admin  true  "Set k-receipt"
// @Success 200 {object} _model.AdminResponse  "Returns new k-receipt"
// @Failure 400 {string} string "err"
// @Router /deductions/k-receipt [post]
// @Router /deductions/k-receipt [post]
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
