package command

var (
	Commands = map[string]Base{
		"help":       &HelpCommand{},
		"init":       &InitCommand{},
		"model":      &ModelCommand{},
		"controller": &ControllerCommand{},
		"scaffold":   &ScaffoldCommand{},
	}
)

func FindCommand(name string) Base {
	switch name {
	case "i":
		{
			name = "init"
		}
	case "m":
		{
			name = "model"
		}
	case "c":
		{
			name = "controller"
		}
	case "s":
		{
			name = "scaffold"
		}
	case "h":
		{
			name = "help"
		}
	}

	return Commands[name]
}
