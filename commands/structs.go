package commands

import "github.com/Raghart/pokeapi"

type Config struct {
	Next     string
	Previous string
	Pokedex  map[string]pokeapi.PokemonStructData
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, string) error
}
