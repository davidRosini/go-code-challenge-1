package adapter_test

import (
	"bytes"
	"math"
	"strings"
	"testing"

	"codechallenge.test/adapter"
	"codechallenge.test/domain"
)

type stubServiceFunc struct {
	funcExecute func(operations []domain.OperationStock) []domain.TaxPay
}

func (fn stubServiceFunc) Execute(operations []domain.OperationStock) []domain.TaxPay {
	return fn.funcExecute(operations)
}

func TestExecute_ValidInput(t *testing.T) {
	s := stubServiceFunc{
		funcExecute: func(operations []domain.OperationStock) []domain.TaxPay {
			return []domain.TaxPay{{Tax: 0}, {Tax: 10000}}
		},
	}

	mockInput := "[{\"operation\":\"buy\", \"unit-cost\":10.00, \"quantity\": 10000}, {\"operation\":\"sell\", \"unit-cost\":20.00, \"quantity\": 5000}]\n"
	expectedOutput := "[{\"tax\":0},{\"tax\":10000}]\n"

	var output bytes.Buffer
	handler := adapter.NewCalculateTaxHandler(strings.NewReader(mockInput), &output, s)
	handler.Execute()

	if output.String() != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, output.String())
	}
}

func TestExecute_InvalidJSON(t *testing.T) {
	s := stubServiceFunc{
		funcExecute: func(operations []domain.OperationStock) []domain.TaxPay {
			t.Fatal("service should not be called for invalid JSON")
			return nil
		},
	}

	mockInput := "invalid json\n"

	var output bytes.Buffer
	handler := adapter.NewCalculateTaxHandler(strings.NewReader(mockInput), &output, s)
	handler.Execute()

	if output.String() != "" {
		t.Errorf("Expected empty output for invalid JSON, but got %q", output.String())
	}
}

func TestExecute_InvalidJSONFollowedByValid(t *testing.T) {
	s := stubServiceFunc{
		funcExecute: func(operations []domain.OperationStock) []domain.TaxPay {
			return []domain.TaxPay{{Tax: 0}}
		},
	}

	mockInput := "not json\n[{\"operation\":\"buy\", \"unit-cost\":10.00, \"quantity\": 100}]\n"
	expectedOutput := "[{\"tax\":0}]\n"

	var output bytes.Buffer
	handler := adapter.NewCalculateTaxHandler(strings.NewReader(mockInput), &output, s)
	handler.Execute()

	if output.String() != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, output.String())
	}
}

func TestExecute_MarshalError(t *testing.T) {
	s := stubServiceFunc{
		funcExecute: func(operations []domain.OperationStock) []domain.TaxPay {
			return []domain.TaxPay{{Tax: math.NaN()}}
		},
	}

	mockInput := "[{\"operation\":\"buy\", \"unit-cost\":10.00, \"quantity\": 100}]\n"

	var output bytes.Buffer
	handler := adapter.NewCalculateTaxHandler(strings.NewReader(mockInput), &output, s)
	handler.Execute()

	if output.String() != "" {
		t.Errorf("Expected empty output on marshal error, but got %q", output.String())
	}
}

func TestExecute_MultipleInputLines(t *testing.T) {
	callCount := 0
	s := stubServiceFunc{
		funcExecute: func(operations []domain.OperationStock) []domain.TaxPay {
			callCount++
			return []domain.TaxPay{{Tax: float64(callCount * 100)}}
		},
	}

	mockInput := "[{\"operation\":\"buy\", \"unit-cost\":10.00, \"quantity\": 100}]\n[{\"operation\":\"sell\", \"unit-cost\":20.00, \"quantity\": 50}]\n"

	var output bytes.Buffer
	handler := adapter.NewCalculateTaxHandler(strings.NewReader(mockInput), &output, s)
	handler.Execute()

	expectedOutput := "[{\"tax\":100}]\n[{\"tax\":200}]\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, output.String())
	}
}
