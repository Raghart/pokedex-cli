package commands

import (
	"fmt"

	"github.com/Raghart/pokeapi"
)

func catchPokemon(cfg *Config, pokemon string) error {
	client := pokeapi.MakeClient("https://pokeapi.co/api/v2/pokemon/" + pokemon)
	res, err := client.GetResponse()

	if err != nil {
		return fmt.Errorf("there was an error while trying to retrieve the info about the pokemon: %w", err)
	}

	tryText := fmt.Sprintf("Throwing a Pokeball at %s...", pokemon)
	fmt.Println(tryText)

	pokemonData, err := client.ParsePokemonData(res)
	if err != nil {
		return fmt.Errorf("there was an error while trying to parse the pokemon response: %w", err)
	}

	isCached := client.TryToCatch(pokemonData)

	if isCached {
		_, inPokedex := cfg.Pokedex[pokemon]
		if !inPokedex {
			cfg.Pokedex[pokemon] = pokemonData
		}

		successText := fmt.Sprintf("%s was caught!", pokemon)
		fmt.Println(successText)
		fmt.Println("You may now inspect it with the 'inspect' command.")
	} else {
		failText := fmt.Sprintf("%s escaped!", pokemon)
		fmt.Println(failText)
	}

	return nil
}
