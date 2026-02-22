package command

type Command struct {
	Name        string
	Description string
	Callback    func() error
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "displays a list of available commands and their usage instructions.",
			Callback:    HelpHandler,
		},
		"exit": {
			Name:        "exit",
			Description: "exits the application or repl (read-eval-print loop) environment.",
			Callback:    ExitHandler,
		},
	}
}
