package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattn/echo-scaffold/command"
)

func printUsageAndExit() {
	fmt.Printf("Echo Scaffold\n")
	for _, command := range command.Commands {
		fmt.Printf("\nCommand `%s`:\n\n", command.Name)
		command.Help()
	}
	os.Exit(0)
}

func checkGOPATH() {
	if os.Getenv("GOPATH") == "" {
		fmt.Printf("$GOPATH is not defined. Exiting.\n")
		os.Exit(2)
	}

	wd, _ := os.Getwd()
	wd = filepath.ToSlash(wd)
	found := false
	for _, p := range filepath.SplitList(os.Getenv("GOPATH")) {
		if strings.HasPrefix(strings.ToLower(wd), strings.ToLower(filepath.ToSlash(p))) {
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("%s is not in the $GOPATH. Exiting.\n", wd)
		os.Exit(3)
	}
}

func main() {
	checkGOPATH()

	if len(os.Args) < 2 {
		printUsageAndExit()
	}

	commandName := os.Args[1]
	commandArgs := []string{}

	if len(os.Args) > 2 {
		commandArgs = os.Args[2:]
	}

	command := command.FindCommand(commandName)
	if command == nil {
		printUsageAndExit()
	}

	command.Execute(commandArgs)
}
