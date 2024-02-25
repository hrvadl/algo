package cli

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	ExitOption = iota
	InverseMatrixOption
	GetRankOption
	SolveLinearEquationOption
	HelpOption
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

func ChooseOption() (int, error) {
	option, err := ReadInt()
	if err != nil {
		return 0, err
	}

	if option != InverseMatrixOption &&
		option != GetRankOption &&
		option != SolveLinearEquationOption &&
		option != HelpOption && option != ExitOption {
		return 0, errors.New("invalid options is chosen")
	}

	return option, nil
}
