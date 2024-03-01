package parse

import (
	"errors"
	"strings"

	"github.com/hrvadl/algo/internal/matrix"
)

func InequationFromString(str string) (matrix.Row, error) {
	length := len(strings.FieldsFunc(str, func(r rune) bool {
		return r == '-' || r == '+'
	}))

	if length == 0 {
		return nil, errors.New("invalid expression")
	}

	split := strings.FieldsFunc(str, func(r rune) bool {
		return r == '>' || r == '<'
	})

	if len(split) != 2 {
		return nil, errors.New("invalid expression")
	}

	r, err := NewEvaluator(strings.ReplaceAll(split[1], "=", "")).EvaluateFromString()
	if err != nil {
		return nil, err
	}

	l, err := NewEvaluator(split[0]).EvaluateFromString()
	if err != nil {
		return nil, err
	}

	lsign, rsign := getSignsForInequality(str)

	for i := range r {
		r[i] *= rsign
	}

	for i := range l {
		l[i] *= lsign
	}

	return append(l, r...), nil
}

func getSignsForInequality(str string) (left float64, right float64) {
	var (
		lsign = 1.
		rsign = 1.
	)

	if strings.Contains(str, ">") {
		rsign = -1
	} else {
		lsign = -1
	}

	return lsign, rsign
}
