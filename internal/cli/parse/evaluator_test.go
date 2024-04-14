package parse

import (
	"fmt"
	"slices"
	"testing"
)

func TestEvaluateFromString(t *testing.T) {
	tc := []struct {
		name     string
		str      string
		expected []float64
	}{
		{
			name:     "Should parse one token correctly",
			str:      "-x1",
			expected: []float64{-1},
		},
		{
			name:     "Should parse one token correctly",
			str:      "-2x1",
			expected: []float64{-2},
		},
		{
			name:     "Should parse two tokens correctly",
			str:      "-x1+3x2",
			expected: []float64{-1, 3},
		},
		{
			name:     "Should parse two tokens correctly",
			str:      "-x1-3x2",
			expected: []float64{-1, -3},
		},
		{
			name:     "Should parse two tokens correctly",
			str:      "x1-3x2",
			expected: []float64{1, -3},
		},
		{
			name:     "Should parse two tokens correctly",
			str:      "1x1-3x2",
			expected: []float64{1, -3},
		},
		{
			name:     "Should parse two tokens correctly",
			str:      "-x1-x2",
			expected: []float64{-1, -1},
		},
		{
			name:     "Should parse two tokens correctly",
			str:      "x1-x2",
			expected: []float64{1, -1},
		},
		{
			name:     "Should parse two tokens correctly",
			str:      "x1-4x3",
			expected: []float64{1, 0, -4},
		},
		{
			name:     "Should parse tokens correctly",
			str:      "-x1-3x2+5x3+4x4",
			expected: []float64{-1, -3, 5, 4},
		},
		{
			name:     "Should parse tokens correctly",
			str:      "-x1-3x2+5x3+4x4-444",
			expected: []float64{-1, -3, 5, 4, -444},
		},
		{
			name:     "Should parse tokens correctly",
			str:      "-2323x1-3x2+5x3+4x4-444",
			expected: []float64{-2323, -3, 5, 4, -444},
		},
		{
			name:     "Should parse tokens correctly",
			str:      "-2323x1-7777x2+898989x3+x4-444x6+424",
			expected: []float64{-2323, -7777, 898989, 1, 0, -444, 424},
		},
	}

	for _, tt := range tc {
		t.Run(fmt.Sprintf("%s: %s", tt.name, tt.str), func(t *testing.T) {
			actual, _ := NewEvaluator(tt.str).EvaluateFromString()
			if !slices.Equal(tt.expected, actual) {
				t.Fatalf("expected: %v\ngot:%v", tt.expected, actual)
			}
		})
	}
}
