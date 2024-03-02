package inequations

import (
	"errors"
	"fmt"

	"github.com/hrvadl/algo/internal/matrix"
)

type Solution struct {
	Optimal []float64
	Support []float64
}

func SolveWithOptimalSolution(m matrix.Matrix) (*Solution, error) {
	fmt.Printf("\nFinding the support solution...\n")
	support, solved, err := solveWithSupportSolution(m)
	if err != nil {
		return nil, err
	}

	m = *solved
	lastCol := len(solved.Rows[0]) - 1
	lastRow := len(solved.Rows) - 1
	res := make([]float64, lastCol)

	fmt.Printf("\nFinding the optimal solution...\n")
	for i := 0; i < len(m.Rows[lastRow])-1; i++ {
		if m.Rows[lastRow][i] < 0 {
			col := i

			row, err := findRowToEliminate(m, col)
			if err != nil {
				return nil, err
			}

			m, err = m.JordanEliminateModified(col, row)
			if err != nil {
				return nil, err
			}
			fmt.Printf("Step #%v. Matrix: \n\n", i+1)
			m.Print()
		}
	}

	for row, variable := range m.LeftTitle {
		if variable.IsX() {
			res[variable.Index] = m.Rows[row][lastCol]
		}
	}

	return &Solution{
		Optimal: res,
		Support: support,
	}, nil
}

func solveWithSupportSolution(m matrix.Matrix) ([]float64, *matrix.Matrix, error) {
	lastCol := len(m.Rows[0]) - 1
	res := make([]float64, lastCol)

	for i := 0; i < len(m.Rows)-1; i++ {
		if m.Rows[i][lastCol] < 0 {
			col, err := m.FirstNegativeInRow(i)
			if err != nil {
				return nil, nil, err
			}

			if col == lastCol {
				return nil, nil, fmt.Errorf("no negative numbers are found in the row %v", i)
			}

			row, err := findRowToEliminate(m, col)
			if err != nil {
				return nil, nil, err
			}

			m, err = m.JordanEliminateModified(col, row)
			if err != nil {
				return nil, nil, err
			}
			fmt.Printf("Step #%v. Matrix: \n\n", i+1)
			m.Print()
		}
	}

	for row, variable := range m.LeftTitle {
		if variable.IsX() {
			res[variable.Index] = m.Rows[row][lastCol]
		}
	}

	return res, &m, nil
}

func findRowToEliminate(m matrix.Matrix, col int) (row int, err error) {
	var min float64
	lastCol := len(m.Rows[0]) - 1
	for j := 0; j < len(m.Rows)-1; j++ {
		res := m.Rows[j][lastCol] / m.Rows[j][col]
		if res <= 0 {
			continue
		}

		if min == 0 || min > res {
			min = res
			row = j
		}
	}

	if min == 0 {
		return 0, errors.New("cannot find row to jordan eliminate")
	}

	return row, nil
}
