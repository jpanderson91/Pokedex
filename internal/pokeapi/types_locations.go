package pokeapi

type RespShallowLocations struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []ShallowLocation `json:"results"`
}

type ShallowLocation struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
