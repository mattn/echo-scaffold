package command

import (
	"fmt"
)

type HelpCommand struct {
}

func (command *HelpCommand) Help() {
	fmt.Printf(`Usage:
	echo-scaffold help <command name>

Description:
	The echo-scaffold help command prints help about the given command.

Example:
	echo-scaffold help init
`)
}

func (command *HelpCommand) Execute(args []string) {
	if len(args) == 0 {
		command.Help()
		return
	}

	targetCommand := FindCommand(args[0])

	if targetCommand == nil {
		command.Help()
	} else {
		targetCommand.Help()
	}
}
