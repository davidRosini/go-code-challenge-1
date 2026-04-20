package service_test

import (
	"reflect"
	"testing"

	"codechallenge.test/domain"
	"codechallenge.test/service"
	"codechallenge.test/usecase"
)

func TestCalculateTaxService_Execute(t *testing.T) {
	tests := []struct {
		name       string
		operations []domain.OperationStock
		wantTaxes  []domain.TaxPay
	}{
		{
			name: "Test case 1",
			operations: []domain.OperationStock{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 5000},
			},
			wantTaxes: []domain.TaxPay{{Tax: 0}, {Tax: 10000}},
		},
		{
			name: "Test case 2",
			operations: []domain.OperationStock{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 5.00, Quantity: 5000},
			},
			wantTaxes: []domain.TaxPay{{Tax: 0}, {Tax: 10000}, {Tax: 0}},
		},
		{
			name: "Test case 3",
			operations: []domain.OperationStock{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 5.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 3000},
			},
			wantTaxes: []domain.TaxPay{{Tax: 0}, {Tax: 0}, {Tax: 1000}},
		},
		{
			name: "Test case 4",
			operations: []domain.OperationStock{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 25.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 10000},
			},
			wantTaxes: []domain.TaxPay{{Tax: 0}, {Tax: 0}, {Tax: 0}},
		},
		{
			name: "Test case 5",
			operations: []domain.OperationStock{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 25.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 5000},
			},
			wantTaxes: []domain.TaxPay{{Tax: 0}, {Tax: 0}, {Tax: 0}, {Tax: 10000}},
		},
		{
			name: "Test case 6",
			operations: []domain.OperationStock{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 2.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 1000},
			},
			wantTaxes: []domain.TaxPay{{Tax: 0}, {Tax: 0}, {Tax: 0}, {Tax: 0}, {Tax: 3000}},
		},
		{
			name: "Test case 7",
			operations: []domain.OperationStock{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 2.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 1000},
				{Operation: "buy", UnitCost: 20.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 30.00, Quantity: 4350},
				{Operation: "sell", UnitCost: 30.00, Quantity: 650},
			},
			wantTaxes: []domain.TaxPay{{Tax: 0}, {Tax: 0}, {Tax: 0}, {Tax: 0}, {Tax: 3000}, {Tax: 0}, {Tax: 0}, {Tax: 3700}, {Tax: 0}},
		},
		{
			name: "Test case 8",
			operations: []domain.OperationStock{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 50.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 20.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 50.00, Quantity: 10000},
			},
			wantTaxes: []domain.TaxPay{{Tax: 0}, {Tax: 80000}, {Tax: 0}, {Tax: 60000}},
		},
		{
			name: "Test case 9 - Error sell more stocks then you have",
			operations: []domain.OperationStock{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 11000},
			},
			wantTaxes: []domain.TaxPay{{Tax: 0}, {Error: "Can't sell more stocks than you have"}},
		},
		{
			name: "Test case 10 - Try to sell more and then sell the correct amount",
			operations: []domain.OperationStock{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 11000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 5000},
			},
			wantTaxes: []domain.TaxPay{{Tax: 0}, {Error: "Can't sell more stocks than you have"}, {Tax: 10000}},
		},
		{
			name: "Test case 11 - Block account operations after 3 error tries",
			operations: []domain.OperationStock{
				{Operation: "sell", UnitCost: 20.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
			},
			wantTaxes: []domain.TaxPay{
				{Error: "Can't sell more stocks than you have"},
				{Error: "Can't sell more stocks than you have"},
				{Error: "Can't sell more stocks than you have"},
				{Error: "Your account is blocked"},
			},
		},
		{
			name:       "Test case 12 - Unknown operation returns error",
			operations: []domain.OperationStock{{Operation: "transfer", UnitCost: 10.00, Quantity: 100}},
			wantTaxes:  []domain.TaxPay{{Error: "Unknown operation: transfer"}},
		},
		{
			name: "Test case 13 - Block account after 3 unknown operations",
			operations: []domain.OperationStock{
				{Operation: "transfer", UnitCost: 10.00, Quantity: 100},
				{Operation: "swap", UnitCost: 10.00, Quantity: 100},
				{Operation: "merge", UnitCost: 10.00, Quantity: 100},
				{Operation: "buy", UnitCost: 10.00, Quantity: 100},
			},
			wantTaxes: []domain.TaxPay{
				{Error: "Unknown operation: transfer"},
				{Error: "Unknown operation: swap"},
				{Error: "Unknown operation: merge"},
				{Error: "Your account is blocked"},
			},
		},
		{
			name:       "Test case 14 - Empty operations list",
			operations: []domain.OperationStock{},
			wantTaxes:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := service.NewCalculateTaxService(usecase.NewBuyOperationUsecase(), usecase.NewSellOperationUsecase(0.2))
			gotTaxes := uc.Execute(tt.operations)

			if len(gotTaxes) != len(tt.wantTaxes) {
				t.Fatalf("Execute() length got = %d, want %d\ngot:  %v\nwant: %v", len(gotTaxes), len(tt.wantTaxes), gotTaxes, tt.wantTaxes)
			}

			if !reflect.DeepEqual(gotTaxes, tt.wantTaxes) {
				t.Errorf("Execute() got = %v, want %v", gotTaxes, tt.wantTaxes)
			}
		})
	}
}
