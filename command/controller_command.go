package command

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"bitbucket.org/pkg/inflect"
	"github.com/mattn/echo-scaffold/template"
)

type ControllerCommand struct {
	PackageName        string
	ControllerName     string
	ModelName          string
	ModelNamePlural    string
	InstanceName       string
	InstanceNamePlural string
	RoutePath          string
	TemplateName       string
	Fields             map[string]string
}

func (command *ControllerCommand) Name() string {
	return "controller"
}

func (command *ControllerCommand) Help() {
	fmt.Printf(`Usage:
	echo-scaffold controller <controller name>

Description:
	The echo-scaffold controller command creates a new controller.

Example:
	echo-scaffold controller Posts
`)
}

func (command *ControllerCommand) Execute(args []string) {
	flag := flag.NewFlagSet(command.Name(), flag.ExitOnError)
	flag.Usage = command.Help
	flag.Parse(args)
	if flag.NArg() == 0 {
		command.Help()
		os.Exit(2)
	}

	command.ControllerName = inflect.Titleize(flag.Arg(0))
	command.RoutePath = inflect.Underscore(command.ControllerName)
	command.ModelName = inflect.Singularize(command.ControllerName)
	command.ModelNamePlural = inflect.Pluralize(command.ModelName)

	command.InstanceName = inflect.CamelizeDownFirst(command.ModelName)
	command.InstanceNamePlural = inflect.Pluralize(command.InstanceName)
	command.PackageName = template.PackageName()

	outputPath := filepath.Join("controllers", inflect.Underscore(command.ControllerName)+".go")
	builder := template.NewBuilder("controller.go.tmpl")
	builder.WriteToPath(outputPath, command)

	outputPath = filepath.Join("controllers", inflect.Underscore(command.ControllerName)+"_helpers.go")
	builder = template.NewBuilder("controller_helpers.go.tmpl")
	builder.WriteToPath(outputPath, command)

	outputPath = filepath.Join("controllers", "suite_test.go")
	builder = template.NewBuilder("suite_test.go.tmpl")
	builder.WriteToPath(outputPath, command)

	command.insertIntoRoutes()
}

func (command *ControllerCommand) insertIntoRoutes() {
	builder := template.NewBuilder("controller_router.go.tmpl")
	builder.InsertAfterToPath("controllers/router.go", "func Setup(", command)
}
