package model

type TaxResponse struct {
	Tax      float64    `json:"tax"`
	Refund   float64    `json:"taxRefund ,omitempty"`
	TaxLevel []TaxLevel `json:"taxLevel"`
}

type TaxLevel struct {
	Level  string  `json:"level"`
	Tax    float64 `json:"tax"`
	Refund float64 `json:"taxRefund ,omitempty"`
}
