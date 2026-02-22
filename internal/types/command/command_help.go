package command

import (
	"fmt"
)

func HelpHandler() error {
	fmt.Println("Pokédex CLI - explore, catch and manage your own, personal Pokémon collection\n")
	fmt.Println("The commands are:\n")

	availableCommands := GetCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("\t- %s : %s\n",
			cmd.Name,
			cmd.Description,
		)
	}
	fmt.Println("")
	return nil
}
