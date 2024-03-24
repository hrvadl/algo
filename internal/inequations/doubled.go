package inequations

import "github.com/hrvadl/algo/internal/matrix"

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

func FindMaxDoubledWithSupportSolution(m matrix.Matrix) (*DoubledSolution, error) {
	support, err := FindSupportSolution(m)
	if err != nil {
		return nil, err
	}

	res := DoubledSolution{
		Max: support,
		Min: Solution{
			Matrix: support.Matrix,
			Result: make([]float64, support.Matrix.GetUCount()),
		},
	}

	lastRow := len(support.Matrix.Rows) - 1
	for col, variable := range support.Matrix.TopTitle {
		if variable.IsU() {
			res.Min.Result[variable.SecondStageIndex] = support.Matrix.Rows[lastRow][col]
		}
	}

	return &res, nil
}

func FindMinDoubledWithSupportSolution(m matrix.Matrix) (*DoubledSolution, error) {
	support, err := FindMinWithSupportSolution(m)
	if err != nil {
		return nil, err
	}

	res := DoubledSolution{
		Min: support,
		Max: Solution{
			Matrix: support.Matrix,
			Result: make([]float64, support.Matrix.GetXCount()),
		},
	}

	lastRow := len(support.Matrix.Rows) - 1
	for col, variable := range support.Matrix.TopTitle {
		if variable.IsU() {
			res.Max.Result[variable.SecondStageIndex] = support.Matrix.Rows[lastRow][col]
		}
	}

	return &res, nil
}
