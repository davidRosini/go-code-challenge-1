package service

import (
	"codechallenge.test/domain"
	"codechallenge.test/usecase"
)

type ICalculateTaxService interface {
	Execute(operations []domain.OperationStock) []domain.TaxPay
}

type CalculateTaxService struct {
	buyOperationUsecase  usecase.IBuyOperationUsecase
	sellOperationUsecase usecase.ISellOperationUsecase
}

func NewCalculateTaxService(buyOperationUsecase usecase.IBuyOperationUsecase, sellOperationUsecase usecase.ISellOperationUsecase) ICalculateTaxService {
	return &CalculateTaxService{
		buyOperationUsecase:  buyOperationUsecase,
		sellOperationUsecase: sellOperationUsecase,
	}
}

func updateOperationState(state *domain.OperationState, newState domain.OperationState) {
	state.OperationsError += newState.OperationsError
	state.SharesHeld = newState.SharesHeld
	state.SharesWeightedAverage = newState.SharesWeightedAverage
	state.TotalLoss = newState.TotalLoss
	state.Tax = newState.Tax
}

func (ctu *CalculateTaxService) Execute(operations []domain.OperationStock) []domain.TaxPay {
	var taxes []domain.TaxPay
	var operationState *domain.OperationState = &domain.OperationState{}

	for _, op := range operations {
		var newState domain.OperationState

		switch op.Operation {
		case "buy":
			newState = ctu.buyOperationUsecase.Execute(op, operationState.SharesHeld, operationState.SharesWeightedAverage)
		case "sell":
			newState = ctu.sellOperationUsecase.Execute(op, *operationState)
		default:
			newState = domain.OperationState{
				Tax: domain.TaxPay{Error: "Unknown operation: " + op.Operation},
			}
			newState.OperationsError = 1
		}

		updateOperationState(operationState, newState)
		taxes = append(taxes, operationState.Tax)

		if operationState.OperationsError >= 3 {
			taxes = append(taxes, domain.TaxPay{Error: "Your account is blocked"})
			break
		}
	}

	return taxes
}
