package usecase_test

import (
	"reflect"
	"testing"

	"codechallenge.test/domain"
	"codechallenge.test/usecase"
)

func TestSellOperationUsecase_Execute(t *testing.T) {
	tests := []struct {
		name      string
		operation domain.OperationStock
		state     domain.OperationState
		want      domain.OperationState
	}{
		{
			name:      "Test case 1",
			operation: domain.OperationStock{Operation: "sell", UnitCost: 20.00, Quantity: 5000},
			state: domain.OperationState{
				SharesHeld:            10000,
				SharesWeightedAverage: 10,
			},
			want: domain.OperationState{
				Tax:                   domain.TaxPay{Tax: 10000},
				SharesHeld:            5000,
				SharesWeightedAverage: 10,
			},
		},
		{
			name:      "Test case 2",
			operation: domain.OperationStock{Operation: "sell", UnitCost: 20.00, Quantity: 10001},
			state: domain.OperationState{
				SharesHeld:            10000,
				SharesWeightedAverage: 10,
			},
			want: domain.OperationState{
				Tax:                   domain.TaxPay{Error: "Can't sell more stocks than you have"},
				OperationsError:       1,
				SharesHeld:            10000,
				SharesWeightedAverage: 10,
			},
		},
		{
			name:      "Test case 3",
			operation: domain.OperationStock{Operation: "sell", UnitCost: 25.00, Quantity: 1000},
			state: domain.OperationState{
				SharesHeld:            1000,
				SharesWeightedAverage: 10,
			},
			want: domain.OperationState{
				Tax:                   domain.TaxPay{Tax: 3000},
				SharesHeld:            0,
				SharesWeightedAverage: 10,
			},
		},
		{
			name:      "Test case 4",
			operation: domain.OperationStock{Operation: "sell", UnitCost: 2.00, Quantity: 5000},
			state: domain.OperationState{
				SharesHeld:            10000,
				SharesWeightedAverage: 10,
			},
			want: domain.OperationState{
				Tax:                   domain.TaxPay{Tax: 0},
				SharesHeld:            5000,
				TotalLoss:             40000,
				SharesWeightedAverage: 10,
			},
		},
		{
			name:      "Test case 5",
			operation: domain.OperationStock{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
			state: domain.OperationState{
				SharesHeld:            5000,
				TotalLoss:             40000,
				SharesWeightedAverage: 10,
			},
			want: domain.OperationState{
				Tax:                   domain.TaxPay{Tax: 0},
				SharesHeld:            3000,
				TotalLoss:             20000,
				SharesWeightedAverage: 10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := usecase.NewSellOperationUsecase(0.2)
			got := uc.Execute(tt.operation, tt.state)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}
