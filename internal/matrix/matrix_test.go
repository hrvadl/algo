package matrix

import (
	"reflect"
	"testing"
)

func TestCalculateDeterminant(t *testing.T) {
	tc := []struct {
		name     string
		m        Matrix
		expected float64
	}{
		{
			name:     "Should calculate correctly",
			expected: 70,
			m: Matrix{
				Rows: []Row{
					{1, 4, 3, 2},
					{3, 2, 1, 1},
					{1, 4, 2, -3},
					{5, 2, -1, 0},
				},
			},
		},
		{
			name:     "Should calculate correctly",
			expected: -215,
			m: Matrix{
				Rows: []Row{
					{2, -2, 3, 1},
					{3, 0, 1, 5},
					{1, 3, 4, -2},
					{4, 2, 2, 1},
				},
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if actual := tt.m.Determinant(); actual != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestInvertMatrix(t *testing.T) {
	tc := []struct {
		name     string
		m        Matrix
		expected Matrix
	}{
		{
			name: "Should calculate correct result",
			m: Matrix{
				Rows: []Row{
					{6, 2, 5},
					{-3, 4, -1},
					{1, 4, 3},
				},
			},
			expected: Matrix{
				Rows: []Row{
					{0.50, 0.44, -0.69},
					{0.25, 0.41, -0.28},
					{-0.50, -0.69, 0.94},
				},
			},
		},
		{
			name: "Should calculate correct result",
			m: Matrix{
				Rows: []Row{
					{5, -3, 7},
					{-1, 4, 3},
					{6, -2, 5},
				},
			},
			expected: Matrix{
				Rows: []Row{
					{-0.28, -0.01, 0.4},
					{-0.25, 0.18, 0.24},
					{0.24, 0.09, -0.18},
				},
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, _ := tt.m.Invert()
			if actual := actual.Round(); !reflect.DeepEqual(actual.Rows, tt.expected.Rows) {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestJordanElimination(t *testing.T) {
	tc := []struct {
		name     string
		m        Matrix
		expected Matrix
	}{
		{
			name: "Should calculate correct result",
			m: Matrix{
				Rows: []Row{
					{6, 2, 5},
					{-3, 4, -1},
					{1, 4, 3},
				},
			},
			expected: Matrix{
				Rows: []Row{
					{0.17, -0.33, -0.83},
					{-0.50, 5, 1.50},
					{0.17, 3.67, 2.17},
				},
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, _ := tt.m.JordanEliminate(0, 0)
			if actual := actual.Round(); !reflect.DeepEqual(actual.Rows, tt.expected.Rows) {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestJordanEliminationModified(t *testing.T) {
	tc := []struct {
		name     string
		m        Matrix
		col      int
		row      int
		expected Matrix
	}{
		{
			name: "Should calculate correct result",
			col:  0,
			row:  1,
			m: Matrix{
				Rows: []Row{
					{1, 1, -1, -2, 6},
					{-1, -1, -1, 1, -5},
					{2, -1, 3, 4, 10},
					{-1, -2, 1, 1, 0},
				},
			},
			expected: Matrix{
				Rows: []Row{
					{1, 0, -2, -1, 1},
					{-1, 1, 1, -1, 5},
					{2, -3, 1, 6, 0},
					{-1, -1, 2, 0, 5},
				},
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, _ := tt.m.JordanEliminateModified(tt.col, tt.row)
			if actual := actual.Round(); !reflect.DeepEqual(actual.Rows, tt.expected.Rows) {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestCalculateRank(t *testing.T) {
	tc := []struct {
		name     string
		m        Matrix
		expected int
	}{
		{
			name: "Should calculate rank correctly",
			m: Matrix{
				Rows: []Row{
					{1, 2, 3, 4},
					{2, 4, 6, 8},
				},
			},
			expected: 1,
		},
		{
			name: "Should calculate rank correctly",
			m: Matrix{
				Rows: []Row{
					{2, 5, 4},
					{-3, 1, -2},
					{-1, 6, 2},
				},
			},
			expected: 2,
		},
		{
			name: "Should calculate rank correctly",
			m: Matrix{
				Rows: []Row{
					{1, 2},
					{3, 6},
					{5, 10},
					{4, 8},
				},
			},
			expected: 1,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			if actual := tt.m.Rank(); actual != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
