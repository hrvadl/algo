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
					{FirstStageName: "y", FirstStageIndex: 0},
					{FirstStageName: "y", FirstStageIndex: 1},
					{FirstStageName: "z"},
				},
				TopTitle: []matrix.Variable{
					{FirstStageName: "x", FirstStageIndex: 0},
					{FirstStageName: "x", FirstStageIndex: 1},
					{FirstStageName: "1"},
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
					{FirstStageName: "y", FirstStageIndex: 0},
					{FirstStageName: "x", FirstStageIndex: 1},
					{FirstStageName: "x", FirstStageIndex: 0},
					{FirstStageName: "z"},
				},
				TopTitle: []matrix.Variable{
					{FirstStageName: "s", FirstStageIndex: 0},
					{FirstStageName: "y", FirstStageIndex: 1},
					{FirstStageName: "1"},
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
					{FirstStageName: "y", FirstStageIndex: 0},
					{FirstStageName: "y", FirstStageIndex: 1},
					{FirstStageName: "y", FirstStageIndex: 2},
					{FirstStageName: "z"},
				},
				TopTitle: []matrix.Variable{
					{FirstStageName: "x", FirstStageIndex: 0},
					{FirstStageName: "x", FirstStageIndex: 1},
					{FirstStageName: "x", FirstStageIndex: 2},
					{FirstStageName: "1"},
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
					{FirstStageName: "x", FirstStageIndex: 0},
					{FirstStageName: "x", FirstStageIndex: 1},
					{FirstStageName: "x", FirstStageIndex: 2},
					{FirstStageName: "y", FirstStageIndex: 1},
					{FirstStageName: "s"},
					{FirstStageName: "z"},
				},
				TopTitle: []matrix.Variable{
					{FirstStageName: "s"},
					{FirstStageName: "y", FirstStageIndex: 0},
					{FirstStageName: "y", FirstStageIndex: 2},
					{FirstStageName: "1"},
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
