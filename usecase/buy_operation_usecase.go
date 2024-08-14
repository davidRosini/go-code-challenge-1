package usecase

import (
	"codechallenge.test/commons"
	"codechallenge.test/domain"
)

type IBuyOperationUsecase interface {
	Execute(op domain.OperationStock, sharesHeld int32, sharesWeightedAverage float64) domain.OperationState
}

type BuyOperationUseCase struct {
}

func NewBuyOperationUsecase() IBuyOperationUsecase {
	return &BuyOperationUseCase{}
}

func (ctu *BuyOperationUseCase) Execute(op domain.OperationStock, sharesHeld int32, sharesWeightedAverage float64) domain.OperationState {
	sharesWeightedAverage = commons.CalculateWeightedAverage(float64(sharesHeld), sharesWeightedAverage, op.UnitCost, float64(op.Quantity))
	sharesHeld += op.Quantity
	return domain.OperationState{
		Tax:                   domain.TaxPay{Tax: 0},
		SharesHeld:            sharesHeld,
		SharesWeightedAverage: sharesWeightedAverage,
	}
}
