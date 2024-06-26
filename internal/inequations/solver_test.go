package inequations

import (
	"reflect"
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
			t.Parallel()
			actual, _ := FindSupportSolution(tt.m)
			if !slices.Equal(actual.Result, tt.expected) {
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
			support, err := FindSupportSolution(tt.m)
			if err != nil {
				t.Fatal(err)
			}

			actual, err := FindMaxWithOptimalSolution(support.Matrix)
			if err != nil {
				t.Fatal(err)
			}

			if !slices.Equal(actual.Result, tt.expected) {
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
			support, err := FindMinWithSupportSolution(tt.m)
			if err != nil {
				t.Fatal(err)
			}

			actual, err := FindMinWithOptimalSolution(support.Matrix)
			if err != nil {
				t.Fatal(err)
			}

			if !slices.Equal(actual.Result, tt.expected) {
				t.Fatalf("Expected %v, \ngot %v", tt.expected, actual.Result)
			}

			if tt.min != actual.Min {
				t.Fatalf("Expected %v, got %v", tt.min, actual.Min)
			}
		})
	}
}

func TestHandleGetMinWithOptimalSolution(t *testing.T) {
	tc := []struct {
		name     string
		m        matrix.Matrix
		expected matrix.Matrix
		max      float64
		solution []float64
	}{
		{
			name: "Should get max correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{-2, 1, 1, 3, 2},
					{-3, 2, -3, 0, 7},
					{-3, 1, 4, 1, 1},
					{3, -2, 2, -2, -9},
					{-10, 1, 42, 52, 0},
				},
				TopTitle: []matrix.Variable{
					{FirstStageName: "x", FirstStageIndex: 0},
					{FirstStageName: "x", FirstStageIndex: 1},
					{FirstStageName: "x", FirstStageIndex: 2},
					{FirstStageName: "x", FirstStageIndex: 3},
					{FirstStageName: "1"},
				},
				LeftTitle: []matrix.Variable{
					{FirstStageName: "0"},
					{FirstStageName: "0"},
					{FirstStageName: "y", FirstStageIndex: 0},
					{FirstStageName: "y", FirstStageIndex: 1},
					{FirstStageName: "z"},
				},
			},
			solution: []float64{9, 17, 0, 1},
			max:      21,
			expected: matrix.Matrix{
				Rows: []matrix.Row{
					{-3, -2, 9},
					{-4, 2, 10},
					{-4.5, -4.5, 17},
					{-0.5, 0.5, 1},
					{0.5, 0.5, 21},
				},
				TopTitle: []matrix.Variable{
					{FirstStageName: "y", FirstStageIndex: 1},
					{FirstStageName: "x", FirstStageIndex: 2},
					{FirstStageName: "1"},
				},
				LeftTitle: []matrix.Variable{
					{FirstStageName: "x", FirstStageIndex: 0},
					{FirstStageName: "y", FirstStageIndex: 0},
					{FirstStageName: "x", FirstStageIndex: 1},
					{FirstStageName: "x", FirstStageIndex: 3},
					{FirstStageName: "z"},
				},
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m, err := tt.m.DeleteZeros()
			if err != nil {
				t.Fatal(err)
			}

			support, err := FindSupportSolution(m)
			if err != nil {
				t.Fatal(err)
			}

			optimal, err := FindMaxWithOptimalSolution(support.Matrix)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(optimal.Matrix, tt.expected) {
				t.Fatalf("expected %v,\ngot %v", tt.expected, optimal.Matrix)
			}

			if !reflect.DeepEqual(optimal.Result, tt.solution) {
				t.Fatalf("expected %v,\ngot %v", tt.solution, optimal.Result)
			}

			if optimal.Max != tt.max {
				t.Fatalf("expected %v,\ngot %v", tt.max, optimal.Max)
			}
		})
	}
}
