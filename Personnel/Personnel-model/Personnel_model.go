package Personnel_model

import (
	"github.com/JiratTha/assessment-tax/db"
)

type Personnel struct {
	TotalIncome float64     `json:"totalIncome"`
	Wht         float64     `json:"wht"`
	Allowance   []Allowance `json:"allowances"`
}

type Allowance struct {
	AllowanceType string  `db:"allowance_type" json:"allowanceType"`
	Amount        float64 `db:"amount" json:"amount"`
}

type PersonnelDataProvider interface {
	GetData() (Personnel, error)
}

func (u Personnel) GetData() (Personnel, error) {

	return Personnel{
		TotalIncome: u.TotalIncome, Wht: u.Wht,
	}, nil
}

func GetAllowanceData() ([]Allowance, error) {
	var allowances []Allowance
	err := db.DB.Select(&allowances, `SELECT allowance_type , amount  FROM project1."allowance"`)
	if err != nil {
		return nil, err
	}

	return allowances, nil
}
