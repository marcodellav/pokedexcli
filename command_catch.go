package main

import (
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {

	pokemonsInfoResp, err := cfg.pokeapiClient.GetPokemonInfo(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	fmt.Printf("Pokemon Base XP: %d\n", pokemonsInfoResp.BaseExperience)

	return nil
}
