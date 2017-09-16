package command

type Cmd struct {
	Short string
	Name  string
	Base  Base
}

func (c *Cmd) Help() {
	c.Base.Help()
}

var (
	Commands = []Cmd{
		{
			Short: "h",
			Name:  "help",
			Base:  &HelpCommand{},
		},
		{
			Short: "i",
			Name:  "init",
			Base:  &InitCommand{},
		},
		{
			Short: "m",
			Name:  "model",
			Base:  &ModelCommand{},
		},
		{
			Short: "c",
			Name:  "controller",
			Base:  &ControllerCommand{},
		},
		{
			Short: "s",
			Name:  "scaffold",
			Base:  &ScaffoldCommand{},
		},
	}
)

func FindCommand(name string) Base {
	for _, c := range Commands {
		if name == c.Short || name == c.Name {
			return c.Base
		}
	}
	return nil
}
