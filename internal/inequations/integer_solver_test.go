package inequations

import (
	"reflect"
	"testing"

	"github.com/hrvadl/algo/internal/matrix"
)

func TestFindIntegerSolution(t *testing.T) {
	tc := []struct {
		name     string
		m        matrix.Matrix
		expected matrix.Matrix
	}{
		{
			name: "Should find integer solution correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{2, 1, 6},
					{1, 3, 4},
					{-1, -4, 0},
				},
				LeftTitle: []matrix.Variable{
					{Name: "y", Index: 0},
					{Name: "y", Index: 1},
					{Name: "z"},
				},
				TopTitle: []matrix.Variable{
					{Name: "x", Index: 0},
					{Name: "x", Index: 1},
					{Name: "1"},
				},
			},
			expected: matrix.Matrix{
				Rows: []matrix.Row{
					{5.06, -2, 3},
					{1, 0, 1},
					{-3.03, 1, 1},
					{1, 1, 5},
				},
				LeftTitle: []matrix.Variable{
					{Name: "y", Index: 0},
					{Name: "x", Index: 1},
					{Name: "x", Index: 0},
					{Name: "z"},
				},
				TopTitle: []matrix.Variable{
					{Name: "s", Index: 0},
					{Name: "y", Index: 1},
					{Name: "1"},
				},
			},
		},
		{
			name: "Should find integer solution correctly",
			m: matrix.Matrix{
				Rows: []matrix.Row{
					{3, 2, 0, 10},
					{1, 4, 0, 11},
					{3, 3, 1, 13},
					{-4, -5, -1, 0},
				},
				LeftTitle: []matrix.Variable{
					{Name: "y", Index: 0},
					{Name: "y", Index: 1},
					{Name: "y", Index: 2},
					{Name: "z"},
				},
				TopTitle: []matrix.Variable{
					{Name: "x", Index: 0},
					{Name: "x", Index: 1},
					{Name: "x", Index: 2},
					{Name: "1"},
				},
			},
			expected: matrix.Matrix{
				Rows: []matrix.Row{
					{-0.18, 0.45, 0, 2},
					{0.27, -0.18, 0, 2},
					{-0.27, -0.82, 1, 1},
					{-0.91, 0.27, 0, 1},
					{-0.73, -0.18, 0, 0},
					{0.36, 0.09, 1, 19},
				},
				LeftTitle: []matrix.Variable{
					{Name: "x", Index: 0},
					{Name: "x", Index: 1},
					{Name: "x", Index: 2},
					{Name: "y", Index: 1},
					{Name: "s"},
					{Name: "z"},
				},
				TopTitle: []matrix.Variable{
					{Name: "s"},
					{Name: "y", Index: 0},
					{Name: "y", Index: 2},
					{Name: "1"},
				},
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, _, err := FindIntegerSolution(tt.m)
			if err != nil {
				t.Fatal(err)
			}

			actual.Matrix = actual.Matrix.Round()
			if !reflect.DeepEqual(actual.Matrix, tt.expected) {
				t.Fatalf("expected %v\ngot %v", tt.expected, actual.Matrix)
			}
		})
	}
}
