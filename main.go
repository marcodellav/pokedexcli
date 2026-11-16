package main

import (
	"time"

	"github.com/marcodellav/pokedexcli/internal/pokeapi"
)

func main() {
	s := "https://pokeapi.co/api/v2/location-area"
	c := &config{
		commands:      getCommands(),
		URL:           s,
		Next:          nil,
		Previous:      nil,
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}
	startRepl(c)
}
