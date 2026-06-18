package main

import (
    "fmt"
    "strings"
)

func commandExplore(cfg *config, args ...string) error {
    if len(args) == 0 {
        return fmt.Errorf("usage: explore <area_name>")
    }
    // Join remaining args in case area name contains spaces (user should use hyphens normally)
    area := strings.Join(args, " ")
    fmt.Printf("Exploring %s...\n", area)

    resp, err := cfg.pokeapiClient.GetLocationArea(area)
    if err != nil {
        return err
    }

    fmt.Println("Found Pokemon:")
    for _, pe := range resp.PokemonEncounters {
        fmt.Printf(" - %s\n", pe.Pokemon.Name)
    }

    return nil
}
