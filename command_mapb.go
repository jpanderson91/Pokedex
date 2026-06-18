package main

import (
	"fmt"
	"errors"

)
func commandMapb(cfg *config, _ ...string) error {
        if cfg.prevPageURL == nil {
        return errors.New("you're on the first page")
    }
	resp, err := cfg.pokeapiClient.ListLocations(cfg.prevPageURL)
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