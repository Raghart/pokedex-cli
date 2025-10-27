package commands

func Get() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    help,
		},
		"map": {
			Name:        "map",
			Description: "travel through 20 locations in the pokemon world",
			Callback:    travelMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "travel through the previous 20 locations in the pokemon world",
			Callback:    travelBackMap,
		},
		"explore": {
			Name:        "explore",
			Description: "explore the location of a zone after seeing it with map",
			Callback:    exploreMap,
		},
		"catch": {
			Name:        "catch",
			Description: "try to catch the pokemon next of the catch command",
			Callback:    catchPokemon,
		},
		"inspect": {
			Name:        "inspect",
			Description: "inspect the pokemon next pokemon if it was registered in the pokedex",
			Callback:    inspectPokemon,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "check out all the pokemons that you have registered in your pokedex!",
			Callback:    checkPokedex,
		},
	}
}
