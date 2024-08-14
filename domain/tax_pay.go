package domain

type TaxPay struct {
	Tax   float64 `json:"tax"`
	Error string  `json:"error"`
}
