package commands

import (
	"fmt"
)

func help(configURL *Config, exploredArea string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	commandList := Get()
	for key, command := range commandList {
		newCommand := fmt.Sprintf("%s: %s", key, command.Description)
		fmt.Println(newCommand)
	}

	return nil
}
