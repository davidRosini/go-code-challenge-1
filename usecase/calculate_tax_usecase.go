package usecase

import (
	"fmt"

	"codechallenge.test/commons"
	"codechallenge.test/domain"
)

type ICalculateTaxUsecase interface {
	Execute(operations []domain.OperationStock) []domain.TaxPay
}

type calculateTaxUsecase struct {
	taxPercent float64
}

func NewCalculateTaxUsecase() ICalculateTaxUsecase {
	return &calculateTaxUsecase{
		taxPercent: 0.2, // tax percent 20% of the profit
	}
}

func (ctu *calculateTaxUsecase) Execute(operations []domain.OperationStock) []domain.TaxPay {
	var taxes []domain.TaxPay

	var operationsError int8 = 0
	var sharesHeld int32 = 0
	var sharesWeightedAverage float64 = 0
	var totalLoss float64 = 0

	for _, op := range operations {
		var tax float64 = 0

		if operationsError == 3 {
			break
		}

		switch op.Operation {

		case "buy":
			sharesWeightedAverage = commons.CalculateWeightedAverage(float64(sharesHeld), sharesWeightedAverage, op.UnitCost, float64(op.Quantity))
			sharesHeld += op.Quantity

		case "sell":
			if sharesHeld < op.Quantity {
				taxes = append(taxes, domain.TaxPay{Error: "Can't sell more stocks than you have"})
				operationsError += 1
				continue
			}

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
		}

		fmt.Printf("shares %d,totalLoss %f, sharesWeightedAverage %f\n", sharesHeld, totalLoss, sharesWeightedAverage)
		taxes = append(taxes, domain.TaxPay{Tax: commons.RoundUpTwoDigits(tax)})
	}

	if operationsError == 3 {
		taxes = append(taxes, domain.TaxPay{Error: "Your account is blocked"})
	}

	return taxes
}
