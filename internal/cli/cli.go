package cli

import (
	"errors"
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hrvadl/algo/internal/equations"
	"github.com/hrvadl/algo/internal/games"
	"github.com/hrvadl/algo/internal/inequations"
	"github.com/hrvadl/algo/internal/matrix"
	"github.com/hrvadl/algo/pkg/tm"
)

const GracefulShutdownTime = 1 * time.Second

const (
	CalculateInequationFlag = 1 << iota
	CalculateIntegerFlag
	CalculateDoubledFlag
)

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
			HandleSolveLinearInequation(CalculateInequationFlag)
		case SolveIntegerLinearInequationOption:
			HandleSolveLinearInequation(CalculateInequationFlag | CalculateIntegerFlag)
		case SolveDoubledLinearInequationOption:
			HandleSolveLinearInequation(CalculateInequationFlag | CalculateDoubledFlag)
		case GetGameStrategies:
			HandleSolveGame()
		case SolveGameWithNature:
			HandleGameWithNature()
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

func HandleSolveGame() {
	m, err := HandleGetMatrix()
	if err != nil {
		PrintError(err)
		return
	}

	startingPoint := m
	fmt.Printf("\nJust confirmation. Your matrix: \n\n")
	m.Print()

	clean, err := m.GetCleanStrategySolution()
	if err == nil {
		fmt.Printf(
			"\n\nFound clean solution: (%d,%d) with game weight: %v\n\n",
			clean.Row,
			clean.Col,
			clean.Val,
		)
		return
	}

	PrintError(err)

	minabs := math.Abs(m.Min())
	m = *m.Add(minabs)
	compat, err := games.CompleteMatrixToCompatible(m)
	if err != nil {
		PrintError(err)
		return
	}

	m.FillLeftTitle()
	m.FillTopTitle()
	m = *compat
	m.InitialCols = len(m.Rows[0])
	m.InitialRows = len(m.Rows) - 1
	fmt.Printf("\nJust confirmation. Your matrix after correcting: \n\n")
	m.Print()

	support, err := inequations.FindSupportSolution(m)
	if err != nil {
		PrintError(err)
		return
	}

	optimal, err := inequations.FindMaxDoubledWithOptimalSolution(support.Matrix)
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nYour support solution: \n%v\n", matrix.RoundRowTo(support.Result, 2))
	fmt.Printf("\nYour optimal solution: \n%v\n", matrix.RoundRowTo(optimal.MaxSolution.Result, 2))
	fmt.Printf(
		"\nYour doubled optimal solution: \n%v\n",
		matrix.RoundRowTo(optimal.MinSolution.Result, 2),
	)
	fmt.Printf("\nYour max: \n%v\n", matrix.RoundTo(optimal.Max, 2))
	fmt.Printf("\nYour min (doubled): \n%v\n", matrix.RoundTo(optimal.Min, 2))

	gameWeight := games.GetGameWeight(optimal.MaxSolution.Matrix)
	correctedGameWeight := games.CorrectGameWeight(gameWeight, minabs)
	firstPlayerStrategy := games.CorrectMixedStrategy(optimal.MaxSolution.Result, gameWeight)
	secondPlayerStrategy := games.CorrectMixedStrategy(optimal.MinSolution.Result, gameWeight)

	fmt.Printf("\n\nFirst player strategy: %v", firstPlayerStrategy)
	fmt.Printf("\nSecond player strategy: %v", secondPlayerStrategy)
	fmt.Printf("\nGame Weight: %v\n\n", correctedGameWeight)

	fmt.Println("How many times do you want to simulate game?")
	n, err := ReadPositiveInt()
	if err != nil {
		PrintError(err)
		return
	}

	steps := games.SimulateGame(games.SimulationOptions{
		Times:                n,
		FirstPlayerStrategy:  firstPlayerStrategy,
		SecondPlayerStrategy: secondPlayerStrategy,
		Matrix:               startingPoint,
	})

	for i, s := range steps {
		fmt.Printf("idx: %d %+v\n\n", i, s)
	}
}

func HandleGameWithNature() {
	m, err := HandleGetMatrix()
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nJust confirmation. Your matrix: \n\n")
	m.Print()

	fmt.Printf("\nType y coefficient for pessimist optimist strategy:\n")
	y, err := ReadFloat()
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nType v coefficient for laplas strategy:\n")
	v, err := ReadFloat()
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nType the p vector for bayes straregy:\n")
	p, err := GetMatrix(1, len(m.Rows[0]))
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf(
		"\n\nSolution with maxmin (Wald) strategy: %v",
		games.ToHumanReadable(games.SolveWithMaxMinStrategy(m)),
	)
	fmt.Printf(
		"\n\nSolution with maxmax strategy: %v",
		games.ToHumanReadable(games.SolveWithMaxMaxStrategy(m)),
	)
	fmt.Printf(
		"\n\nSolution with pessimist optimist (Gurwic) strategy: %v",
		games.ToHumanReadable(games.SolveWithPessimismOptimismStrategy(m, y)),
	)
	fmt.Printf(
		"\n\nSolution with minmax (Sevige) strategy: %v",
		games.ToHumanReadable(games.SolveWithMinMaxStrategy(m)),
	)

	bsol, err := games.SolveWithBayesPrinicple(m, p.Rows[0])
	if err != nil {
		PrintError(fmt.Errorf("cannot solve with bayes principle: %w", err))
		return
	}

	fmt.Printf("\n\nSolution with bayes strategy: %v", games.ToHumanReadable(bsol))

	lsol, err := games.SolveWithLaplasPrinciple(m, v)
	if err != nil {
		PrintError(fmt.Errorf("cannot solve with laplas principle: %w", err))
		return
	}

	fmt.Printf("\n\nSolution with laplas strategy: %v\n\n", games.ToHumanReadable(lsol))
}

func HandleSolveLinearInequation(flag uint8) {
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
		LeftTitle: make([]matrix.Variable, n+1),
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
			m.LeftTitle[i] = matrix.Variable{
				FirstStageName:   "0",
				SecondStageName:  "u",
				SecondStageIndex: i,
			}
		} else {
			m.LeftTitle[i] = matrix.Variable{
				FirstStageName:   "y",
				FirstStageIndex:  i,
				SecondStageName:  "u",
				SecondStageIndex: i,
			}
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

	m.FillTopTitle()
	m.InitialCols = len(m.Rows[0])
	m.InitialRows = len(m.Rows) - 1

	switch minMax {
	case "max":
		if flag&CalculateIntegerFlag != 0 {
			HandleGetMaxWithIntegerOptimalSolution(m)
			return
		}

		if flag&CalculateDoubledFlag != 0 {
			HandleGetDoubledMaxWithOptimalSolution(m)
			return
		}

		HandleGetMaxWithOptimalSolution(m)
		return

	case "min":
		if flag&CalculateIntegerFlag != 0 {
			HandleGetMinWithIntegerOptimalSolution(m)
			return
		}

		if flag&CalculateDoubledFlag != 0 {
			HandleGetDoubledMinWithOptimalSolution(m)
			return
		}

		HandleGetMinWithOptimalSolution(m)
		return
	}
}

func HandleGetMinWithIntegerOptimalSolution(m matrix.Matrix) {
	m, err := m.DeleteZeros()
	if err != nil {
		PrintError(err)
		return
	}

	optimal, support, err := inequations.FindMinIntegerSolution(m)
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nYour support integer solution: \n%v\n", support.Result)
	fmt.Printf("\nYour optimal integer solution: \n%v\n", optimal.Result)
	fmt.Printf("\nYour min: \n%v\n", optimal.Min)
}

func HandleGetMaxWithIntegerOptimalSolution(m matrix.Matrix) {
	m, err := m.DeleteZeros()
	if err != nil {
		PrintError(err)
		return
	}

	optimal, support, err := inequations.FindMaxIntegerSolution(m)
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nYour support integer solution: \n%v\n", support.Result)
	fmt.Printf("\nYour optimal integer solution: \n%v\n", optimal.Result)
	fmt.Printf("\nYour max: \n%v\n", optimal.Max)
}

func HandleGetDoubledMinWithOptimalSolution(m matrix.Matrix) {
	m, err := m.DeleteZeros()
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nFinding the support solution...\n")
	support, err := inequations.FindMinWithSupportSolution(m)
	if err != nil {
		PrintError(err)
		return
	}

	optimal, err := inequations.FindMinDoubledWithOptimalSolution(support.Matrix)
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nYour support solution: \n%v\n", support.Result)

	fmt.Printf("\nYour optimal solution (min): \n%v\n", optimal.MinSolution.Result)
	fmt.Printf("\nYour optimal solution (max): \n%v\n", optimal.MaxSolution.Result)

	fmt.Printf("\nYour min: \n%v\n", optimal.Min)
	fmt.Printf("\nYour max (doubled): \n%v\n", optimal.Max)
}

func HandleGetDoubledMaxWithOptimalSolution(m matrix.Matrix) {
	m, err := m.DeleteZeros()
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nFinding the support solution...\n")
	support, err := inequations.FindSupportSolution(m)
	if err != nil {
		PrintError(err)
		return
	}

	optimal, err := inequations.FindMaxDoubledWithOptimalSolution(support.Matrix)
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nYour support solution: \n%v\n", support.Result)
	fmt.Printf("\nYour optimal solution: \n%v\n", optimal.MaxSolution.Result)
	fmt.Printf("\nYour doubled optimal solution: \n%v\n", optimal.MinSolution.Result)
	fmt.Printf("\nYour max: \n%v\n", optimal.Max)
	fmt.Printf("\nYour min (doubled): \n%v\n", optimal.Min)
}

func HandleGetMinWithOptimalSolution(m matrix.Matrix) {
	m, err := m.DeleteZeros()
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nFinding the support solution...\n")
	support, err := inequations.FindMinWithSupportSolution(m)
	if err != nil {
		PrintError(err)
		return
	}

	optimal, err := inequations.FindMinWithOptimalSolution(support.Matrix)
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nYour support solution: \n%v\n", support.Result)
	fmt.Printf("\nYour optimal solution: \n%v\n", optimal.Solution.Result)
	fmt.Printf("\nYour min: \n%v\n", optimal.Min)
}

func HandleGetMaxWithOptimalSolution(m matrix.Matrix) {
	m, err := m.DeleteZeros()
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nFinding the support solution...\n")
	support, err := inequations.FindSupportSolution(m)
	if err != nil {
		PrintError(err)
		return
	}

	optimal, err := inequations.FindMaxWithOptimalSolution(support.Matrix)
	if err != nil {
		PrintError(err)
		return
	}

	fmt.Printf("\nYour support solution: \n%v\n", support.Result)
	fmt.Printf("\nYour optimal solution: \n%v\n", optimal.Solution.Result)
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
