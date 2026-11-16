package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	fmt.Printf("Exploring %s...\n", args[0])
	pokemonsResp, err := cfg.pokeapiClient.ListPokemons(args[0])
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonsResp.PokemonEncounters {
		fmt.Println(" -", pokemon.Pokemon.Name)
	}
	// url := cfg.URL + "/" + args[0]
	// fmt.Println(url)
	return nil
}
