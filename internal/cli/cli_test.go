package cli

import (
	"reflect"
	"testing"

	"github.com/hrvadl/algo/internal/inequations"
	"github.com/hrvadl/algo/internal/matrix"
)

func TestHandleGetMinWithOptimalSolution(t *testing.T) {
	tc := []struct {
		name     string
		m        matrix.Matrix
		expected matrix.Matrix
		max      float64
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

			_, solved, err := inequations.FindMaxWithSupportSolution(m)
			if err != nil {
				t.Fatal(err)
			}

			m = *solved

			actual, err := inequations.FindMaxWithOptimalSolution(m)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(actual.Matrix, tt.expected) {
				t.Fatalf("expected %v,\ngot %v", tt.expected, actual.Matrix)
			}
		})
	}
}
