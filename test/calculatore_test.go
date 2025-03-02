package calculator

import (
	"testing"

	"github.com/SashaMorkovkin/Final_task_2/internal/calculator"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
	}{
		{"2+2*2", 6},
		{"(3+5)*2", 16},
		{"10/2+3", 8},
		{"(2+2)*2-3", 5},
	}

	for _, test := range tests {
		result, err := calculator.Calculate(test.expression)
		if err != nil {
			t.Errorf("Error calculating expression %s: %v", test.expression, err)
		}
		if result != test.expected {
			t.Errorf("Expected %f, got %f for expression %s", test.expected, result, test.expression)
		}
	}
}
