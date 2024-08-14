package usecase_test

import (
	"reflect"
	"testing"

	"codechallenge.test/domain"
	"codechallenge.test/usecase"
)

func TestBuyOperationUsecase_Execute(t *testing.T) {
	tests := []struct {
		name                  string
		operation             domain.OperationStock
		sharesHeld            int32
		sharesWeightedAverage float64
		want                  domain.OperationState
	}{
		{
			name:                  "Test case 1",
			operation:             domain.OperationStock{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
			sharesHeld:            100,
			sharesWeightedAverage: 10,
			want: domain.OperationState{
				Tax:                   domain.TaxPay{Tax: 0},
				SharesHeld:            10100,
				SharesWeightedAverage: 10,
			},
		},
		{
			name:                  "Test case 2",
			operation:             domain.OperationStock{Operation: "buy", UnitCost: 25.00, Quantity: 5000},
			sharesHeld:            10000,
			sharesWeightedAverage: 10,
			want: domain.OperationState{
				Tax:                   domain.TaxPay{Tax: 0},
				SharesHeld:            15000,
				SharesWeightedAverage: 15,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := usecase.NewBuyOperationUsecase()
			got := uc.Execute(tt.operation, tt.sharesHeld, tt.sharesWeightedAverage)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}
