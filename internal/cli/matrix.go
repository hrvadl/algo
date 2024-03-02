package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hrvadl/algo/internal/cli/parse"
	"github.com/hrvadl/algo/internal/matrix"
)

func GetMatrix(rows, columns int) (matrix.Matrix, error) {
	if rows < 0 || columns < 0 {
		return matrix.Matrix{}, errors.New("can't create a matrix with the zero size")
	}

	m := matrix.Matrix{
		Rows: make([]matrix.Row, rows),
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\nType the elements in the row  with the space between each:\n")
	for i := 0; i < rows && scanner.Scan(); i++ {
		tokens := strings.Fields(scanner.Text())

		if len(tokens) != columns {
			return matrix.Matrix{}, fmt.Errorf("invalid amount of tokens, expected: %v", columns)
		}

		for _, token := range tokens {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return matrix.Matrix{}, err
			}

			m.Rows[i] = append(m.Rows[i], num)
		}
	}

	return m, nil
}

func GetNegativeRowFromExpression() (matrix.Row, bool, error) {
	inequality, isEquation, err := GetRowFromExpression()
	if err != nil {
		return nil, false, err
	}

	for i := 0; i < len(inequality)-1; i++ {
		inequality[i] /= -1
	}

	return inequality, isEquation, nil
}

func GetRowFromExpression() (matrix.Row, bool, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return parse.EquationOrInequationFromString(scanner.Text())
}

func GetNegativeFunctionRow() (matrix.Row, error) {
	r, err := GetFunctionRow()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(r)-1; i++ {
		r[i] /= -1
	}

	return r, nil
}

func GetFunctionRow() (matrix.Row, error) {
	r, err := GetRowFromOneSide()
	if err != nil {
		return nil, err
	}

	return append(r, 0), nil
}

func GetRowFromOneSide() (matrix.Row, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return parse.NewEvaluator(scanner.Text()).EvaluateFromString()
}
