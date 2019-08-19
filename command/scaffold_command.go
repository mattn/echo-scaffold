package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/markbates/inflect"
)

type ScaffoldCommand struct {
}

func (command *ScaffoldCommand) Name() string {
	return "scaffold"
}

func (command *ScaffoldCommand) Help() {
	fmt.Printf(`Usage:
	echo-scaffold scaffold <controller name> <field name>:<field type> ...

Description:
	The echo-scaffold scaffold command creates a new controller and model with the given fields.

Example:
	echo-scaffold controller Post Title:string Body:string
`)
}

func (command *ScaffoldCommand) Execute(args []string) {
	flag := flag.NewFlagSet(command.Name(), flag.ExitOnError)
	flag.Usage = command.Help
	flag.Parse(args)
	if flag.NArg() == 0 {
		command.Help()
		os.Exit(2)
	}

	flag.Args()[0] = inflect.Singularize(flag.Arg(0))
	modelCommand := &ModelCommand{}
	modelCommand.Execute(args)

	controllerCommand := &ControllerCommand{}
	controllerCommand.Execute([]string{modelCommand.ModelNamePlural})
}
