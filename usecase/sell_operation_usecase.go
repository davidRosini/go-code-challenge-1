package usecase

import (
	"codechallenge.test/commons"
	"codechallenge.test/domain"
)

type ISellOperationUsecase interface {
	Execute(op domain.OperationStock, state domain.OperationState) domain.OperationState
}

type SellOperationUsecase struct {
	taxPercent float64
}

func NewSellOperationUsecase(taxPercent float64) ISellOperationUsecase {
	return &SellOperationUsecase{
		taxPercent: taxPercent,
	}
}

func (ctu *SellOperationUsecase) Execute(op domain.OperationStock, state domain.OperationState) domain.OperationState {

	var sharesHeld int32 = state.SharesHeld
	var sharesWeightedAverage float64 = state.SharesWeightedAverage
	var totalLoss float64 = state.TotalLoss

	if sharesHeld < op.Quantity {
		return domain.OperationState{
			Tax:                   domain.TaxPay{Error: "Can't sell more stocks than you have"},
			OperationsError:       1,
			SharesWeightedAverage: sharesWeightedAverage,
			SharesHeld:            sharesHeld,
			TotalLoss:             totalLoss,
		}
	}

	var tax float64 = 0

	sharesSaleValue := float64(op.Quantity) * op.UnitCost
	sharesAverageValue := float64(op.Quantity) * sharesWeightedAverage

	profit := sharesSaleValue - sharesAverageValue
	if profit < 0 {
		profit = 0
	}

	loss := sharesAverageValue - sharesSaleValue
	if loss < 0 {
		loss = 0
	}

	if profit > totalLoss && loss == 0 && sharesSaleValue > 20000 {
		tax = commons.CalculatePercentage((profit - totalLoss), ctu.taxPercent)
	}

	sharesHeld -= op.Quantity
	totalLoss = (totalLoss + loss) - profit
	if totalLoss < 0 {
		totalLoss = 0
	}

	return domain.OperationState{
		Tax:                   domain.TaxPay{Tax: commons.RoundUpTwoDigits(tax)},
		SharesHeld:            sharesHeld,
		SharesWeightedAverage: sharesWeightedAverage,
		TotalLoss:             totalLoss,
	}
}
