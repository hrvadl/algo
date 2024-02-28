package cli

import (
	"errors"
	"fmt"
	"strconv"
)

type Option = string

const (
	ExitOption                = "exit"
	InverseMatrixOption       = "inverse"
	GetRankOption             = "rank"
	SolveLinearEquationOption = "solve_linear"
	HelpOption                = "help"
	ClearOption               = "clear"
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

func ChooseOption() (Option, error) {
	var option string
	if _, err := fmt.Scanln(&option); err != nil {
		return "", errors.New("invalid input")
	}

	if option != InverseMatrixOption &&
		option != GetRankOption &&
		option != SolveLinearEquationOption &&
		option != HelpOption &&
		option != ExitOption &&
		option != ClearOption {
		return "", errors.New("invalid options is chosen")
	}

	return option, nil
}
