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
				TopTitle: map[int]matrix.Variable{
					0: {Name: "x", Index: 0},
					1: {Name: "x", Index: 1},
					2: {Name: "x", Index: 2},
					3: {Name: "x", Index: 3},
					4: {Name: "1"},
				},
				LeftTitle: map[int]matrix.Variable{
					0: {Name: "0"},
					1: {Name: "0"},
					2: {Name: "y", Index: 0},
					3: {Name: "y", Index: 1},
					4: {Name: "z"},
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
				TopTitle: map[int]matrix.Variable{
					0: {Name: "y", Index: 1},
					1: {Name: "x", Index: 2},
					2: {Name: "1"},
				},
				LeftTitle: map[int]matrix.Variable{
					0: {Name: "x", Index: 0},
					1: {Name: "y", Index: 0},
					2: {Name: "x", Index: 1},
					3: {Name: "x", Index: 3},
					4: {Name: "z"},
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

			_, solved, err := FindMaxWithSupportSolution(m)
			if err != nil {
				t.Fatal(err)
			}

			m = *solved
			actual, err := FindMaxWithOptimalSolution(m)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(actual.Matrix, tt.expected) {
				t.Fatalf("expected %v,\ngot %v", tt.expected, actual.Matrix)
			}

			if !reflect.DeepEqual(actual.Solution, tt.solution) {
				t.Fatalf("expected %v,\ngot %v", tt.solution, actual.Solution)
			}

			if actual.Max != tt.max {
				t.Fatalf("expected %v,\ngot %v", tt.max, actual.Max)
			}
		})
	}
}
