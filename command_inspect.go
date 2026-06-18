package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	// check if the pokemon name passed in the arguments exists in your caughtPokemon name passedi n the arguments exists in your caughtPokemon map (which is in config struct)
	if len(args) < 1 {
		return fmt.Errorf("Please provide a pokemon name to inspect")
	}
	_, exists := cfg.caughtPokemon[args[0]]
	if !exists {
		return fmt.Errorf("you have not caught that pokemon")
	}

		// if it exists, print out the details of the pokemon
		pokemon := cfg.caughtPokemon[args[0]]
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, t := range pokemon.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
	fmt.Printf("\n")
	return nil
}