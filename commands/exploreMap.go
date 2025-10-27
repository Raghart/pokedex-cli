package commands

import (
	"fmt"

	"github.com/Raghart/pokeapi"
)

func exploreMap(cfg *Config, exploreArea string) error {
	client := pokeapi.MakeClient("https://pokeapi.co/api/v2/location-area/" + exploreArea)
	res, err := client.GetResponse()
	if err != nil {
		return fmt.Errorf("there was an error while trying to get the response: %w", err)
	}
	pokemonList, err := client.PackPokemons(res)
	if err != nil {
		return fmt.Errorf("there was an error while trying to pack the pokemons: %w", err)
	}

	for _, pokemon := range pokemonList {
		fmt.Println(pokemon)
	}
	return nil
}
