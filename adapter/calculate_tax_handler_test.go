package adapter_test

import (
	"bytes"
	"os"
	"testing"

	"codechallenge.test/adapter"
	"codechallenge.test/domain"
)

type stubUseCaseFunc struct {
	funcExecute func(operations []domain.OperationStock) []domain.TaxPay
}

func (fn stubUseCaseFunc) Execute(operations []domain.OperationStock) []domain.TaxPay {
	return fn.funcExecute(operations)
}

// Test is breaking cant read output
func TestExecute(t *testing.T) {
	s := stubUseCaseFunc{
		funcExecute: func(operations []domain.OperationStock) []domain.TaxPay {
			return []domain.TaxPay{{Tax: 0}, {Tax: 10000}}
		},
	}

	mockInput := `[{"operation":"buy", "unit-cost":10.00, "quantity": 10000}, {"operation":"sell", "unit-cost":20.00, "quantity": 5000}]`
	expectedOutput := `[{"tax":0}, {"tax":10000}]`

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	r, w, _ := os.Pipe()
	os.Stdin = r

	var output bytes.Buffer
	handler := adapter.NewCalculateTaxHandler(&output, s)

	w.Write([]byte(mockInput))
	w.Close()

	go handler.Execute()

	actualOutput := output.String()

	if actualOutput != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, actualOutput)
	}
}
