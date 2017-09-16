package command

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"bitbucket.org/pkg/inflect"
	"github.com/mattn/echo-scaffold/template"
)

// ModelCommand generates files related to model.
type ModelCommand struct {
	PackageName        string
	ModelName          string
	ModelNamePlural    string
	InstanceName       string
	InstanceNamePlural string
	TemplateName       string
	Fields             map[string]string
}

func (command *ModelCommand) Name() string {
	return "model"
}

// Help prints a help message for this command.
func (command *ModelCommand) Help() {
	fmt.Printf(`Usage:
	echo-scaffold model <model name> <field name>:<field type> ...

Description:
	The echo-scaffold model command creates a new model with the given fields.

Example:
	echo-scaffold model Post Title:string Body:string 
`)
}

func findFieldType(name string) string {
	switch name {
	case "text":
		name = "string"
	case "float":
		name = "float64"
	case "boolean":
		name = "bool"
	case "integer":
		name = "int"
	case "time", "datetime":
		name = "int64"
	}
	return name
}

// Converts "<fieldname>:<type>" to {"<fieldname>": "<type>"}
func processFields(args []string) map[string]string {
	fields := map[string]string{}
	for _, arg := range args {
		fieldNameAndType := strings.SplitN(arg, ":", 2)
		fields[inflect.Titleize(fieldNameAndType[0])] = findFieldType(fieldNameAndType[1])
	}

	return fields
}

// Execute runs this command.
func (command *ModelCommand) Execute(args []string) {
	flag := flag.NewFlagSet(command.Name(), flag.ExitOnError)
	flag.Usage = command.Help
	flag.Parse(args)
	if flag.NArg() < 2 {
		command.Help()
		os.Exit(2)
	}

	command.ModelName = inflect.Titleize(flag.Arg(0))
	command.ModelNamePlural = inflect.Pluralize(command.ModelName)

	command.Fields = processFields(flag.Args()[1:])
	command.InstanceName = inflect.CamelizeDownFirst(command.ModelName)
	command.InstanceNamePlural = inflect.Pluralize(command.InstanceName)
	command.PackageName = template.PackageName()

	outputPath := filepath.Join("models", inflect.Underscore(command.ModelName)+".go")

	builder := template.NewBuilder("model.go.tmpl")
	builder.WriteToPath(outputPath, command)

	outputPath = filepath.Join("models", inflect.Underscore(command.ModelName)+"_dbsession.go")
	builder = template.NewBuilder("model_dbsession.go.tmpl")
	builder.WriteToPath(outputPath, command)
}
