package commands

import (
	"fmt"

	"github.com/Raghart/pokeapi"
)

func travelMap(configURL *Config, exploreArea string) error {
	client := pokeapi.MakeClient(configURL.Next)
	res, err := client.GetResponse()

	if err != nil {
		return fmt.Errorf("there was a problem while trying to retrieve the next locations: %w", err)
	}

	newLocations, err := client.PackLocations(res)

	if err != nil {
		return fmt.Errorf("there was an error while trying to pack the response: %w", err)
	}

	for _, location := range newLocations.Results {
		fmt.Println(location.Name)
	}

	configURL.Next = newLocations.Next
	configURL.Previous = newLocations.Previous

	return nil
}
