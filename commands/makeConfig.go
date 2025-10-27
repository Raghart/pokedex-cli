package commands

import "github.com/Raghart/pokeapi"

func MakeConfig(url string) *Config {
	return &Config{Next: url, Previous: "", Pokedex: make(map[string]pokeapi.PokemonStructData)}
}
