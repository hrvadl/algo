package equations

import "github.com/hrvadl/algo/internal/matrix"

type EquationSystem struct {
	A matrix.Matrix
	B matrix.Matrix
}

func SolveSystem(s EquationSystem) []float64 {
	swapped, _ := s.A.SwapAll()
	res := make([]float64, 0, len(swapped.Rows))

	for _, row := range swapped.Rows {
		var x float64
		for i, el := range row {
			x += el * s.B.Rows[i][0]
		}
		res = append(res, matrix.RoundTo(x, 2))
	}

	return res
}
