package model

type Admin struct {
	Amount float64 `db:"amount" json:"amount"`
}

type AdminResponse struct {
	PersonalDeduction float64 `db:"personnel_deduction" json:"personalDeduction ,omitempty"`
	Kreceipt          float64 `db:"k-receipt" json:"kReceipt ,omitempty"`
}
