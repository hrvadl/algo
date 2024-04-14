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

	minResLen := optimal.Matrix.InitialCols
	if minResLen == 0 {
		minResLen = optimal.Matrix.GetUCount()
	}

	res := DoubledOptimalSolution{
		MinSolution: *optimal,
		MaxSolution: MaxSolution{
			Solution: Solution{
				Matrix: optimal.Matrix,
				Result: make([]float64, minResLen),
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

	maxResLen := optimal.Matrix.InitialCols
	if maxResLen == 0 {
		maxResLen = optimal.Matrix.GetUCount()
	}

	res := DoubledOptimalSolution{
		MaxSolution: *optimal,
		MinSolution: MinSolution{
			Solution: Solution{
				Matrix: optimal.Matrix,
				Result: make([]float64, maxResLen),
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
