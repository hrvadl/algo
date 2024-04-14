package parse

import (
	"errors"
	"strconv"
	"unicode"
)

func NewEvaluator(str string) *ExpressionEvaluator {
	return &ExpressionEvaluator{
		String:     str,
		expression: []rune(str),
		result:     make([]float64, 0, 50),
	}
}

type ExpressionEvaluator struct {
	String     string
	expression []rune
	result     []float64
	i          int
}

func (i ExpressionEvaluator) EvaluateFromString() ([]float64, error) {
	for i.IsNotOver() {
		cursor := i.Cursor()
		if i.isVariableWithCoefficient(cursor) {
			coef, col, n, err := i.ParseCoefficientAndColumn(cursor)
			if err != nil {
				return nil, err
			}

			i.SetColumnWith(col, coef)
			i.MoveCursorTo(n + 1)
			continue
		}

		if i.isVariableWithSign(cursor) {
			sign, col, err := i.ParseSignAndColumn(cursor)
			if err != nil {
				return nil, err
			}

			i.SetColumnWith(col, sign)
			i.MoveCursorTo(3)
			continue
		}

		if i.isVariable(cursor) {
			col, err := i.ParseColumn(cursor + 1)
			if err != nil {
				return nil, err
			}

			i.SetColumnWith(col, 1)
			i.MoveCursorTo(2)
			continue
		}

		if isDigit, _ := i.isDigit(cursor); isDigit {
			dig, n, err := i.ParseDigit(cursor)
			if err != nil {
				return nil, err
			}
			i.MoveCursorTo(n)
			i.SetColumnWith(len(i.result)+1, dig)
		}
	}

	return i.result, nil
}

func (i *ExpressionEvaluator) IsNotOver() bool {
	return i.Cursor() < len(i.expression)
}

func (i *ExpressionEvaluator) Cursor() int {
	return i.i
}

func (i *ExpressionEvaluator) MoveCursorTo(n int) {
	i.i += n
}

func (i *ExpressionEvaluator) ParseCoefficientAndColumn(
	cursor int,
) (coef float64, col int, n int, err error) {
	coef, n, err = i.ParseDigit(cursor)
	if err != nil {
		return 0, 0, 0, err
	}

	col, err = i.ParseColumn(cursor + n + 1)
	if err != nil {
		return 0, 0, 0, err
	}

	return coef, col, n + 1, nil
}

func (i *ExpressionEvaluator) ParseSignAndColumn(cursor int) (float64, int, error) {
	sign, err := i.ParseSign(cursor)
	if err != nil {
		return 0, 0, err
	}

	col, err := i.ParseColumn(cursor + 2)
	if err != nil {
		return 0, 0, err
	}

	return sign, col, nil
}

func (i *ExpressionEvaluator) ParseColumn(cursor int) (int, error) {
	return strconv.Atoi(string(i.expression[cursor]))
}

func (i *ExpressionEvaluator) ParseSign(cursor int) (float64, error) {
	return strconv.ParseFloat(string(i.expression[cursor])+"1", 64)
}

func (i *ExpressionEvaluator) ParseDigit(cursor int) (float64, int, error) {
	if i.expression[cursor] != '+' && i.expression[cursor] != '-' &&
		!unicode.IsDigit(i.expression[cursor]) {
		return 0, 0, errors.New("invalid digit")
	}

	res := string(i.expression[cursor])
	for j := 1; cursor+j < len(i.expression); j++ {
		if !unicode.IsDigit(i.expression[cursor+j]) {
			break
		}
		res += string(i.expression[cursor+j])
	}

	parsed, err := strconv.ParseFloat(res, 64)
	if err != nil {
		return 0, len(res), err
	}

	return parsed, len(res), nil
}

func (i *ExpressionEvaluator) SetColumnWith(col int, num float64) {
	if col > len(i.result) {
		i.result = i.result[:col]
	}
	i.result[col-1] = num
}

func (i *ExpressionEvaluator) isVariableWithSign(idx int) bool {
	if len(i.expression) <= idx+1 {
		return false
	}
	return (i.expression[idx] == '-' || i.expression[idx] == '+') && i.isVariable(idx+1)
}

func (i *ExpressionEvaluator) isVariableWithCoefficient(idx int) bool {
	if len(i.expression) <= idx+2 {
		return false
	}

	isDigit, n := i.isDigit(idx)
	if !isDigit {
		return false
	}

	return i.isVariable(idx + n)
}

func (i *ExpressionEvaluator) isVariable(idx int) bool {
	if len(i.expression) <= idx+1 {
		return false
	}

	isDigit, _ := i.isDigit(idx + 1)
	if !isDigit {
		return false
	}

	return unicode.IsLetter(i.expression[idx])
}

func (i *ExpressionEvaluator) isDigit(cursor int) (bool, int) {
	isSign := i.expression[cursor] == '+' || i.expression[cursor] == '-'
	if !isSign && !unicode.IsDigit(i.expression[cursor]) {
		return false, 0
	}

	res := 1
	for j := 1; cursor+j < len(i.expression); j++ {
		if unicode.IsDigit(i.expression[cursor+j]) {
			res++
			continue
		}

		if j == 1 && isSign {
			return false, 0
		}
		break
	}

	return true, res
}
