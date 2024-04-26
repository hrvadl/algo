package games

import (
	"errors"

	"github.com/hrvadl/algo/internal/matrix"
	"github.com/hrvadl/algo/pkg/sliceh"
)

func SolveWithMaxMinStrategy(m matrix.Matrix) []int {
	return sliceh.MaxIdxs(sliceh.MinFor2D(m.Rows))
}

func SolveWithMaxMaxStrategy(m matrix.Matrix) []int {
	return sliceh.MaxIdxs(sliceh.MaxFor2D(m.Rows))
}

func SolveWithPessimismOptimismStrategy(m matrix.Matrix, y float64) []int {
	mins := sliceh.MinFor2D(m.Rows)
	maxs := sliceh.MaxFor2D(m.Rows)
	res := make(matrix.Row, len(maxs))
	for i, min := range mins {
		res[i] = y*min + (1-y)*maxs[i]
	}
	return sliceh.MaxIdxs(res)
}

func SolveWithMinMaxStrategy(m matrix.Matrix) []int {
	r := m.NewRisk()
	return sliceh.MinIdxs(sliceh.MaxFor2D(r.Rows))
}

func SolveWithLaplasPrinciple(m matrix.Matrix, v float64) ([]int, error) {
	if len(m.Rows) == 0 {
		return nil, errors.New("p slice should be same length as matrix row")
	}

	p := make([]float64, 0, len(m.Rows))
	for range m.Rows[0] {
		p = append(p, v)
	}

	return SolveWithBayesPrinicple(m, p)
}

func SolveWithBayesPrinicple(m matrix.Matrix, p matrix.Row) ([]int, error) {
	multiplied, err := m.MultiplyByVector(p)
	if err != nil {
		return nil, err
	}

	return sliceh.MaxIdxs(multiplied.SumRows()), nil
}
