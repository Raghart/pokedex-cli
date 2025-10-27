package commands

import "fmt"

func inspectPokemon(cfg *Config, pokemon string) error {
	pokemonData, inPokedex := cfg.Pokedex[pokemon]

	if !inPokedex {
		return fmt.Errorf("you have no caught %s", pokemon)
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemonData.Name, pokemonData.Height, pokemonData.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonData.Stats {
		pokemonStrStat := fmt.Sprintf("	-%s: %d", stat.Stat.Name, stat.BaseStat)
		fmt.Println(pokemonStrStat)
	}

	fmt.Println("Types:")
	for _, pokemonType := range pokemonData.Types {
		pokemonStrType := fmt.Sprintf("	- %s", pokemonType.Type.Name)
		fmt.Println(pokemonStrType)
	}

	return nil
}
