package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
	"github.com/jpanderson91/Pokedex/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	nextPageURL *string
	prevPageURL *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: "map",
			description: "Get the next page of locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Get the previous page of locations",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Explore a location area: explore <area_name>",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Attempt to catch a pokemon: catch <pokemon_name>",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect a caught pokemon: inspect <pokemon_name>",
			callback: commandInspect,
		},
		"pokedex": {
			
			name: "pokedex",
			description: "View all caught pokemon",
			callback: commandPokedex,

		},
	}
}
func cleanInput(text string) []string {
	// split user's input into "words" based on whitespace
	// should also lowercase the input and trim any leading or trailing whitespace. For example:
	// hello world -> ["hello", "world"]
	trimmed := strings.TrimSpace(strings.ToLower(text))
	words := strings.Fields(trimmed)
	return words
}


func startRepl(cfg *config) {
	//infinite for loop. This lool will execute once for every command the user types in (we don't want to exit the program after just one command)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		// Use the scanner's .Scan and .Text methods to get the user's input as a string
		scanner.Scan()
		input := scanner.Text()
		// clean the user's input string
		cleanedInput := cleanInput(input)
		// capture the first "word" of the input and use it to print: Your command was : <first word>
		if len(cleanedInput) > 0 {
			commandName := cleanedInput[0]
			command, exists := getCommands()[commandName]
			if exists {
				// call the callback here, passing any args after the command
				args := []string{}
				if len(cleanedInput) > 1 {
					args = cleanedInput[1:]
				}
				err := command.callback(cfg, args...)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				// print "Unknown command"
				fmt.Println("Unknown command")
			}
		}
	}
}
