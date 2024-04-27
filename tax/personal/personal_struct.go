package personal

type Personnel struct {
	TotalIncome float64     `json:"totalIncome" validate:"gte=0"`
	Wht         float64     `json:"wht" validate:"gte=0"`
	Allowance   []Allowance `json:"allowances" validate:"dive"`
}

type Allowance struct {
	AllowanceType string  `db:"allowance_type" json:"allowanceType" validate:"oneof=donation k-receipt"`
	Amount        float64 `db:"amount" json:"amount" validate:"gte=0"`
}
