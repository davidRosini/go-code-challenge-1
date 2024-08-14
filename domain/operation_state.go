package domain

type OperationState struct {
	Tax                   TaxPay
	OperationsError       int8
	SharesHeld            int32
	SharesWeightedAverage float64
	TotalLoss             float64
}
