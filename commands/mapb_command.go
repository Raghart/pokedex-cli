package commands

import (
	"fmt"

	"github.com/Raghart/pokeapi"
)

func travelBackMap(configURL *Config, exploreArea string) error {
	if configURL.Previous == "" {
		return fmt.Errorf("you're on the first page")
	}
	client := pokeapi.MakeClient(configURL.Previous)
	res, err := client.GetResponse()

	if err != nil {
		return fmt.Errorf("there was a problem while trying to retrieve the previous locations: %w", err)
	}
	newLocations, err := client.PackLocations(res)

	if err != nil {
		return fmt.Errorf("there was a problem while trying to pack the locations: %w", err)
	}

	for _, location := range newLocations.Results {
		fmt.Println(location.Name)
	}

	configURL.Previous = newLocations.Previous
	configURL.Next = newLocations.Next

	return nil
}
