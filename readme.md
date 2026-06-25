![Pokedex Animation](./assets/pokedextwo.gif)

# Pokedex

A command-line Pokedex application built in Go that lets you explore locations, catch Pokémon, and build your collection using the [PokeAPI](https://pokeapi.co/).

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.26+

### Installation

```bash
git clone https://github.com/jpanderson91/Pokedex.git
cd Pokedex
go build -o pokedex
./pokedex
```

## Usage

Once running, you'll be presented with a `Pokedex >` prompt. Available commands:

| Command | Description |
|---------|-------------|
| `help` | Displays a help message |
| `map` | Get the next page of locations |
| `mapb` | Get the previous page of locations |
| `explore <area_name>` | Explore a location area to see available Pokémon |
| `catch <pokemon_name>` | Attempt to catch a Pokémon |
| `inspect <pokemon_name>` | Inspect a caught Pokémon's details |
| `pokedex` | View all caught Pokémon |
| `exit` | Exit the Pokedex |

### Example

```
Pokedex > map
Pokedex > explore canalave-city-area
Pokedex > catch pikachu
Pokedex > inspect pikachu
Pokedex > pokedex
```
