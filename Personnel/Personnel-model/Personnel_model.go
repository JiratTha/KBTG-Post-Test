package Personnel_model

type Personnel struct {
	TotalIncome float64     `json:"totalIncome"`
	Wht         float64     `json:"wht"`
	Allowance   []Allowance `json:"allowances"`
}

type Allowance struct {
	AllowanceType string  `db:"allowance_type" json:"allowanceType"`
	Amount        float64 `db:"amount" json:"amount"`
}
