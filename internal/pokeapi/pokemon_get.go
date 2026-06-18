package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
url := baseURL + "/pokemon/" + pokemonName

if val, ok := c.cache.Get(url); ok {
	var pokemon Pokemon
	err := json.Unmarshal(val, &pokemon)
	if err != nil {
	return Pokemon{}, err
	}
	return pokemon, nil
	}
	resp, err := c.httpClient.Get(url)
if err != nil {
	return Pokemon{}, err
	}
	
	defer resp.Body.Close()
data, err := io.ReadAll(resp.Body)
if err != nil {
	return Pokemon{}, err
	}
c.cache.Add(url, data)
var pokemon Pokemon
err = json.Unmarshal(data, &pokemon)
if err != nil {
	return Pokemon{}, err
	}
return pokemon, nil
}