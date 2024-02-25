package equations

import (
	"reflect"
	"testing"

	"github.com/hrvadl/algo/internal/matrix"
)

func TestSolveSystem(t *testing.T) {
	tc := []struct {
		name     string
		sys      EquationSystem
		expected []float64
	}{
		{
			name: "Should calculate correctly",
			sys: EquationSystem{
				A: matrix.Matrix{
					Rows: []matrix.Row{
						{5, -3, 7},
						{-1, 4, 3},
						{6, -2, 5},
					},
				},
				B: matrix.Matrix{
					Rows: []matrix.Row{
						{13},
						{13},
						{12},
					},
				},
			},
			expected: []float64{1, 2, 2},
		},
		{
			name: "Should calculate correctly",
			sys: EquationSystem{
				A: matrix.Matrix{
					Rows: []matrix.Row{
						{6, 2, 5},
						{-3, 4, -1},
						{1, 4, 3},
					},
				},
				B: matrix.Matrix{
					Rows: []matrix.Row{
						{1},
						{6},
						{6},
					},
				},
			},
			expected: []float64{-1, 1, 1},
		},
	}

	for _, tt := range tc {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if actual := SolveSystem(tt.sys); !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, actual)
			}
		})
	}
}
