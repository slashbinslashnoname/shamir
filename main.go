package main

import (
	"os"
	"syscall"

	"github.com/slashbinslashnoname/shamir/cmd"
	"github.com/slashbinslashnoname/shamir/utils"
	"golang.org/x/term"
)

func main() {
	inputSource := os.Stdin
	outputDestination := os.Stdout
	errorDestination := os.Stderr
	isTerminal := term.IsTerminal(int(syscall.Stdin))

	rootCommand := cmd.GenerateRootCommand(
		inputSource,
		outputDestination,
		errorDestination,
		isTerminal,
	)

	err := rootCommand.Execute()
	if err != nil {
		utils.ExitWithError(errorDestination, err)
	}
}
