package main
import (
	"fmt"

)
func commandMap(cfg *config) error {
    resp, err := cfg.pokeapiClient.ListLocations(cfg.nextPageURL)
    if err != nil {
        return err
    }
	cfg.nextPageURL = resp.Next
	cfg.prevPageURL = resp.Previous
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
    return nil
}