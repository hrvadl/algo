package inequations

import (
	"github.com/hrvadl/algo/internal/matrix"
)

func FindMinIntegerSolution(m matrix.Matrix) (*MinSolution, *Solution, error) {
	optimal, support, err := FindIntegerSolution(m)
	if err != nil {
		return nil, nil, err
	}

	lastCol := len(optimal.Matrix.Rows[0]) - 1
	lastRow := len(optimal.Matrix.Rows) - 1

	return &MinSolution{
		Solution: *optimal,
		Min:      matrix.RoundTo(-1*optimal.Matrix.Rows[lastRow][lastCol], 2),
	}, support, nil
}

func FindMaxIntegerSolution(m matrix.Matrix) (*MaxSolution, *Solution, error) {
	optimal, support, err := FindIntegerSolution(m)
	if err != nil {
		return nil, nil, err
	}

	lastCol := len(optimal.Matrix.Rows[0]) - 1
	lastRow := len(optimal.Matrix.Rows) - 1

	return &MaxSolution{
		Solution: *optimal,
		Max:      matrix.RoundTo(optimal.Matrix.Rows[lastRow][lastCol], 2),
	}, support, nil
}

func FindIntegerSolution(m matrix.Matrix) (*Solution, *Solution, error) {
	support, err := FindSupportSolution(m)
	if err != nil {
		return nil, nil, err
	}

	optimal, err := FindOptimalSolution(support.Matrix)
	if err != nil {
		return nil, nil, err
	}

	lastCol := len(optimal.Matrix.Rows[0]) - 1
	for row, variable := range optimal.Matrix.LeftTitle {
		el := matrix.RoundTo(optimal.Matrix.Rows[row][lastCol], 2)
		if variable.IsX() && el != float64(int(el)) {
			l := m.NegativeRowFor(optimal.Matrix.IntegerLimitationFor(row))
			withLimit := optimal.Matrix.InsertRow(l)
			return FindIntegerSolution(withLimit.Round())
		}
	}

	return optimal, &support, nil
}
