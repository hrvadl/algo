package inequations

import (
	"fmt"

	"github.com/hrvadl/algo/internal/matrix"
)

type MaxSolution struct {
	Matrix   matrix.Matrix
	Solution []float64
	Max      float64
}

type MinSolution struct {
	Matrix   matrix.Matrix
	Solution []float64
	Min      float64
}

func FindMinWithOptimalSolution(m matrix.Matrix) (*MinSolution, error) {
	sol, err := FindMaxWithOptimalSolution(m)
	if err != nil {
		return nil, err
	}

	return &MinSolution{
		Matrix:   m,
		Solution: sol.Solution,
		Min:      -1 * sol.Max,
	}, nil
}

func FindMinWithSupportSolution(m matrix.Matrix) ([]float64, *matrix.Matrix, error) {
	lastRow := len(m.Rows) - 1
	for i := range m.Rows[lastRow] {
		m.Rows[lastRow][i] /= -1
	}

	return FindMaxWithSupportSolution(m)
}

func FindMaxWithOptimalSolution(m matrix.Matrix) (*MaxSolution, error) {
	lastCol := len(m.Rows[0]) - 1
	lastRow := len(m.Rows) - 1
	res := make([]float64, m.GetXCount())

	col, err := m.FirstNegativeColumnInLastRow()
	if err != nil {
		for row, variable := range m.LeftTitle {
			if variable.IsX() {
				res[variable.Index] = m.Rows[row][lastCol]
			}
		}

		return &MaxSolution{
			Matrix:   m,
			Solution: res,
			Max:      m.Rows[lastRow][lastCol],
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

	return FindMaxWithOptimalSolution(m)
}

func FindMaxWithSupportSolution(m matrix.Matrix) ([]float64, *matrix.Matrix, error) {
	lastCol := len(m.Rows[0]) - 1
	res := make([]float64, m.GetXCount())

	negativeInLastCol, err := m.FirstNegativeRowInLastColumn()
	if err != nil {
		for row, variable := range m.LeftTitle {
			if variable.IsX() {
				res[variable.Index] = m.Rows[row][lastCol]
			}
		}
		return res, &m, nil
	}

	col, err := m.FirstNegativeInRowExceptLastColumn(negativeInLastCol)
	if err != nil {
		return nil, nil, err
	}

	row, err := m.FindMinPositiveFor(col)
	if err != nil {
		return nil, nil, err
	}

	m, err = m.JordanEliminateModified(col, row)
	if err != nil {
		return nil, nil, err
	}

	fmt.Printf("\nElement col: %v row: %v, Matrix:\n\n", col, row)
	rm := m.Round()
	rm.Print()
	return FindMaxWithSupportSolution(m)
}
