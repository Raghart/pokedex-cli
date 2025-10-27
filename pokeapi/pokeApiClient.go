package pokeapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Raghart/pokecache"
)

type PokeApiClient struct {
	url   string
	cache pokecache.Cache
}

func MakeClient(url string) *PokeApiClient {
	return &PokeApiClient{url: url, cache: *pokecache.NewCache(30 * time.Second)}
}

func (p *PokeApiClient) GetResponse() (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", p.url, nil)

	if err != nil {
		return &http.Response{}, fmt.Errorf("there was an error while trying to create the req: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return &http.Response{}, fmt.Errorf("there was an error while trying to make the req: %w", err)
	}
	return res, nil
}
