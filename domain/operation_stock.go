package domain

type OperationStock struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int32   `json:"quantity"`
}
