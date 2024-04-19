package Personnel_model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"log"
)

type Personnel struct {
	TotalIncome float64 `json:"totalincome"`
	Wht         float64 `json:"wht"`
}

type Allowance struct {
	Allowances []struct {
		AllowanceType string  `db:"allowancetype"`
		Amount        float64 `db:"amount"`
	} `db:"allowances"`
}

type PersonnelDataProvider interface {
	GetData() (Personnel, error)
}

func (u Personnel) GetData() (Personnel, error) {

	return Personnel{
		TotalIncome: u.TotalIncome, Wht: u.Wht,
	}, nil
}

var db *sqlx.DB

func InitDB(dataSourceName string) error {
	var err error
	db, err = sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return nil
}

func GetPersonnelData(provider PersonnelDataProvider) (Personnel, error) {
	return provider.GetData()
}

func GetAllowanceData() ([]Allowance, error) {
	var allowances []Allowance
	err := db.Select(&allowances, `SELECT allowancetype , amount  FROM allowance."allowance"`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return allowances, nil
}
