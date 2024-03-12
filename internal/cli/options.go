package cli

import (
	"errors"
	"fmt"
	"strconv"
)

type Option = string

const (
	ExitOption                         = "exit"
	InverseMatrixOption                = "inverse"
	GetRankOption                      = "rank"
	SolveLinearEquationOption          = "solve_equtaion"
	SolveLinearInequationOption        = "solve_inequation"
	SolveIntegerLinearInequationOption = "solve_integer"
	HelpOption                         = "help"
	ClearOption                        = "clear"
)

func ReadInt() (int, error) {
	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		return 0, err
	}

	option, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return option, err
}

func ReadPositiveInt() (int, error) {
	n, err := ReadInt()
	if err != nil {
		return 0, err
	}

	if n <= 0 {
		return 0, errors.New("number should be greater than 0")
	}

	return n, nil
}

func ReadWord() (string, error) {
	var str string
	if _, err := fmt.Scanln(&str); err != nil {
		return "", errors.New("invalid input")
	}
	return str, nil
}

func ChooseOption() (Option, error) {
	option, err := ReadWord()
	if err != nil {
		return "", errors.New("invalid input")
	}

	if option != InverseMatrixOption &&
		option != GetRankOption &&
		option != SolveLinearEquationOption &&
		option != HelpOption &&
		option != ExitOption &&
		option != SolveLinearInequationOption &&
		option != SolveIntegerLinearInequationOption &&
		option != ClearOption {
		return "", errors.New("invalid options is chosen")
	}

	return option, nil
}
