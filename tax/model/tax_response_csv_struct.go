package model

type TaxResponseCSVDataStruct struct {
	TotalIncomeCSV float64 `json:"totalIncome" validate:"gte=0"`
	TaxCSV         float64 `json:"tax" validate:"gte=0"`
}

type TaxResponseCSVStruct struct {
	TaxesCSV []TaxResponseCSVDataStruct `json:"taxes" validate:"gte=0"`
}
