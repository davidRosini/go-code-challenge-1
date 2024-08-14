package commons_test

import (
	"math"
	"testing"

	"codechallenge.test/commons"
)

func TestCalculateSharesWeightedAverage(t *testing.T) {
	tests := []struct {
		name              string
		sharesHeld        float64
		sharesWeightedAvg float64
		sharesCost        float64
		sharesQuantity    float64
		expected          float64
	}{
		{
			name:              "Test case 1: normal calculation",
			sharesHeld:        100,
			sharesWeightedAvg: 100.0,
			sharesCost:        120.0,
			sharesQuantity:    50,
			expected:          106.67,
		},
		{
			name:              "Test case 2: zero shares held",
			sharesHeld:        0,
			sharesWeightedAvg: 0.0,
			sharesCost:        100.0,
			sharesQuantity:    100,
			expected:          100.0,
		},
		{
			name:              "Test case 3: zero shares cost",
			sharesHeld:        1,
			sharesWeightedAvg: 0.0,
			sharesCost:        0.0,
			sharesQuantity:    1,
			expected:          0.0,
		},
		{
			name:              "Test case 4: all zeros",
			sharesHeld:        0,
			sharesWeightedAvg: 0.0,
			sharesCost:        0.0,
			sharesQuantity:    0,
			expected:          math.NaN(),
		},
		{
			name:              "Test case 5: negative shares cost",
			sharesHeld:        0,
			sharesWeightedAvg: -100.0,
			sharesCost:        -100.0,
			sharesQuantity:    0,
			expected:          math.NaN(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := commons.CalculateWeightedAverage(tt.sharesHeld, tt.sharesWeightedAvg, tt.sharesCost, tt.sharesQuantity)
			if math.IsNaN(tt.expected) && math.IsNaN(got) {
				return
			}
			if commons.RoundUpTwoDigits(got) != tt.expected {
				t.Errorf("CalculateSharesWeightedAverage() = %f, but want %f", got, tt.expected)
			}
		})
	}
}

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		name     string
		profit   float64
		expected float64
	}{
		{
			name:     "Test case 1: positive profit",
			profit:   1000.0,
			expected: 200.0,
		},
		{
			name:     "Test case 1: zero profit",
			profit:   0.0,
			expected: 0.0,
		},
		{
			name:     "Test case 1: negative profit",
			profit:   -500.0,
			expected: 0.0,
		},
		{
			name:     "Test case 1: large profit",
			profit:   1000000.0,
			expected: 200000.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := commons.CalculatePercentage(tt.profit, 0.2)
			if got != tt.expected {
				t.Errorf("CalculateTax() = %f, but want %f", got, tt.expected)
			}
		})
	}
}

func TestCalculateTaxUsecase_RoundUpTwoDigits_NegativeInput(t *testing.T) {
	testCases := []struct {
		name string
		val  float64
		want float64
	}{
		{
			name: "Test case 1: normal value",
			val:  -12.345,
			want: -12.35,
		},
		{
			name: "Test case 2: zero value",
			val:  0.0,
			want: 0.0,
		},
		{
			name: "Test case 3: large value",
			val:  123456789.012345,
			want: 123456789.01,
		},
		{
			name: "Test case 4: very small value",
			val:  0.0000000000000001,
			want: 0.0,
		},
		{
			name: "Test case 5: value with precision",
			val:  12.344999999999999,
			want: 12.35,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := commons.RoundUpTwoDigits(tc.val)
			if got != tc.want {
				t.Errorf("RoundUpTwoDigits() = %v, but want %v", got, tc.want)
			}
		})
	}
}
