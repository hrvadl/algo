package cli

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hrvadl/algo/internal/equations"
	"github.com/hrvadl/algo/internal/inequations"
	"github.com/hrvadl/algo/internal/matrix"
	"github.com/hrvadl/algo/pkg/tm"
)

const GracefulShutdownTime = 1 * time.Second

func Start() {
	for {
		PrintHelp()
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
		case ClearOption:
			tm.Clear()
		case SolveLinearEquationOption:
			HandleSolveLinearEquation()
		case SolveLinearInequationOption:
			HandleSolveLinearInequation()
		case ExitOption:
			PrintExitMessage()
			os.Exit(0)
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
	_, rows := a.GetDimensions()
	b, err := GetMatrix(rows, 1)
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Println()

	res := equations.SolveSystem(equations.EquationSystem{
		A: a,
		B: b,
	})

	fmt.Println("\nThe result of solving equation: ")
	fmt.Printf("%v\n", res)
}

func HandleSolveLinearInequation() {
	fmt.Printf("\nInput Z: \n")
	z, err := GetNegativeFunctionRow()
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nInput amount of inequation limitations: \n")
	n, err := ReadPositiveInt()
	if err != nil {
		PrintError(err)
		return
	}

	m := matrix.Matrix{
		Rows:      make([]matrix.Row, 0, n+1),
		LeftTitle: make(map[int]matrix.Variable),
		TopTitle:  make(map[int]matrix.Variable),
	}

	for i := range n {
		fmt.Printf("\nInput the inequation: \n")
		row, isEquation, err := GetNegativeRowFromExpression()
		if err != nil {
			PrintError(err)
			return
		}

		if len(row) != len(z) {
			PrintError(errors.New("rows should have the same size"))
			return
		}

		if isEquation {
			m.LeftTitle[i] = matrix.Variable{Name: "0"}
		} else {
			m.LeftTitle[i] = matrix.Variable{Name: "y", Index: i}
		}

		m.Rows = append(m.Rows, row)
	}

	m.Rows = append(m.Rows, z)
	fmt.Printf("\nJust confirmation. Your matrix: \n\n")
	m.Print()

	fmt.Printf("\nDo you want to find min or max?\n")
	minMax, err := ReadWord()
	if err != nil {
		PrintError(err)
		return
	}

	if minMax == "max" {
		HandleGetMaxWithOptimalSolution(m)
		return
	}

	HandleGetMinWithOptimalSolution(m)
}

func HandleGetMinWithOptimalSolution(m matrix.Matrix) {
	m, err := m.DeleteZeros()
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nFinding the support solution...\n")
	support, solved, err := inequations.FindMaxWithSupportSolution(m)
	if err != nil {
		PrintError(err)
		return
	}
	m = *solved

	optimal, err := inequations.FindMinWithOptimalSolution(m)
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nYour support solution: \n%v\n", support)
	fmt.Printf("\nYour optimal solution: \n%v\n", optimal.Solution)
	fmt.Printf("\nYour min: \n%v\n", optimal.Min)
}

func HandleGetMaxWithOptimalSolution(m matrix.Matrix) {
	m, err := m.DeleteZeros()
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nFinding the support solution...\n")
	support, solved, err := inequations.FindMaxWithSupportSolution(m)
	if err != nil {
		PrintError(err)
		return
	}

	m = *solved

	optimal, err := inequations.FindMaxWithOptimalSolution(m)
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nYour support solution: \n%v\n", support)
	fmt.Printf("\nYour optimal solution: \n%v\n", optimal.Solution)
	fmt.Printf("\nYour max: \n%v\n", optimal.Max)
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

func HandleGracefulShutdown() {
	terminated := make(chan os.Signal, 1)
	signal.Notify(terminated, syscall.SIGINT, syscall.SIGTERM)
	reason := <-terminated
	fmt.Printf("\nReceived: %v. Terminating...\n", reason.String())
	time.Sleep(GracefulShutdownTime)
}
