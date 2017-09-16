package command

import (
	"flag"
	"fmt"
)

type HelpCommand struct {
}

func (command *HelpCommand) Name() string {
	return "help"
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
	flag := flag.NewFlagSet(command.Name(), flag.ExitOnError)
	flag.Parse(args)
	if flag.NArg() == 0 {
		command.Help()
		return
	}

	targetCommand := FindCommand(flag.Arg(0))

	if targetCommand == nil {
		command.Help()
	} else {
		targetCommand.Help()
	}
}
