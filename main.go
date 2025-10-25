package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	configURL := &config{Next: "https://pokeapi.co/api/v2/location-area/", Previous: ""}
	commandList := getCommandList()
	for {
		fmt.Println("Input your command! type help to see all the avaibles commands :D")
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		command, isCommand := commandList[input]

		if isCommand {
			err := command.callback(configURL)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
