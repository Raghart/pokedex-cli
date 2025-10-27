package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (p PokeApiClient) ParsePokemonData(res *http.Response) (PokemonStructData, error) {
	var pokemonData PokemonStructData
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return pokemonData, fmt.Errorf("there was an error with the res: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return pokemonData, fmt.Errorf("there was an error while trying to read the response: %w", err)
	}

	if err := json.Unmarshal(data, &pokemonData); err != nil {
		return pokemonData, fmt.Errorf("there was an error while trying to unpack the pokemon data: %w", err)
	}

	return pokemonData, nil
}
