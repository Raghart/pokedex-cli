package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (p *PokeApiClient) PackPokemons(res *http.Response) ([]string, error) {
	var pokemonList []string
	var pokemonData pokemonMapData

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("the area wasn't found: %s", res.Status)
	}

	pokemonCache, existsCache := p.cache.Get(p.url)
	if existsCache {
		if err := json.Unmarshal(pokemonCache, &pokemonData); err != nil {
			return pokemonList, fmt.Errorf("there was an error while trying to unpack the cache json: %w", err)
		}
	} else {
		data, err := io.ReadAll(res.Body)

		if err != nil {
			return pokemonList, fmt.Errorf("there was an error while trying to read the pokemon response: %w", err)
		}

		p.cache.Add(p.url, data)

		if err := json.Unmarshal(data, &pokemonData); err != nil {
			return pokemonList, fmt.Errorf("there was an error while trying to unpack the json: %w", err)
		}
	}

	for _, pokemon := range pokemonData.PokemonEncounters {
		pokemonList = append(pokemonList, pokemon.Pokemon.Name)
	}

	return pokemonList, nil
}
