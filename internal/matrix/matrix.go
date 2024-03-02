package matrix

import (
	"errors"
	"fmt"
	"math"
)

type Row = []float64

type Variable struct {
	Name  string
	Index int
}

func (v Variable) IsX() bool {
	return v.Name == "x"
}

type Matrix struct {
	Rows      []Row
	LeftTitle map[int]Variable
	TopTitle  map[int]Variable
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
		eliminated := RoundTo(resm.Rows[i][i], 2)
		if resm, err = resm.JordanEliminate(i, i); err == nil {
			fmt.Printf("\nStep #%v. Element: %v. Results: \n", i+1, eliminated)
			rounded := resm.Round()
			rounded.Print()
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
	resm.SetSwapped(col, row)
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

func (m *Matrix) JordanEliminateModified(col, row int) (Matrix, error) {
	resm := m.Copy()
	eliminated := resm.Rows[row][col]
	if eliminated == 0 {
		return Matrix{}, errors.New("divide by zero")
	}

	resm.Rows[row][col] = 1
	resm.SetSwapped(col, row)
	for i := range resm.Rows {
		if i != row {
			resm.Rows[i][col] *= -1
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

	for i := range len(m.Rows) {
		for j := range len(m.Rows[i]) {
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
		Rows:      make([][]float64, len(m.Rows)),
		LeftTitle: m.LeftTitle,
		TopTitle:  m.TopTitle,
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

func (m *Matrix) FirstNegativeInRow(row int) (col int, err error) {
	for i, el := range m.Rows[row] {
		if el < 0 {
			return i, nil
		}
	}

	return 0, fmt.Errorf("no negative numbers found in the row %v", row)
}

func (m *Matrix) FindMinPositiveFor(col int) (row int, err error) {
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

func (m *Matrix) SetSwapped(col, row int) {
	if m.LeftTitle == nil || m.TopTitle == nil {
		m.fillTitleMaps()
	}

	m.TopTitle[col], m.LeftTitle[row] = m.LeftTitle[row], m.TopTitle[col]
}

func (m *Matrix) fillTitleMaps() {
	m.LeftTitle = make(map[int]Variable)
	m.TopTitle = make(map[int]Variable)
	for i := range m.Rows {
		m.LeftTitle[i] = Variable{"y", i}
	}
	for i := range m.Rows[0] {
		m.TopTitle[i] = Variable{"x", i}
	}
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
}

func RoundTo(num float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(num*ratio) / ratio
}
