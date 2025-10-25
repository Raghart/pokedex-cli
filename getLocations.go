package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getLocations(url string) (locationResults, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return locationResults{}, fmt.Errorf("there was an error while trying to create the req: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return locationResults{}, fmt.Errorf("there was an error while trying to make the req: %w", err)
	}

	defer res.Body.Close()
	var locations locationResults

	if res.StatusCode > 299 {
		return locationResults{}, fmt.Errorf("there was an error while trying to fetch the locations: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationResults{}, fmt.Errorf("there was a problem while trying to read the res: %w", err)
	}

	if err := json.Unmarshal(data, &locations); err != nil {
		return locationResults{}, fmt.Errorf("there was a problem while trying to unmarshal the res: %w", err)
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return locations, nil
}
