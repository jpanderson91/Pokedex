package pokeapi

import (
	"encoding/json"

)


func (c Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	// accept a *string URL parameter (which may be nil on the first call), if the URL is nil, use the default endpoint: baseURL + "/location-area", make a GET request to that URL, read the response body, unmarshal the JSON into a RespShallowLocations struct, return it
	var url string
	if pageURL == nil {
		url = baseURL + "/location-area"
	} else {
		url = *pageURL
	}
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()
	var locations RespShallowLocations
	err = json.NewDecoder(resp.Body).Decode(&locations)
	if err != nil {
		return RespShallowLocations{}, err
	}
	return locations, nil
}