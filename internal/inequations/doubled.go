package inequations

import (
	"github.com/hrvadl/algo/internal/matrix"
)

type DoubledOptimalSolution struct {
	MinSolution
	MaxSolution
}

type DoubledSolution struct {
	Min Solution
	Max Solution
}

func FindMinDoubledWithOptimalSolution(m matrix.Matrix) (*DoubledOptimalSolution, error) {
	optimal, err := FindMinWithOptimalSolution(m)
	if err != nil {
		return nil, err
	}

	res := DoubledOptimalSolution{
		MinSolution: *optimal,
		MaxSolution: MaxSolution{
			Solution: Solution{
				Matrix: optimal.Matrix,
				Result: make([]float64, optimal.Matrix.GetUCount()),
			},
			Max: optimal.Min,
		},
	}

	lastRow := len(optimal.Matrix.Rows) - 1
	for col, variable := range optimal.Matrix.TopTitle {
		if variable.IsU() {
			res.MaxSolution.Result[variable.SecondStageIndex] = optimal.Matrix.Rows[lastRow][col]
		}
	}

	return &res, nil
}

func FindMaxDoubledWithOptimalSolution(m matrix.Matrix) (*DoubledOptimalSolution, error) {
	optimal, err := FindMaxWithOptimalSolution(m)
	if err != nil {
		return nil, err
	}

	res := DoubledOptimalSolution{
		MaxSolution: *optimal,
		MinSolution: MinSolution{
			Solution: Solution{
				Matrix: optimal.Matrix,
				Result: make([]float64, optimal.Matrix.GetUCount()),
			},
			Min: optimal.Max,
		},
	}

	lastRow := len(optimal.Matrix.Rows) - 1
	for col, variable := range optimal.Matrix.TopTitle {
		if variable.IsU() {
			res.MinSolution.Result[variable.SecondStageIndex] = optimal.Matrix.Rows[lastRow][col]
		}
	}

	return &res, nil
}
