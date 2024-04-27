package main

import (
	"github.com/hrvadl/algo/internal/cli"
)

func main() {
	cli.PrintStudentInfo()
	go cli.Start()
	cli.HandleGracefulShutdown()
	cli.PrintExitMessage()
}
