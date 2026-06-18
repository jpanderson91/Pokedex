package pokeapi

import (
    "encoding/json"
    "io"
    "bytes"
    "net/url"
)

type RespLocationArea struct {
    ID int `json:"id"`
    Name string `json:"name"`
    PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
    Pokemon ShallowLocation `json:"pokemon"`
}

// GetLocationArea fetches detailed info for a location-area by name or id.
func (c Client) GetLocationArea(nameOrID string) (RespLocationArea, error) {
    // construct URL
    // ensure nameOrID is URL-safe
    escaped := url.PathEscape(nameOrID)
    urlStr := baseURL + "/location-area/" + escaped

    // check cache
    if val, ok := c.cache.Get(urlStr); ok {
        var area RespLocationArea
        err := json.NewDecoder(bytes.NewReader(val)).Decode(&area)
        if err != nil {
            return RespLocationArea{}, err
        }
        return area, nil
    }	

    resp, err := c.httpClient.Get(urlStr)
    if err != nil {
        return RespLocationArea{}, err
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return RespLocationArea{}, err
    }

    c.cache.Add(urlStr, data)

    var area RespLocationArea
    err = json.NewDecoder(bytes.NewReader(data)).Decode(&area)
    if err != nil {
        return RespLocationArea{}, err
    }
    return area, nil
}
