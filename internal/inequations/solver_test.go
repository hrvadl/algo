package inequations

import (
	"slices"
	"testing"

	"github.com/hrvadl/algo/internal/matrix"
)

func TestSolveWithSupportSolution(t *testing.T) {
	tc := []struct {
		name     string
		m        matrix.Matrix
		expected matrix.Row
	}{
		{
			name: "Should solve inequation correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{-1, 0, 3, -2, 1, 3},
					{1, -1, 0, 1, 1, 3},
					{-1, -3, 1, 1, -1, -2},
					{-1, 1, 0, 0, 1, 0},
				},
			},
			expected: matrix.Row{2, 0, 0, 0, 0},
		},
		{
			name: "Should solve inequation correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{1, 1, -1, -2, 6},
					{-1, -1, -1, 1, -5},
					{2, -1, 3, 4, 10},
					{-1, -2, 1, 1, 0},
				},
			},
			expected: matrix.Row{5, 0, 0, 0},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, _, _ := solveWithSupportSolution(tt.m)
			if !slices.Equal(actual, tt.expected) {
				t.Fatalf("Expected %v, \ngot %v", tt.expected, actual)
			}
		})
	}
}

func TestSolveWithOptimalSolution(t *testing.T) {
	tc := []struct {
		name     string
		m        matrix.Matrix
		expected matrix.Row
	}{
		{
			name: "Should solve inequation correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{-1, 0, 3, -2, 1, 3},
					{1, -1, 0, 1, 1, 3},
					{-1, -3, 1, 1, -1, -2},
					{-1, 1, 0, 0, 1, 0},
				},
			},
			expected: matrix.Row{3, 0, 0, 0, 0},
		},
		{
			name: "Should solve inequation correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{1, 1, -1, -2, 6},
					{-1, -1, -1, 1, -5},
					{2, -1, 3, 4, 10},
					{-1, -2, 1, 1, 0},
				},
			},
			expected: matrix.Row{0, 22, 0, 8},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, _ := SolveWithOptimalSolution(tt.m)
			if !slices.Equal(actual.Optimal, tt.expected) {
				t.Fatalf("Expected %v, \ngot %v", tt.expected, actual.Optimal)
			}
		})
	}
}
