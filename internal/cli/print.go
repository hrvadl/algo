package cli

import (
	"fmt"
)

func PrintStudentInfo() {
	fmt.Println("Lab 1 Algorithms Vadym Hrashchenko 634p")
}

func PrintHelp() {
	fmt.Println()
	fmt.Println("Exit the program: 0")
	fmt.Println("Calculate inverse matrix: 1")
	fmt.Println("Calculate rank of the matrix: 2")
	fmt.Println("Calculate linear equation system: 3")
	fmt.Println("Print this message: 4")
	fmt.Println()
}

func PrintError(err error) {
	fmt.Printf("Error occurred: %v", err)
}

func PrintExitMessage() {
	fmt.Printf("\nThe program is terminated... Bye!")
}
