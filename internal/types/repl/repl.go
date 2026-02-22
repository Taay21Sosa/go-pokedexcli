package repl

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/types/command"
	"strings"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" > ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := CleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		availableCommands := command.GetCommands()
		action, ok := availableCommands[commandName]
		if !ok {
			fmt.Printf("Unknown command: %s\n", commandName)
			continue
		}
		action.Callback()
	}
}

func CleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
