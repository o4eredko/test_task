package calculator_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"calculator/calculator"
)

func TestCalculator_Evaluate(t *testing.T) {
	type (
		testCase struct {
			name       string
			expression string
			want       float64
			wantErr    error
		}
	)
	testCases := []testCase{
		{
			name:       "single number",
			expression: "+1000000",
			want:       1_000_000,
		},
		{
			name:       "simple addition",
			expression: "-2 + 2 + 2",
			want:       2,
		},
		{
			name:       "simple substraction",
			expression: "10 - 100 - 5",
			want:       -95,
		},
		{
			name:       "simple division",
			expression: "18 / 3 / -5",
			want:       6.0 / -5.0,
		},
		{
			name:       "zero division",
			expression: "100 / -0",
			want:       math.Inf(-1),
		},
		{
			name:       "simple multiplication",
			expression: "1000 * 234",
			want:       234_000,
		},
		{
			name:       "order of operators",
			expression: "2 + 2 * 2",
			want:       6,
		},
		{
			name:       "order of operators",
			expression: "(2 + 2) * 2",
			want:       8,
		},
		{
			name:       "order of operators",
			expression: "(2 + (2 - (2))) / 2",
			want:       1,
		},
		{
			name:       "order of operators",
			expression: "(2 * (2 * 4)) / 2",
			want:       8,
		},
		{
			name:       "int limit",
			expression: "2147483647 * 2",
			want:       2147483647 * 2,
		},
		{
			name:       "invalid expression",
			expression: "100 -",
			wantErr:    calculator.ErrInvalidExpression,
		},
		{
			name:       "invalid symbols",
			expression: "100 - abc",
			wantErr:    calculator.ErrInvalidExpression,
		},
		{
			name:       "invalid symbols",
			expression: "100 * 11abc",
			wantErr:    calculator.ErrInvalidExpression,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calc := calculator.NewCalculator(tc.expression)
			got, gotErr := calc.Evaluate()
			assert.Equal(t, tc.want, got)
			assert.ErrorIs(t, gotErr, tc.wantErr)
		})
	}
}
