package main

import (
	"github.com/hrvadl/algo/internal/cli"
)

func main() {
	cli.PrintStudentInfo()
	cli.PrintHelp()
	cli.Start()
	cli.PrintExitMessage()
}
