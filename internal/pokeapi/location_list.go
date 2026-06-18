package pokeapi

import (
	"encoding/json"
	"io"
	"bytes"


)


func (c Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	// accept a *string URL parameter (which may be nil on the first call), if the URL is nil, use the default endpoint: baseURL + "/location-area", make a GET request to that URL, read the response body, unmarshal the JSON into a RespShallowLocations struct, return it
	var url string
	if pageURL == nil {
		url = baseURL + "/location-area"
	} else {
		url = *pageURL
	}
	// check cache first
	if val, ok := c.cache.Get(url); ok {
    // decode val ([]byte) into locations and return
		var locations RespShallowLocations
		err := json.NewDecoder(bytes.NewReader(val)).Decode(&locations)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locations, nil	
	}
	resp, err := c.httpClient.Get(url)
	if err != nil {
    	return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
    	return RespShallowLocations{}, err
	}

	c.cache.Add(url, data)

	var locations RespShallowLocations
	err = json.NewDecoder(bytes.NewReader(data)).Decode(&locations)
	if err != nil {
    	return RespShallowLocations{}, err
	}
	return locations, nil
}