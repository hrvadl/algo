package cli

import (
	"errors"
	"fmt"

	"github.com/hrvadl/algo/internal/equations"
	"github.com/hrvadl/algo/internal/matrix"
)

func Start() {
loop:
	for {
		option, err := ChooseOption()
		if err != nil {
			PrintError(err)
			continue
		}

		switch option {
		case InverseMatrixOption:
			HandleInverseMatrix()
		case GetRankOption:
			HandleGetRank()
		case HelpOption:
			PrintHelp()
		case SolveLinearEquationOption:
			HandleSolveLinearEquation()
		case ExitOption:
			break loop
		}
	}
}

func HandleInverseMatrix() {
	fmt.Println("\nType the size of your matrix: ")
	size, err := ReadInt()
	if err != nil {
		PrintError(err)
		return
	}

	m, err := GetMatrix(size, size)
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Println("\nJust confirmation. Your matrix: ")
	m.Print()

	m, err = m.Invert()
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Println("\nThe result is:")
	m = m.Round()
	m.Print()
}

func HandleGetRank() {
	m, err := HandleGetMatrix()
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Println("\nJust confirmation. Your matrix: ")
	m.Print()

	rank := m.Rank()
	fmt.Printf("\nThe rank of your matrix is: %v\n", rank)
}

func HandleSolveLinearEquation() {
	fmt.Println("\nInput your A matrix: ")
	a, err := HandleGetMatrix()
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Println("\nInput your B matrix: ")
	b, err := HandleGetMatrix()
	if err != nil {
		PrintError(err)
		return
	}

	_, ha := a.GetDimensions()
	_, hb := b.GetDimensions()
	if ha != hb {
		fmt.Printf("ha: %v, hb: %v", ha, hb)
		PrintError(errors.New("matrixes are not compatible"))
		return
	}

	res := equations.SolveSystem(equations.EquationSystem{
		A: a,
		B: b,
	})

	fmt.Println("\nThe result of solving equation: ")
	fmt.Printf("%v", res)
}

func HandleGetMatrix() (matrix.Matrix, error) {
	fmt.Println("\nType the amount of rows of your matrix: ")
	rows, err := ReadInt()
	if err != nil {
		return matrix.Matrix{}, err
	}

	fmt.Println("\nType the amount of columns of your matrix: ")
	col, err := ReadInt()
	if err != nil {
		return matrix.Matrix{}, err
	}

	return GetMatrix(rows, col)
}
