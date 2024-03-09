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

func (v Variable) IsZ() bool {
	return v.Name == "z"
}

func (v Variable) IsZero() bool {
	return v.Name == "0"
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

func (m *Matrix) DeleteZeros() (Matrix, error) {
	toDelete := -1
	for row, variable := range m.LeftTitle {
		if variable.IsZero() {
			toDelete = row
			break
		}
	}

	if toDelete == -1 {
		return *m, nil
	}

	newM, err := m.DeleteRow(toDelete)
	if err != nil {
		return Matrix{}, err
	}

	return newM.DeleteZeros()
}

func (m *Matrix) DeleteRow(row int) (Matrix, error) {
	col, err := m.FirstPositiveInRowExceptLastColumn(row)
	if err != nil {
		return Matrix{}, fmt.Errorf("can't delete. no positive elements in row %v", row)
	}

	toDeleteRow, err := m.FindMinPositiveFor(col)
	if err != nil {
		return Matrix{}, fmt.Errorf("can't delete. %w", err)
	}

	newM, err := m.JordanEliminateModified(col, toDeleteRow)
	if err != nil {
		return Matrix{}, err
	}

	if !m.TopTitle[col].IsZero() {
		return newM, nil
	}

	for i, row := range newM.Rows {
		for j := range row {
			if j <= col {
				continue
			}

			newM.Rows[i][j-1] = newM.Rows[i][j]
		}

		newM.Rows[i] = newM.Rows[i][:len(row)-1]
	}

	delete(newM.TopTitle, col)
	tmp := make([]Variable, len(newM.TopTitle))
	for i, variable := range newM.TopTitle {
		if i < col {
			tmp[i] = variable
		} else {
			tmp[i-1] = variable
		}
	}

	clear(newM.TopTitle)
	for i, variable := range tmp {
		newM.TopTitle[i] = variable
	}

	return newM, nil
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

func (m *Matrix) FirstNegativeInRowExceptLastColumn(row int) (col int, err error) {
	for i := 0; i < len(m.Rows[row])-1; i++ {
		if m.Rows[row][i] < 0 {
			return i, nil
		}
	}

	return 0, fmt.Errorf("no negative numbers found in the row %v", row)
}

func (m *Matrix) FirstPositiveInRowExceptLastColumn(row int) (int, error) {
	for i := 0; i < len(m.Rows[row])-1; i++ {
		if m.Rows[row][i] > 0 {
			return i, nil
		}
	}

	return 0, fmt.Errorf("no negative numbers found in the row %v", row)
}

func (m *Matrix) FirstNegativeRowInLastColumn() (int, error) {
	lastCol := len(m.Rows[0]) - 1
	for i := 0; i < len(m.Rows)-1; i++ {
		if m.Rows[i][lastCol] < 0 {
			return i, nil
		}
	}

	return 0, errors.New("no negatives found")
}

func (m *Matrix) FirstNegativeColumnInLastRow() (int, error) {
	lastRow := len(m.Rows) - 1
	fmt.Printf("\nFinding the optimal solution...\n")
	for i := 0; i < len(m.Rows[lastRow])-1; i++ {
		if m.Rows[lastRow][i] < 0 {
			return i, nil
		}
	}

	return 0, errors.New("no negatives found")
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
	if len(m.LeftTitle) == 0 {
		m.FillLeftTitle()
	}

	if len(m.TopTitle) == 0 {
		m.FillTopTitle()
	}

	m.TopTitle[col], m.LeftTitle[row] = m.LeftTitle[row], m.TopTitle[col]
}

func (m *Matrix) FillLeftTitle() {
	m.LeftTitle = make(map[int]Variable)
	for i := range m.Rows {
		if i == len(m.Rows)-1 {
			m.LeftTitle[i] = Variable{"1", 0}
		} else {
			m.LeftTitle[i] = Variable{"y", i}
		}
	}
}

func (m *Matrix) FillTopTitle() {
	m.TopTitle = make(map[int]Variable)
	for i := range m.Rows[0] {
		if i == len(m.Rows[0])-1 {
			m.TopTitle[i] = Variable{"z", i}
		} else {
			m.TopTitle[i] = Variable{"x", i}
		}
	}
}

func (m *Matrix) GetXCount() int {
	var total int
	for _, variable := range m.LeftTitle {
		if variable.IsX() {
			total++
		}
	}

	for _, variable := range m.TopTitle {
		if variable.IsX() {
			total++
		}
	}

	return total
}

func (m *Matrix) IntegerLimitationFor(row int) Row {
	res := make(Row, 0, len(m.Rows[row]))

	for _, el := range m.Rows[row] {
		integer, fraction := math.Modf(el)
		if fraction == 0 {
			res = append(res, el)
			continue
		}

		if integer < 0 || fraction < 0 {
			integer--
		}

		res = append(res, RoundTo(el-integer, 2))
	}

	return res
}

func (m *Matrix) NegativeRowFor(row Row) Row {
	for i := range row {
		row[i] /= -1
	}
	return row
}

func (m *Matrix) InsertRow(row Row) Matrix {
	newM := m.Copy()
	lastCol := len(newM.Rows) - 1
	old := newM.Rows[lastCol]
	newM.Rows = append(newM.Rows, old)
	newM.Rows[lastCol] = row

	for row, variable := range newM.LeftTitle {
		if variable.IsZ() {
			newM.LeftTitle[row+1] = variable
			newM.LeftTitle[row] = Variable{Name: "s"}
			break
		}
	}

	return newM
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
