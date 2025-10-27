package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Raghart/commands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	configURL := commands.MakeConfig("https://pokeapi.co/api/v2/location-area/")
	commandList := commands.Get()

	var exploreArea string
	for {
		fmt.Println("Input your command! type help to see all the avaibles commands :D")
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) > 1 {
			exploreArea = input[1]
		}
		command, isCommand := commandList[input[0]]

		if isCommand {
			err := command.Callback(configURL, exploreArea)
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
