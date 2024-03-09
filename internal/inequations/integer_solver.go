package inequations

import (
	"github.com/hrvadl/algo/internal/matrix"
)

func FindIntegerSolution(m matrix.Matrix) (*Solution, error) {
	support, err := FindSupportSolution(m)
	if err != nil {
		return nil, err
	}

	optimal, err := FindOptimalSolution(support.Matrix)
	if err != nil {
		return nil, err
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

	return optimal, nil
}
