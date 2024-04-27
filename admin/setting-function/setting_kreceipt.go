package setting

import (
	"github.com/JiratTha/assessment-tax/admin/model"
	"github.com/JiratTha/assessment-tax/db"
	"log"
)

// SettingKReceipt using SetKReceiptPost
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
