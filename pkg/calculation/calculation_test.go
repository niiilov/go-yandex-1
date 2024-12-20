package calculation

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	testCases := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "+",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "-",
			expression:     "2-1",
			expectedResult: 1,
		},
		{
			name:           "*",
			expression:     "2*2",
			expectedResult: 4,
		},
		{
			name:           "/",
			expression:     "4/2",
			expectedResult: 2,
		},
		{
			name:           "()",
			expression:     "1+(3*2)",
			expectedResult: 7,
		},
		{
			name:           "double operand",
			expression:     "22+10-(40/20)",
			expectedResult: 30,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", testCase.expression)
			}
			if val != fmt.Sprintf("%g", testCase.expectedResult) {
				t.Fatalf("%s should be equal %s", val, fmt.Sprintf("%g", testCase.expectedResult))
			}
		})
	}
}
