package inequations

import (
	"fmt"

	"github.com/hrvadl/algo/internal/matrix"
)

type Solution struct {
	Matrix matrix.Matrix
	Result []float64
}

type MaxSolution struct {
	Solution
	Max float64
}

type MinSolution struct {
	Solution
	Min float64
}

func FindMinWithOptimalSolution(m matrix.Matrix) (*MinSolution, error) {
	sol, err := FindMaxWithOptimalSolution(m)
	if err != nil {
		return nil, err
	}

	return &MinSolution{
		Solution: sol.Solution,
		Min:      -1 * sol.Max,
	}, nil
}

func FindMinWithSupportSolution(m matrix.Matrix) (Solution, error) {
	lastRow := len(m.Rows) - 1
	for i := range m.Rows[lastRow] {
		m.Rows[lastRow][i] /= -1
	}

	return FindSupportSolution(m)
}

func FindMaxWithOptimalSolution(m matrix.Matrix) (*MaxSolution, error) {
	optimal, err := FindOptimalSolution(m)
	if err != nil {
		return nil, err
	}

	lastCol := len(optimal.Matrix.Rows[0]) - 1
	lastRow := len(optimal.Matrix.Rows) - 1
	return &MaxSolution{
		Solution: *optimal,
		Max:      optimal.Matrix.Rows[lastRow][lastCol],
	}, nil
}

func FindOptimalSolution(m matrix.Matrix) (*Solution, error) {
	lastCol := len(m.Rows[0]) - 1
	res := make([]float64, m.GetXCount())

	col, err := m.FirstNegativeColumnInLastRow()
	if err != nil {
		for row, variable := range m.LeftTitle {
			if variable.IsX() {
				res[variable.FirstStageIndex] = matrix.RoundTo(m.Rows[row][lastCol], 2)
			}
		}

		return &Solution{
			Matrix: m,
			Result: res,
		}, nil
	}

	row, err := m.FindMinPositiveFor(col)
	if err != nil {
		return nil, err
	}

	m, err = m.JordanEliminateModified(col, row)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\nElement col: %v row: %v, Matrix:\n\n", col, row)
	rm := m.Round()
	rm.Print()

	return FindOptimalSolution(m)
}

func FindSupportSolution(m matrix.Matrix) (Solution, error) {
	lastCol := len(m.Rows[0]) - 1
	res := make([]float64, m.GetXCount())

	negativeInLastCol, err := m.FirstNegativeRowInLastColumn()
	if err != nil {
		for row, variable := range m.LeftTitle {
			if variable.IsX() {
				res[variable.FirstStageIndex] = m.Rows[row][lastCol]
			}
		}
		return Solution{Matrix: m, Result: res}, nil
	}

	col, err := m.FirstNegativeInRowExceptLastColumn(negativeInLastCol)
	if err != nil {
		return Solution{}, err
	}

	row, err := m.FindMinPositiveFor(col)
	if err != nil {
		return Solution{}, err
	}

	m, err = m.JordanEliminateModified(col, row)
	if err != nil {
		return Solution{}, err
	}

	fmt.Printf("\nElement col: %v row: %v, Matrix:\n\n", col, row)
	rm := m.Round()
	rm.Print()
	return FindSupportSolution(m)
}
