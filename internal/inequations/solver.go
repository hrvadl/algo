package inequations

import (
	"errors"
	"fmt"

	"github.com/hrvadl/algo/internal/matrix"
)

type MaxSolution struct {
	Solution []float64
	Max      float64
}

type MinSolution struct {
	Solution []float64
	Min      float64
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
	res := make([]float64, lastCol)
	col := -1

	fmt.Printf("\nFinding the optimal solution...\n")
	for i := 0; i < len(m.Rows[lastRow])-1; i++ {
		if m.Rows[lastRow][i] < 0 {
			col = i
			break
		}
	}

	if col == -1 {
		for row, variable := range m.LeftTitle {
			if variable.IsX() {
				res[variable.Index] = m.Rows[row][lastCol]
			}
		}

		return &MaxSolution{
			Solution: res,
			Max:      m.Rows[lastRow][lastCol],
		}, nil
	}

	row, err := findRowToEliminate(m, col)
	if err != nil {
		return nil, err
	}

	m, err = m.JordanEliminateModified(col, row)
	if err != nil {
		return nil, err
	}

	rm := m.Round()
	rm.Print()

	return FindMaxWithOptimalSolution(m)
}

func FindMaxWithSupportSolution(m matrix.Matrix) ([]float64, *matrix.Matrix, error) {
	lastCol := len(m.Rows[0]) - 1
	res := make([]float64, lastCol)
	negativeInLastCol := -1

	for i := 0; i < len(m.Rows)-1; i++ {
		if m.Rows[i][lastCol] < 0 {
			negativeInLastCol = i
			break
		}
	}

	if negativeInLastCol == -1 {
		for row, variable := range m.LeftTitle {
			if variable.IsX() {
				res[variable.Index] = m.Rows[row][lastCol]
			}
		}
		return res, &m, nil
	}

	col, err := m.FirstNegativeInRow(negativeInLastCol)
	if err != nil {
		return nil, nil, err
	}

	if col == lastCol {
		return nil, nil, fmt.Errorf(
			"no negative numbers are found in the row %v",
			negativeInLastCol,
		)
	}

	row, err := findRowToEliminate(m, col)
	if err != nil {
		return nil, nil, err
	}

	m, err = m.JordanEliminateModified(col, row)
	if err != nil {
		return nil, nil, err
	}

	rm := m.Round()
	rm.Print()
	return FindMaxWithSupportSolution(m)
}

func findRowToEliminate(m matrix.Matrix, col int) (row int, err error) {
	min := 0.
	row = -1
	lastCol := len(m.Rows[0]) - 1
	for j := 0; j < len(m.Rows)-1; j++ {
		if m.Rows[j][col] == 0 {
			continue
		}

		res := m.Rows[j][lastCol] / m.Rows[j][col]
		if res < 0 {
			continue
		}

		if res == 0 && m.Rows[j][col] < 0 {
			continue
		}

		if (min == 0 && row == -1) || min > res {
			min = res
			row = j
		}
	}

	if row == -1 {
		return 0, errors.New("cannot find element to jordan eliminate")
	}

	return row, nil
}
