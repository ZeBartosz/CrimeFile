package app

import (
	"errors"
	"fmt"
)

var (
	ErrMustProvideArgs = errors.New("Please provide a commamnd. Type -h for help")
)

func App(args []string) error {
	if len(args) <= 0 {
		return ErrMustProvideArgs
	}

	commands := map[string]func([]string) error{
		"generate": generateScen,
	}

	commandName := args[0]

	command, exists := commands[commandName]
	if !exists {
		return fmt.Errorf("unknown command: %s", commandName)
	}

	commandArgs := args[1:]
	return command(commandArgs)

}

func generateScen(args []string) error {
	if len(args) < 1 {
		return errors.New("generate requires a scenario name")
	}

	scenario := args[0]
	fmt.Printf("Generating scenario: %s\n", scenario)
	return nil
}
