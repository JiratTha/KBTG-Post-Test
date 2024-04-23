package Personnel_model

type TaxResponse struct {
	Tax    float64 `json:"tax"`
	Refund float64 `json:"taxRefund ,omitempty"`
}
