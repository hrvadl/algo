package inequations

import (
	"slices"
	"testing"

	"github.com/hrvadl/algo/internal/matrix"
)

func TestFindMaxWithSupportSolution(t *testing.T) {
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
			actual, _, _ := FindMaxWithSupportSolution(tt.m)
			if !slices.Equal(actual, tt.expected) {
				t.Fatalf("Expected %v, \ngot %v", tt.expected, actual)
			}
		})
	}
}

func TestFindMaxWithOptimalSolution(t *testing.T) {
	tc := []struct {
		name     string
		m        matrix.Matrix
		expected matrix.Row
		max      float64
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
			max:      3,
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
			max:      36,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			_, m, err := FindMaxWithSupportSolution(tt.m)
			if err != nil {
				t.Fatal(err)
			}

			actual, err := FindMaxWithOptimalSolution(*m)
			if err != nil {
				t.Fatal(err)
			}

			if !slices.Equal(actual.Solution, tt.expected) {
				t.Errorf("Expected %v, \ngot %v", tt.expected, actual.Solution)
			}

			if tt.max != actual.Max {
				t.Errorf("Expected %v, got %v", tt.max, actual.Max)
			}
		})
	}
}

func TestFindMinWithOptimalSolution(t *testing.T) {
	tc := []struct {
		name     string
		m        matrix.Matrix
		expected matrix.Row
		min      float64
	}{
		{
			name: "Should solve inequation correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{1, 1, -1, -2, 6},
					{-1, -1, -1, 1, -5},
					{2, -1, 3, 4, 10},
					{2, -3, 0, 3, 0},
				},
			},
			expected: matrix.Row{5, 0, 0, 0},
			min:      -10,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, m, err := FindMinWithSupportSolution(tt.m)
			if err != nil {
				t.Fatal(err)
			}

			actual, err := FindMinWithOptimalSolution(*m)
			if err != nil {
				t.Fatal(err)
			}

			if !slices.Equal(actual.Solution, tt.expected) {
				t.Fatalf("Expected %v, \ngot %v", tt.expected, actual.Solution)
			}

			if tt.min != actual.Min {
				t.Fatalf("Expected %v, got %v", tt.min, actual.Min)
			}
		})
	}
}
