package games

import (
	"reflect"
	"slices"
	"testing"

	"github.com/hrvadl/algo/internal/matrix"
)

func TestWithMaxMinStrategy(t *testing.T) {
	tc := []struct {
		name       string
		m          matrix.Matrix
		strategies []int
	}{
		{
			name: "Should find max min strategy solution correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{-1, 1, 1, 4},
					{-1, -2, 2, 3},
					{3, -1, 3, 2},
				},
			},
			strategies: []int{0, 2},
		},
		{
			name: "Should find max min strategy solution correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{-2, 1, 1, 4},
					{-1, -2, 2, 3},
					{3, -1, 3, 2},
				},
			},
			strategies: []int{2},
		},
		{
			name:       "Should return nil when rows are empty",
			m:          matrix.Matrix{},
			strategies: nil,
		},
	}

	for _, tt := range tc {
		if got := SolveWithMaxMinStrategy(tt.m); !reflect.DeepEqual(tt.strategies, got) {
			t.Errorf("Expected to get: %v, got: %v", tt.strategies, got)
		}
	}
}

func TestWithMaxMaxStrategy(t *testing.T) {
	tc := []struct {
		name       string
		m          matrix.Matrix
		strategies []int
	}{
		{
			name: "Should find max max strategy solution correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{-1, 1, 1, 4},
					{-1, -2, 2, 3},
					{3, -1, 3, 2},
				},
			},
			strategies: []int{0},
		},
		{
			name: "Should find max max strategy solution correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{-1, 1, 1, 4},
					{-1, -2, 2, 3},
					{3, -1, 3, 4},
				},
			},
			strategies: []int{0, 2},
		},
		{
			name: "Should find max max strategy solution correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{-1, 1, 1, 4},
					{-1, -2, 2, 9},
					{3, -1, 3, 4},
				},
			},
			strategies: []int{1},
		},
		{
			name:       "Should return nil when rows are empty",
			m:          matrix.Matrix{},
			strategies: nil,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := SolveWithMaxMaxStrategy(tt.m); !reflect.DeepEqual(tt.strategies, got) {
				t.Errorf("Expected to get: %v, got: %v", tt.strategies, got)
			}
		})
	}
}

func TestWithPessimistOptimistStrategy(t *testing.T) {
	tc := []struct {
		name       string
		m          matrix.Matrix
		strategies []int
		y          float64
	}{
		{
			name: "Should find pes/opt strategy solution correctly",
			y:    0.3,
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{-1, 1, 1, 4},
					{-1, -2, 2, 3},
					{3, -1, 3, 2},
				},
			},
			strategies: []int{0},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := SolveWithPessimismOptimismStrategy(tt.m, tt.y); !reflect.DeepEqual(
				tt.strategies,
				got,
			) {
				t.Errorf("Expected to get: %v, got: %v", tt.strategies, got)
			}
		})
	}
}

func TestWithMinMaxStrategy(t *testing.T) {
	tc := []struct {
		name       string
		m          matrix.Matrix
		strategies []int
	}{
		{
			name: "Should find min max strategy solution correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{-1, 1, 1, 4},
					{-1, -2, 2, 3},
					{3, -1, 3, 2},
				},
			},
			strategies: []int{2},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := SolveWithMinMaxStrategy(tt.m); !reflect.DeepEqual(
				tt.strategies,
				got,
			) {
				t.Errorf("Expected to get: %v, got: %v", tt.strategies, got)
			}
		})
	}
}

func TestSolveWithBayesPrinciple(t *testing.T) {
	tc := []struct {
		name     string
		m        matrix.Matrix
		p        matrix.Row
		expected []int
	}{
		{
			name: "Should solve with bayes principle correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{-1, 1, 1, 4},
					{-1, -2, 2, 3},
					{3, -1, 3, 2},
				},
			},
			p:        matrix.Row{0.2, 0.4, 0.1, 0.3},
			expected: []int{0},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got, _ := SolveWithBayesPrinicple(tt.m, tt.p); !slices.Equal(tt.expected, got) {
				t.Fatalf("Expected to get: %v, got: %v", tt.expected, got)
			}
		})
	}
}

func TestSolveWithLaplasPrinciple(t *testing.T) {
	tc := []struct {
		name     string
		m        matrix.Matrix
		v        float64
		expected []int
	}{
		{
			name: "Should solve with bayes principle correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{-1, 1, 1, 4},
					{-1, -2, 2, 3},
					{3, -1, 3, 2},
				},
			},
			v:        1. / 4,
			expected: []int{2},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got, _ := SolveWithLaplasPrinciple(tt.m, tt.v); !slices.Equal(tt.expected, got) {
				t.Fatalf("Expected to get: %v, got: %v", tt.expected, got)
			}
		})
	}
}
