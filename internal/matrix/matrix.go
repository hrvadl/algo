package matrix

import (
	"errors"
	"fmt"
	"math"
)

type Row = []float64

type Matrix struct {
	Rows []Row
}

func (m *Matrix) Rank() int {
	_, rank, _ := m.SwapAll()
	return rank
}

func (m *Matrix) SwapAll() (Matrix, int, error) {
	var rank int
	resm := *m
	for i := 0; i < len(resm.Rows) && i < len(resm.Rows[i]); i++ {
		var err error
		if resm, err = resm.JordanEliminate(i, i); err == nil {
			rank++
		}
	}
	return resm, rank, nil
}

func (m *Matrix) Invert() (Matrix, error) {
	if !m.IsSquare() {
		return Matrix{}, errors.New("cannot inverse not square matrix")
	}

	if m.IsDegenerate() {
		return Matrix{}, errors.New("cannot inverse degenerate matrix")
	}

	resm := *m
	for i := range m.Rows {
		var err error
		if resm, err = resm.JordanEliminate(i, i); err != nil {
			return resm, err
		}
	}

	return resm, nil
}

func (m *Matrix) JordanEliminate(col, row int) (Matrix, error) {
	resm := m.Copy()
	eliminated := resm.Rows[row][col]
	if eliminated == 0 {
		return Matrix{}, errors.New("divide by zero")
	}

	resm.Rows[row][col] = 1

	for i := range resm.Rows[row] {
		if i != col {
			resm.Rows[row][i] *= -1
		}
	}

	for i, rowRes := range resm.Rows {
		for j := range rowRes {
			if i != row && j != col {
				resm.Rows[i][j] = m.Rows[i][j]*eliminated - m.Rows[i][col]*m.Rows[row][j]
			}
		}
	}

	return resm.DivideBy(eliminated)
}

func (m *Matrix) IsDegenerate() bool {
	return m.Determinant() == 0
}

func (m *Matrix) IsSquare() bool {
	col, row := m.GetDimensions()
	return col == row
}

func (m *Matrix) Determinant() float64 {
	if !m.IsSquare() {
		return 0
	}

	if len(m.Rows) < 1 {
		return 0
	}

	if len(m.Rows) == 2 {
		return m.Rows[0][0]*m.Rows[1][1] - m.Rows[0][1]*m.Rows[1][0]
	}

	var (
		det float64
		col int
	)

	for row := range m.Rows[0] {
		coef := m.Rows[row][col]
		sign := m.GetSignForAlgebraicAddition(col, row)
		minor := m.MinorFor(col, row)
		det += coef * sign * minor.Determinant()
	}

	return det
}

func (m *Matrix) GetDimensions() (col, rows int) {
	rows = len(m.Rows)
	if rows > 0 {
		col = len(m.Rows[0])
	}
	return col, rows
}

func (m *Matrix) GetSignForAlgebraicAddition(col, row int) float64 {
	if (col+row)%2 == 0 {
		return 1
	}
	return -1
}

func (m *Matrix) MinorFor(col, row int) Matrix {
	res := Matrix{
		Rows: make([][]float64, len(m.Rows)-1),
	}

	for i := 0; i < len(m.Rows); i++ {
		for j := 0; j < len(m.Rows[i]); j++ {
			if j == col || i == row {
				continue
			}

			var correction int
			if i > row {
				correction = -1
			}

			res.Rows[i+correction] = append(res.Rows[i+correction], m.Rows[i][j])
		}
	}

	return res
}

func (m *Matrix) Copy() Matrix {
	res := Matrix{
		Rows: make([][]float64, len(m.Rows)),
	}

	for i, row := range m.Rows {
		dst := make(Row, len(row))
		copy(dst, row)
		res.Rows[i] = dst
	}

	return res
}

func (m *Matrix) DivideBy(n float64) (Matrix, error) {
	if n == 0 {
		return Matrix{}, errors.New("divide by zero")
	}

	resm := m.Copy()
	for i, row := range m.Rows {
		for j := range row {
			resm.Rows[i][j] /= n
		}
	}

	return resm, nil
}

func (m *Matrix) Round() Matrix {
	resm := m.Copy()
	for row, el := range resm.Rows {
		for col := range el {
			resm.Rows[row][col] = RoundTo(resm.Rows[row][col], 2)
		}
	}

	return resm
}

func (m *Matrix) Print() {
	for _, row := range m.Rows {
		for _, col := range row {
			var space string
			if col >= 0 {
				space = " "
			}
			fmt.Printf("%s%v   ", space, col)
		}
		fmt.Println()
	}

	fmt.Println()
}

func RoundTo(num float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(num*ratio) / ratio
}
