package model

type Admin struct {
	Amount float64 `db:"amount" json:"amount" validate:"gte=0"`
}

type AdminResponse struct {
	PersonalDeduction float64 `db:"personnel_deduction" json:"personalDeduction,omitempty" validate:"gte=0"`
	KReceipt          float64 `db:"k-receipt" json:"kReceipt,omitempty" validate:"gte=0"`
}
