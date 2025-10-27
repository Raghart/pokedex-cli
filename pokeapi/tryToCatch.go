package pokeapi

import "math/rand"

func (p PokeApiClient) TryToCatch(pokemonData PokemonStructData) bool {
	baseExp := pokemonData.BaseExperience
	failChance := int(float64(baseExp) * 0.55)
	randomInt := rand.Intn(baseExp)

	if randomInt <= failChance {
		return false
	}

	return true
}
