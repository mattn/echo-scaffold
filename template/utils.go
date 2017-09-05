package template

import (
	"fmt"

	"github.com/mattn/go-colorable"
	"github.com/mgutz/ansi"
)

var out = colorable.NewColorableStdout()

func printAction(color string, action string, target string) {
	fmt.Fprintf(out, "	%s  %s\n", ansi.Color(action, color), target)
}
