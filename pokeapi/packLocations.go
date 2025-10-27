package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (p *PokeApiClient) PackLocations(res *http.Response) (locationResults, error) {
	defer res.Body.Close()
	var locations locationResults

	if res.StatusCode > 299 {
		return locationResults{}, fmt.Errorf("there was an error while trying to fetch the locations: %s", res.Status)
	}

	cachedData, cacheExists := p.cache.Get(p.url)

	if cacheExists {
		if err := json.Unmarshal(cachedData, &locations); err != nil {
			return locationResults{}, fmt.Errorf("there was a problem while trying to unmarshal the res: %w", err)
		}
	} else {
		data, err := io.ReadAll(res.Body)
		p.cache.Add(p.url, data)

		if err != nil {
			return locationResults{}, fmt.Errorf("there was a problem while trying to read the res: %w", err)
		}

		if err := json.Unmarshal(data, &locations); err != nil {
			return locationResults{}, fmt.Errorf("there was a problem while trying to unmarshal the res: %w", err)
		}
	}

	return locations, nil
}
