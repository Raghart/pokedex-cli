package commands

import "fmt"

func checkPokedex(cfg *Config, pokemon string) error {
	if len(cfg.Pokedex) == 0 {
		return fmt.Errorf("you have no pokemon registered in your pokedex, use the 'catch' command to catch some!")
	}

	fmt.Println("Your pokedex:")
	for key, _ := range cfg.Pokedex {
		pokemonStr := fmt.Sprintf("	- %s", key)
		fmt.Println(pokemonStr)
	}
	return nil
}
