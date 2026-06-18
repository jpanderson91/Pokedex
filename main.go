package main


import (
	"time"
	"github.com/jpanderson91/Pokedex/internal/pokeapi"
)



func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Minute),
	}
	startRepl(cfg)
}



