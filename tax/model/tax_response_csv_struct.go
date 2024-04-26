package model

type TaxResponseCSVDataStruct struct {
	TotalIncomeCSV float64 `json:"totalIncome" validate:"gte=0"`
	TaxCSV         float64 `json:"tax" validate:"gte=0"`
	TaxRefund      float64 `json:"taxRefund,omitempty" validate:"gte=0" `
}

type TaxResponseCSVStruct struct {
	TaxesCSV []TaxResponseCSVDataStruct `json:"taxes" validate:"gte=0"`
}
