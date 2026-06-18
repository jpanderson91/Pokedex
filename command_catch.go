package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("requires exactly one Pokemon name")
	}

	name := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	max := pokemon.BaseExperience
	if max < 1 {
		max = 1
	}

	threshold := 50
	if rand.Intn(max) < threshold {
		cfg.caughtPokemon[pokemon.Name] = pokemon
		fmt.Printf("You caught %s!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}