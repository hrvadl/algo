package games

import (
	"errors"

	"github.com/hrvadl/algo/internal/matrix"
)

func CompleteMatrixToCompatible(m matrix.Matrix) (*matrix.Matrix, error) {
	m = m.Copy()
	if len(m.Rows) == 0 {
		return nil, errors.New("cannot complete empty matrix")
	}

	if len(m.Rows[0]) == 0 {
		return nil, errors.New("cannot complete empty matrix")
	}

	m.Rows = append(m.Rows, make(matrix.Row, 0, len(m.Rows)))
	lastRowIdx := len(m.Rows) - 1
	for range m.Rows[0] {
		m.Rows[lastRowIdx] = append(m.Rows[lastRowIdx], -1)
	}

	for i := range m.Rows {
		if i == lastRowIdx {
			m.Rows[i] = append(m.Rows[i], 0)
			continue
		}
		m.Rows[i] = append(m.Rows[i], 1)
	}

	return &m, nil
}

func GetGameWeight(m matrix.Matrix) float64 {
	return matrix.RoundTo(1/m.Rows[len(m.Rows)-1][len(m.Rows[0])-1], 2)
}

func CorrectGameWeight(w, minabs float64) float64 {
	return matrix.RoundTo(w-minabs, 2)
}

func CorrectMixedStrategy(strategy []float64, w float64) []float64 {
	res := make([]float64, len(strategy))
	for i := range strategy {
		res[i] = matrix.RoundTo(strategy[i]*w, 2)
	}
	return res
}
