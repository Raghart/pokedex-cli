package main

import (
	"fmt"
)

func travelMap(configURL *config) error {
	newLocations, err := getLocations(configURL.Next)

	if err != nil {
		return fmt.Errorf("there was a problem while trying to retrieve the next locations: %w", err)
	}

	for _, location := range newLocations.Results {
		fmt.Println(location.Name)
	}

	configURL.Next = newLocations.Next
	configURL.Previous = newLocations.Previous

	return nil
}
