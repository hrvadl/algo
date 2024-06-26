package cli

import (
	"fmt"
)

func PrintStudentInfo() {
	fmt.Println("Lab 1 Algorithms Vadym Hrashchenko 634p")
}

func PrintHelp() {
	fmt.Println()
	fmt.Printf("Exit the program:                    %s\n", ExitOption)
	fmt.Printf("Calculate inverse matrix:            %s\n", InverseMatrixOption)
	fmt.Printf("Calculate rank of the matrix:        %s\n", GetRankOption)
	fmt.Printf("Calculate linear equation system:    %s\n", SolveLinearEquationOption)
	fmt.Printf("Calculate linear inequation:         %s\n", SolveLinearInequationOption)
	fmt.Printf("Calculate integer linear inequation: %s\n", SolveIntegerLinearInequationOption)
	fmt.Printf("Calculate doubled linear inequation: %s\n", SolveDoubledLinearInequationOption)
	fmt.Printf("Get game strategies:                 %s\n", GetGameStrategies)
	fmt.Printf("Solve game with nature:              %s\n", SolveGameWithNature)
	fmt.Printf("Print this message:                  %s\n", HelpOption)
	fmt.Printf("Clear the screen:                    %s\n", ClearOption)
	fmt.Println()
}

func PrintError(err error) {
	fmt.Printf("\nError occurred: %v\n", err)
}

func PrintExitMessage() {
	fmt.Printf("\nThe program is terminated... Bye!\n")
}
