package main

import (
	"fmt"
)

func help(configURL *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	commandList := getCommandList()
	for key, command := range commandList {
		newCommand := fmt.Sprintf("%s: %s", key, command.description)
		fmt.Println(newCommand)
	}

	return nil
}
