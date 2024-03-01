package parse

import (
	"fmt"
	"slices"
	"testing"
)

func TestInequalityFromString(t *testing.T) {
	tc := []struct {
		name     string
		str      string
		expected []float64
	}{
		{
			name:     "Should parse one token correctly",
			str:      "2x1>=0",
			expected: []float64{2, 0},
		},
		{
			name:     "Should parse one token correctly",
			str:      "x1>=7",
			expected: []float64{1, -7},
		},
		{
			name:     "Should parse one token correctly",
			str:      "1x1>=9",
			expected: []float64{1, -9},
		},
		{
			name:     "Should parse one token correctly",
			str:      "-x1<=0",
			expected: []float64{1, 0},
		},
		{
			name:     "Should parse one token correctly",
			str:      "-2x1>=55",
			expected: []float64{-2, -55},
		},
		{
			name:     "Should parse two tokens correctly",
			str:      "-x1+3x2<=1",
			expected: []float64{1, -3, 1},
		},
		{
			name:     "Should parse two tokens correctly",
			str:      "-x1-3x2<=0",
			expected: []float64{1, 3, 0},
		},
		{
			name:     "Should parse two tokens correctly",
			str:      "x1-3x2<=0",
			expected: []float64{-1, 3, 0},
		},
		{
			name:     "Should parse two tokens correctly",
			str:      "1x1-3x2>=4",
			expected: []float64{1, -3, -4},
		},
		{
			name:     "Should parse two tokens correctly",
			str:      "-x1-x2<=0",
			expected: []float64{1, 1, 0},
		},
		{
			name:     "Should parse two tokens correctly",
			str:      "x1-x2>=-1",
			expected: []float64{1, -1, 1},
		},
		{
			name:     "Should parse two tokens correctly",
			str:      "x1-4x3<=0",
			expected: []float64{-1, 0, 4, 0},
		},
		{
			name:     "Should parse tokens correctly",
			str:      "-x1-3x2+5x3+4x4>=-6",
			expected: []float64{-1, -3, 5, 4, 6},
		},
	}

	for _, tt := range tc {
		t.Run(fmt.Sprintf("%s: %s", tt.name, tt.str), func(t *testing.T) {
			t.Parallel()
			actual, _ := InequationFromString(tt.str)
			if !slices.Equal(tt.expected, actual) {
				t.Fatalf("expected: %v\ngot:%v", tt.expected, actual)
			}
		})
	}
}
