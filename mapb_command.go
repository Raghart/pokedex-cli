package main

import "fmt"

func travelBackMap(configURL *config) error {
	if configURL.Previous == "" {
		return fmt.Errorf("you're on the first page")
	}
	newLocations, err := getLocations(configURL.Previous)
	if err != nil {
		return fmt.Errorf("there was a problem while trying to retrieve the previous locations: %w", err)
	}

	for _, location := range newLocations.Results {
		fmt.Println(location.Name)
	}

	configURL.Previous = newLocations.Previous
	configURL.Next = newLocations.Next

	return nil
}
